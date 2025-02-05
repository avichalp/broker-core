package dealer

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/textileio/broker-core/cmd/dealerd/store"
	mbroker "github.com/textileio/broker-core/msgbroker"
	"github.com/textileio/broker-core/ratelim"
)

var (
	failureDealMakingMaxRetries = "deal rejected: reached maximum amount of deal-making retries"
)

func (d *Dealer) daemonDealMaker() {
	defer d.daemonWg.Done()

	for {
		select {
		case <-d.daemonCtx.Done():
			log.Infof("deal maker daemon closed")
			return
		case <-time.After(d.config.dealMakingFreq):
			if err := d.daemonDealMakerTick(); err != nil {
				log.Errorf("deal maker tick: %s", err)
			}
		}
	}
}

func (d *Dealer) daemonDealMakerTick() error {
	rl, err := ratelim.New(d.config.dealMakingRateLim)
	if err != nil {
		return fmt.Errorf("create ratelim: %s", err)
	}

	for {
		if d.daemonCtx.Err() != nil {
			break
		}

		ctx, cancel := context.WithTimeout(d.daemonCtx, time.Second*15)
		defer cancel()
		aud, ok, err := d.store.GetNextPending(ctx, store.StatusDealMaking)
		if err != nil {
			return fmt.Errorf("get pending auction deals: %s", err)
		}
		if !ok {
			break
		}
		rl.Exec(func() error {
			ctx, cancel := context.WithTimeout(d.daemonCtx, time.Second*30)
			defer cancel()
			if err := d.executePendingDealMaking(ctx, aud); err != nil {
				log.Errorf("executing pending deal making: %s", err)
				aud.ErrorCause = err.Error()
				aud.ReadyAt = time.Unix(0, 0)
				if err := d.store.SaveAndMoveAuctionDeal(d.daemonCtx, aud, store.StatusReportFinalized); err != nil {
					log.Errorf("saving auction deal: %s", err)
				}
			}
			// We're not interested in ratelim error inspection.
			return nil
		})
		cancel()
	}
	rl.Wait()

	return nil
}

func (d *Dealer) executePendingDealMaking(ctx context.Context, aud store.AuctionDeal) error {
	ad, err := d.store.GetAuctionData(ctx, aud.AuctionDataID)
	if err != nil {
		return fmt.Errorf("get auction data %s: %s", aud.AuctionDataID, err)
	}

	rw, err := d.store.GetRemoteWallet(ctx, aud.AuctionDataID)
	if err != nil {
		return fmt.Errorf("get remote wallet info: %s", err)
	}

	allowBoost, err := d.store.IsBoostAllowed(ctx, aud.StorageProviderID)
	if err != nil {
		return fmt.Errorf("is boost allowed for %s: %s", aud.StorageProviderID, err)
	}
	log.Debugf("%s executing deal from SD %s for %s with storage-provider %s, boost: %v",
		aud.ID, ad.BatchID, ad.PayloadCid, aud.StorageProviderID, allowBoost)

	dealIdentifier, retry, err := d.filclient.ExecuteAuctionDeal(ctx, ad, aud, rw, allowBoost)
	if err != nil {
		return fmt.Errorf("executing auction deal: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// If retry, then we move from ExecutingDealMaking back to PendingDealMaking
	// with some delay as to retry again. If we tried dealMakingMaxRetries,
	// we give up and error the auction deal.
	if retry {
		aud.Retries++
		if aud.Retries > d.config.dealMakingMaxRetries {
			log.Warnf("deal for %s with %s failed the max number of retries, failing", ad.PayloadCid, aud.StorageProviderID)
			aud.ErrorCause = failureDealMakingMaxRetries
			aud.ReadyAt = time.Unix(0, 0)
			if err := d.store.SaveAndMoveAuctionDeal(ctx, aud, store.StatusReportFinalized); err != nil {
				return fmt.Errorf("saving auction deal for final retry: %s", err)
			}
			return nil
		}

		log.Warnf("deal for %s with %s failed, we'll retry soon...", ad.PayloadCid, aud.StorageProviderID)
		aud.ReadyAt = time.Now().Add(d.config.dealMakingRetryDelay * time.Duration(aud.Retries))
		if err := d.store.SaveAndMoveAuctionDeal(ctx, aud, store.StatusDealMaking); err != nil {
			return fmt.Errorf("saving auction deal for future retry: %s", err)
		}
		return nil
	}

	log.Infof("deal with payloadcid %s with %s successfully executed", ad.PayloadCid, aud.StorageProviderID)
	aud.Retries = 0
	aud.ProposalCid = dealIdentifier
	aud.ReadyAt = time.Unix(0, 0)
	if err := d.store.SaveAndMoveAuctionDeal(ctx, aud, store.StatusConfirmation); err != nil {
		return fmt.Errorf("changing status to WaitingConfirmation: %s", err)
	}

	log.Debugf("accepted deal proposal %s from payloadcid %s", dealIdentifier, ad.PayloadCid)

	// If the dealIdentifier isn't a UUID, then it's a PropsalCid. (Legacy case)
	if _, err := uuid.Parse(dealIdentifier); err != nil {
		proposalCid, err := cid.Decode(dealIdentifier)
		if err != nil {
			return fmt.Errorf("invalid proposal cid %s: %s", dealIdentifier, err)
		}

		log.Debugf("notifying about proposalcid %s", proposalCid)
		if err := mbroker.PublishMsgDealProposalAccepted(
			ctx,
			d.mb,
			ad.BatchID,
			aud.AuctionID,
			aud.BidID,
			aud.StorageProviderID,
			proposalCid); err != nil {
			return fmt.Errorf("publish deal-proposal-accepted msg of proposal %s to msgbroker: %s", proposalCid, err)
		}
	} else {
		// It's a UUID, thus a Boost UUID identifier.
		log.Debugf("notifying about dealuuid %s", dealIdentifier)
		if err := mbroker.PublishMsgBoostDealProposalAccepted(
			ctx,
			d.mb,
			ad.BatchID,
			aud.AuctionID,
			aud.BidID,
			aud.StorageProviderID,
			dealIdentifier); err != nil {
			return fmt.Errorf("publish boost-deal-proposal-accepted msg for %s to msgbroker: %s", dealIdentifier, err)
		}
	}

	return nil
}

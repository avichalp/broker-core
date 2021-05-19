package auctioneer

import (
	"context"

	"github.com/textileio/broker-core/cmd/auctioneerd/metrics"
	"go.opentelemetry.io/otel/metric"
)

func (a *Auctioneer) initMetrics() {
	a.metricNewAuction = metrics.Meter.NewInt64Counter(metrics.Prefix + ".auctions_total")
	a.metricNewFinalizedAuction = metrics.Meter.NewInt64Counter(metrics.Prefix + ".finalized_auctions_total")
	a.metricNewBid = metrics.Meter.NewInt64Counter(metrics.Prefix + ".bids_total")
	a.metricLastCreatedAuction = metrics.Meter.NewInt64ValueObserver(metrics.Prefix+".last_created_auction_epoch", a.lastCreatedAuctionCb)
}

func (a *Auctioneer) lastCreatedAuctionCb(ctx context.Context, r metric.Int64ObserverResult) {
	r.Observe(a.statLastCreatedAuction.Unix())
}

package httpapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/textileio/broker-core/broker"
	bidstore "github.com/textileio/broker-core/cmd/bidbot/service/store"
	golog "github.com/textileio/go-log/v2"
)

var (
	log = golog.Logger("bidbot/api")
)

// Service provides scoped access to the bidbot service.
type Service interface {
	ListBids(query bidstore.Query) ([]*bidstore.Bid, error)
	GetBid(id broker.BidID) (*bidstore.Bid, error)
	WriteCar(ctx context.Context, cid cid.Cid) (string, error)
}

// NewServer returns a new http server for bidbot commands.
func NewServer(listenAddr string, service Service) (*http.Server, error) {
	httpServer := &http.Server{
		Addr:              listenAddr,
		ReadHeaderTimeout: time.Second * 5,
		WriteTimeout:      time.Second * 10,
		Handler:           createMux(service),
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("stopping http server: %s", err)
		}
	}()

	log.Infof("http server started at %s", listenAddr)
	return httpServer, nil
}

func createMux(service Service) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", getOnly(healthHandler))
	// allow both with and without trailing slash
	deals := getOnly(dealsHandler(service))
	mux.HandleFunc("/deals", deals)
	mux.HandleFunc("/deals/", deals)
	mux.HandleFunc("/cid/", getOnly(writeCarRequestHandler(service)))
	return mux
}

func getOnly(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			httpError(w, "only GET method is allowed", http.StatusBadRequest)
			return
		}
		f(w, r)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func dealsHandler(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bids []*bidstore.Bid
		urlParts := strings.SplitN(r.URL.Path, "/", 3)
		if len(urlParts) < 3 || urlParts[2] == "" {
			statusFilters := strings.Split(r.URL.Query().Get("status"), ",")
			var proceed bool
			if bids, proceed = listBids(w, service, statusFilters); !proceed {
				return
			}
		} else {
			bid, err := service.GetBid(broker.BidID(urlParts[2]))
			if err == nil {
				bids = append(bids, bid)
			} else if err != bidstore.ErrBidNotFound {
				httpError(w, fmt.Sprintf("get bid: %s", err), http.StatusInternalServerError)
				return
			}
		}
		data, err := json.Marshal(bids)
		if err != nil {
			httpError(w, fmt.Sprintf("json encoding: %s", err), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(data)
		if err != nil {
			log.Errorf("write failed: %+v", err)
		}
	}
}

func writeCarRequestHandler(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.SplitN(r.URL.Path, "/", 3)
		if len(urlParts) < 3 {
			httpError(w, "the url should be /cid/{cid}", http.StatusBadRequest)
			return
		}
		id, err := cid.Decode(urlParts[2])
		if err != nil {
			httpError(w, fmt.Sprintf("decoding cid: %s", err), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), bidstore.DataCidFetchTimeout)
		defer cancel()
		dest, err := service.WriteCar(ctx, id)
		if err != nil {
			httpError(w, fmt.Sprintf("writing car: %s", err), http.StatusInternalServerError)
			return
		}

		resp := fmt.Sprintf("downloaded to %s", dest)
		if _, err = w.Write([]byte(resp)); err != nil {
			httpError(w, fmt.Sprintf("writing response: %s", err), http.StatusInternalServerError)
			return
		}
	}
}

func httpError(w http.ResponseWriter, err string, status int) {
	log.Debugf("request error: %s", err)
	http.Error(w, err, status)
}

func listBids(w http.ResponseWriter, service Service, statusFilters []string) (bids []*bidstore.Bid, proceed bool) {
	var filters []bidstore.BidStatus
	for _, s := range statusFilters {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if bs, err := bidstore.BidStatusByString(s); err != nil {
			httpError(w, fmt.Sprintf("%s: %s", s, err), http.StatusBadRequest)
			return nil, false
		} else {
			filters = append(filters, bs)
		}
	}
	// for simplicity we apply filters after retrieving. if performance
	// becomes an issue, we can add query filters.
	fullList, err := service.ListBids(bidstore.Query{Limit: -1})
	if err != nil {
		httpError(w, fmt.Sprintf("listing bids: %s", err), http.StatusInternalServerError)
		return nil, false
	}
	if len(filters) == 0 {
		return fullList, true
	}
	for _, bid := range fullList {
		for _, status := range filters {
			if bid.Status == status {
				bids = append(bids, bid)
				break
			}
		}
	}
	return bids, true
}

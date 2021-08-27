package auctioneer

import (
	"container/heap"
	"context"
	"math/rand"

	"github.com/textileio/broker-core/auctioneer"
)

// Cmp is the interface for a comparator.
type Cmp interface {
	// Cmp returns arbitrary number with the following semantics:
	// negative: i is considered to be less than j
	// zero: i is considered to be equal to j
	// positive: i is considered to be greater than j
	Cmp(a *auctioneer.Auction, i auctioneer.Bid, j auctioneer.Bid) int
}

// CmpFn is a helper which turns a function to a Cmp interface.
func CmpFn(f func(a *auctioneer.Auction, i auctioneer.Bid, j auctioneer.Bid) int) Cmp {
	return fnCmp{f: f}
}

type fnCmp struct {
	f func(*auctioneer.Auction, auctioneer.Bid, auctioneer.Bid) int
}

func (c fnCmp) Cmp(a *auctioneer.Auction, i auctioneer.Bid, j auctioneer.Bid) int {
	return c.f(a, i, j)
}

type ordered struct {
	cmps []Cmp
}

// Ordered executes each comparator in order, i.e., if the first comparator
// judges the two bids to be equal, continues to the next comparator, and so
// on. It considers two bids to be equal if all comparators are exhausted.
func Ordered(cmps ...Cmp) Cmp {
	return ordered{cmps}
}

func (c ordered) Cmp(a *auctioneer.Auction, i auctioneer.Bid, j auctioneer.Bid) int {
	for _, c := range c.cmps {
		result := c.Cmp(a, i, j)
		switch result {
		case 0:
			continue
		default:
			return result
		}
	}
	return 0
}

// Weighed combines comparators togethers with different weights. The result
// depends on both the weights given to each comparator and the scale of the
// comparison result of each comparator. Be aware to not cause integer overflow.
type Weighed struct {
	cmps    []Cmp
	weights []int
}

// Add returns a new weighed comparator with the comparator being added with
// the given weight.
func (wc Weighed) Add(cmp Cmp, weight int) Weighed {
	w := Weighed{cmps: wc.cmps, weights: wc.weights}
	w.cmps = append(w.cmps, cmp)
	w.weights = append(w.weights, weight)
	return w
}

func (wc Weighed) Cmp(a *auctioneer.Auction, i auctioneer.Bid, j auctioneer.Bid) int {
	var weighed int
	for k, cmp := range wc.cmps {
		weighed += wc.weights[k] * cmp.Cmp(a, i, j)
	}
	return weighed
}

// ByPrice returns a comparator which prefers lower ask price or verified ask
// price depending on if the auction is verified. The price difference is returned.
func ByPrice() Cmp {
	return CmpFn(func(a *auctioneer.Auction, i auctioneer.Bid, j auctioneer.Bid) int {
		if a.DealVerified {
			return int(i.VerifiedAskPrice - j.VerifiedAskPrice)
		}
		return int(i.AskPrice - j.AskPrice)
	})
}

// ByStartEpoch returns a comparator which prefers bid with smaller start
// epoch. The difference of the start epoch is returned.
func ByStartEpoch() Cmp {
	return CmpFn(func(a *auctioneer.Auction, i auctioneer.Bid, j auctioneer.Bid) int {
		return int(i.StartEpoch) - int(j.StartEpoch)
	})
}

type providerRate struct {
	rates map[string]int
}

func (wc providerRate) Cmp(a *auctioneer.Auction, i auctioneer.Bid, j auctioneer.Bid) int {
	return wc.rates[i.StorageProviderID] - wc.rates[j.StorageProviderID]
}

// ByProviderRate returns a comparator which considers some rate of the storage
// provider. Provider with lower rate gets a higher chance to win. Provider
// not in the provided rates table are considered to have zero rate.
func ByProviderRate(rates map[string]int) Cmp {
	return providerRate{rates}
}

// Random returns a comparator which randomly returns -1, 0, or 1 using the
// given source.
func Random(s rand.Source) Cmp {
	return CmpFn(func(a *auctioneer.Auction, i auctioneer.Bid, j auctioneer.Bid) int {
		return int(s.Int63()%3 - 1)
	})
}

// BidsSorter constructs a sorter from the given comparator.
func BidsSorter(cmp Cmp) Sorter { return Sorter{cmp} }

// Sorter has a single sort method which takes an aunction and some bids, then
// sort the bids based on the comparator given.
type Sorter struct {
	cmp Cmp
}

// Sort does the sort work and sends the result to the channel one by one. The
// channel is closed when all bids are exhausted or the context is done.
func (s Sorter) Sort(ctx context.Context, auction *auctioneer.Auction, bids []auctioneer.Bid) chan auctioneer.Bid {
	ret := make(chan auctioneer.Bid)
	go func() {
		defer close(ret)
		bh := &BidHeap{a: auction, cmp: s.cmp}
		heap.Init(bh)
		for _, b := range bids {
			heap.Push(bh, b)
		}
		for {
			if bh.Len() == 0 {
				return
			}
			b := heap.Pop(bh).(auctioneer.Bid)
			select {
			case <-ctx.Done():
				return
			case ret <- b:
			}
		}
	}()
	return ret
}

// BidHeap is used to efficiently select auction winners.
type BidHeap struct {
	a   *auctioneer.Auction
	h   []auctioneer.Bid
	cmp Cmp
}

// Len returns the length of h.
func (bh BidHeap) Len() int {
	return len(bh.h)
}

// Less returns true if the value at j is less than the value at i.
func (bh BidHeap) Less(i, j int) bool {
	return bh.cmp.Cmp(bh.a, bh.h[i], bh.h[j]) < 0
}

// Swap index i and j.
func (bh BidHeap) Swap(i, j int) {
	bh.h[i], bh.h[j] = bh.h[j], bh.h[i]
}

// Push adds x to h.
func (bh *BidHeap) Push(x interface{}) {
	bh.h = append(bh.h, x.(auctioneer.Bid))
}

// Pop removes and returns the last element in h.
func (bh *BidHeap) Pop() (x interface{}) {
	x, bh.h = bh.h[len(bh.h)-1], bh.h[:len(bh.h)-1]
	return x
}

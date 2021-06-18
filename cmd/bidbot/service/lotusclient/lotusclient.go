package lotusclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/ipfs/go-cid"
	ma "github.com/multiformats/go-multiaddr"
	dns "github.com/multiformats/go-multiaddr-dns"
	"github.com/textileio/broker-core/finalizer"
	golog "github.com/textileio/go-log/v2"
)

var (
	log = golog.Logger("bidbot/lotus")

	requestTimeout = time.Second * 10
)

// LotusClient provides access to Lotus for importing deal data.
type LotusClient interface {
	io.Closer

	ImportData(pcid cid.Cid, file string) error
}

// Client provides access to Lotus for importing deal data.
type Client struct {
	c        v0api.StorageMiner
	fakeMode bool

	ctx       context.Context
	finalizer *finalizer.Finalizer
}

// New returns a new *LotusClient.
func New(maddr string, authToken string, connRetries int, fakeMode bool) (*Client, error) {
	fin := finalizer.NewFinalizer()
	ctx, cancel := context.WithCancel(context.Background())
	fin.Add(finalizer.NewContextCloser(cancel))

	lc := &Client{
		fakeMode:  fakeMode,
		ctx:       ctx,
		finalizer: fin,
	}
	if lc.fakeMode {
		return lc, nil
	}

	builder, err := newBuilder(maddr, authToken, connRetries)
	if err != nil {
		return nil, fmt.Errorf("building lotus client: %v", err)
	}
	c, closer, err := builder(lc.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting lotus client: %v", err)
	}
	fin.AddFn(closer)

	return &Client{
		c:         c,
		fakeMode:  fakeMode,
		ctx:       ctx,
		finalizer: fin,
	}, nil
}

// Close the client.
func (c *Client) Close() error {
	return c.finalizer.Cleanup(nil)
}

// ImportData imports deal data into Lotus.
func (c *Client) ImportData(pcid cid.Cid, file string) error {
	if c.fakeMode {
		return nil
	}
	ctx, cancel := context.WithTimeout(c.ctx, requestTimeout)
	defer cancel()
	if err := c.c.DealsImportData(ctx, pcid, file); err != nil {
		return fmt.Errorf("calling storage miner deals import data: %v", err)
	}
	return nil
}

type clientBuilder func(ctx context.Context) (*v0api.StorageMinerStruct, func(), error)

func newBuilder(maddrs string, authToken string, connRetries int) (clientBuilder, error) {
	maddr, err := ma.NewMultiaddr(maddrs)
	if err != nil {
		return nil, fmt.Errorf("parsing multiaddress: %v", err)
	}
	addr, err := tcpAddrFromMultiAddr(maddr)
	if err != nil {
		return nil, fmt.Errorf("getting tcp address from multiaddress: %v", err)
	}
	headers := http.Header{
		"Authorization": []string{"Bearer " + authToken},
	}

	return func(ctx context.Context) (*v0api.StorageMinerStruct, func(), error) {
		var api v0api.StorageMinerStruct
		var closer jsonrpc.ClientCloser
		var err error
		for i := 0; i < connRetries; i++ {
			if ctx.Err() != nil {
				return nil, nil, fmt.Errorf("canceled by context")
			}
			closer, err = jsonrpc.NewMergeClient(context.Background(), "ws://"+addr+"/rpc/v0", "Filecoin",
				[]interface{}{
					&api.Internal,
				}, headers)
			if err == nil {
				break
			}
			log.Warnf("failed to connect to Lotus API %s, retrying...", err)
			time.Sleep(time.Second * 10)
		}
		if err != nil {
			return nil, nil, fmt.Errorf("couldn't connect to Lotus API: %s", err)
		}

		return &api, closer, nil
	}, nil
}

func tcpAddrFromMultiAddr(maddr ma.Multiaddr) (string, error) {
	if maddr == nil {
		return "", fmt.Errorf("invalid address")
	}

	var ip string
	if _, err := maddr.ValueForProtocol(ma.P_DNS4); err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		maddrs, err := dns.Resolve(ctx, maddr)
		if err != nil {
			return "", fmt.Errorf("resolving dns: %s", err)
		}
		for _, m := range maddrs {
			if ip, err = getIPFromMaddr(m); err == nil {
				break
			}
		}
	} else {
		ip, err = getIPFromMaddr(maddr)
		if err != nil {
			return "", fmt.Errorf("getting ip from maddr: %s", err)
		}
	}

	tcp, err := maddr.ValueForProtocol(ma.P_TCP)
	if err != nil {
		return "", fmt.Errorf("getting port from maddr: %s", err)
	}
	return fmt.Sprintf("%s:%s", ip, tcp), nil
}

func getIPFromMaddr(maddr ma.Multiaddr) (string, error) {
	if ip, err := maddr.ValueForProtocol(ma.P_IP4); err == nil {
		return ip, nil
	}
	if ip, err := maddr.ValueForProtocol(ma.P_IP6); err == nil {
		return fmt.Sprintf("[%s]", ip), nil
	}
	return "", fmt.Errorf("no ip in multiaddr")
}

package marketpeer

import (
	"time"

	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/spf13/viper"
	"github.com/textileio/broker-core/cmd/common"
)

// Flags defines daemon flags for a marketpeer.
var Flags = []common.Flag{
	{
		Name:        "repo",
		DefValue:    "${HOME}/.miner",
		Description: "Repo path",
	},
	{
		Name:        "listen-multiaddr",
		DefValue:    "/ip4/0.0.0.0/tcp/4001",
		Description: "Libp2p listen multiaddr; repeatable",
		Repeatable:  true,
	},
	{
		Name:        "bootstrap-multiaddr",
		DefValue:    "",
		Description: "Libp2p bootstrap peer multiaddr; repeatable",
		Repeatable:  true,
	},
	{
		Name:        "announce-multiaddr",
		DefValue:    "",
		Description: "Libp2p annouce multiaddr; repeatable",
		Repeatable:  true,
	},
	{
		Name:        "conn-low",
		DefValue:    256,
		Description: "Libp2p connection manager low water mark",
	},
	{
		Name:        "conn-high",
		DefValue:    512,
		Description: "Libp2p connection manager high water mark",
	},
	{
		Name:        "conn-grace",
		DefValue:    time.Second * 120,
		Description: "Libp2p connection manager grace period",
	},
	{
		Name:        "quic",
		DefValue:    false,
		Description: "Enable the QUIC transport",
	},
	{
		Name:        "mdns",
		DefValue:    false,
		Description: "Enable MDNS peer discovery",
	},
	{
		Name:        "nat",
		DefValue:    false,
		Description: "Enable NAT port mapping",
	},
}

// ConfigFromFlags returns a Config from a *viper.Viper instance.
func ConfigFromFlags(v *viper.Viper) Config {
	return Config{
		RepoPath:           v.GetString("repo"),
		ListenMultiaddrs:   v.GetStringSlice("listen-multiaddr"),
		AnnounceMultiaddrs: v.GetStringSlice("announce-multiaddr"),
		BootstrapAddrs:     v.GetStringSlice("bootstrap-multiaddr"),
		ConnManager: connmgr.NewConnManager(
			v.GetInt("conn-low"),
			v.GetInt("conn-high"),
			v.GetDuration("conn-grace"),
		),
		EnableQUIC:       v.GetBool("quic"),
		EnableNATPortMap: v.GetBool("nat"),
	}
}

package main

// TODO: Add filter and bid params to config.

import (
	"encoding/json"
	_ "net/http/pprof"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/textileio/broker-core/broker"
	"github.com/textileio/broker-core/cmd/common"
	"github.com/textileio/broker-core/cmd/minerd/service"
	"github.com/textileio/broker-core/marketpeer"
)

var (
	daemonName = "minerd"
	log        = logging.Logger(daemonName)
	v          = viper.New()
)

func init() {
	flags := []common.Flag{
		{Name: "debug", DefValue: false, Description: "Enable debug level logs"},
		{Name: "repo", DefValue: ".miner", Description: "Repo path"},
		{Name: "host-multiaddr", DefValue: "/ip4/0.0.0.0/tcp/4001", Description: "Libp2p host listen multiaddr"},
		{Name: "metrics-addr", DefValue: ":9090", Description: "Prometheus listen address"},
	}

	common.ConfigureCLI(v, "MINER", flags, rootCmd)
}

var rootCmd = &cobra.Command{
	Use:   daemonName,
	Short: "minerd is used by a miner to listen for deals from the Broker",
	Long:  "minerd is used by a miner to listen for deals from the Broker",
	PersistentPreRun: func(c *cobra.Command, args []string) {
		logging.SetAllLoggers(logging.LevelInfo)
		if v.GetBool("debug") {
			logging.SetAllLoggers(logging.LevelDebug)
		}
	},
	Run: func(c *cobra.Command, args []string) {
		settings, err := json.MarshalIndent(v.AllSettings(), "", "  ")
		common.CheckErr(err)
		log.Infof("loaded config: %s", string(settings))

		if err := common.SetupInstrumentation(v.GetString("metrics.addr")); err != nil {
			log.Fatalf("booting instrumentation: %s", err)
		}

		config := service.Config{
			RepoPath: v.GetString("repo"),
			Peer: marketpeer.Config{
				RepoPath:      v.GetString("repo"),
				HostMultiaddr: v.GetString("host-multiaddr"),
			},
			BidParams: service.BidParams{
				AskPrice: 100000000000,
			},
			AuctionFilters: service.AuctionFilters{
				DealDuration: service.MinMaxFilter{
					Min: broker.MinDealEpochs,
					Max: broker.MaxDealEpochs,
				},
				DealSize: service.MinMaxFilter{
					Min: 56 * 1024,               // 56KiB
					Max: 32 * 1000 * 1000 * 1000, // 32GB
				},
			},
		}
		serv, err := service.New(config)
		common.CheckErr(err)

		serv.Bootstrap()

		common.HandleInterrupt(func() {
			if err := serv.Close(); err != nil {
				log.Errorf("closing service: %s", err)
			}
		})
	},
}

func main() {
	common.CheckErr(rootCmd.Execute())
}

package main

import (
	"encoding/hex"
	"fmt"
	_ "net/http/pprof"
	"os"
	"path/filepath"

	golog "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	mbase "github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/textileio/broker-core/broker"
	"github.com/textileio/broker-core/chain"
	"github.com/textileio/broker-core/cmd/bidbot/service"
	"github.com/textileio/broker-core/cmd/common"
	"github.com/textileio/broker-core/finalizer"
	"github.com/textileio/broker-core/marketpeer"
)

var (
	cliName           = "bidbot"
	defaultConfigPath = filepath.Join(os.Getenv("HOME"), "."+cliName)
	log               = golog.Logger(cliName)
	v                 = viper.New()
)

func init() {
	rootCmd.AddCommand(initCmd, daemonCmd)

	flags := []common.Flag{
		{
			Name:        "ask-price",
			DefValue:    100000000000,
			Description: "Bid ask price for deals in attoFIL per GiB per epoch; default is 100 nanoFIL",
		},
		{
			Name:        "verified-ask-price",
			DefValue:    100000000000,
			Description: "Bid ask price for verified deals in attoFIL per GiB per epoch; default is 100 nanoFIL",
		},
		{
			Name:        "deal-duration-min",
			DefValue:    broker.MinDealEpochs,
			Description: "Minimum deal duration to bid on in epochs; default is ~6 months",
		},
		{
			Name:        "deal-duration-max",
			DefValue:    broker.MaxDealEpochs,
			Description: "Maximum deal duration to bid on in epochs; default is ~1 year",
		},
		{
			Name:        "deal-size-min",
			DefValue:    56 * 1024,
			Description: "Minimum deal size to bid on in bytes",
		},
		{
			Name:        "deal-size-max",
			DefValue:    32 * 1000 * 1000 * 1000,
			Description: "Maximum deal size to bid on in bytes",
		},
		{
			Name:        "deal-start-window",
			DefValue:    60 * 24 * 2,
			Description: "Number of epochs after which won deals must start on-chain; default is ~one day",
		},
		{
			Name:        "fast-retrieval",
			DefValue:    false,
			Description: "Offer deals with fast retrieval",
		},
		{Name: "metrics-addr", DefValue: ":9090", Description: "Prometheus listen address"},
		{Name: "log-debug", DefValue: false, Description: "Enable debug level log"},
		{Name: "log-json", DefValue: false, Description: "Enable structured logging"},
	}
	flags = append(flags, marketpeer.Flags...)

	cobra.OnInitialize(func() {
		v.SetConfigType("json")
		v.SetConfigName("config")
		v.AddConfigPath(os.Getenv("BIDBOT_PATH"))
		v.AddConfigPath(defaultConfigPath)
		if err := v.ReadInConfig(); err != nil {
			common.CheckErrf("reading configuration: %s", err)
		}
	})

	common.ConfigureCLI(v, "BIDBOT", flags, rootCmd.PersistentFlags())
}

var rootCmd = &cobra.Command{
	Use:   cliName,
	Short: "Bidbot listens for Filecoin storage deal auctions from deal brokers",
	Long: `Bidbot listens for Filecoin storage deal auctions from deal brokers.

bidbot will automatically bid on storage deals that pass configured filters at
the configured prices.

To get started, run 'bidbot init' and follow the instructions. 
`,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes bidbot configuration file",
	Long: `Initializes bidbot configuration file and generates a new keypair.

bidbot uses a repository in the local file system. By default, the repo is
located at ~/.bidbot. To change the repo location, set the $BIDBOT_PATH
environment variable:

    export BIDBOT_PATH=/path/to/bidbotrepo
`,
	Run: func(c *cobra.Command, args []string) {
		path, err := marketpeer.WriteConfig(v, "BIDBOT_PATH", defaultConfigPath)
		common.CheckErrf("writing config: %v", err)
		fmt.Printf("Initialized configuration file: %s\n\n", path)

		_, key, err := mbase.Decode(v.GetString("private-key"))
		common.CheckErrf("decoding private key: %v", err)
		priv, err := crypto.UnmarshalPrivateKey(key)
		common.CheckErrf("unmarshaling private key: %v", err)
		id, err := peer.IDFromPrivateKey(priv)
		common.CheckErrf("getting peer id: %v", err)

		signingToken := hex.EncodeToString([]byte(id))

		fmt.Printf(`Bidbot needs a signature from a miner wallet address to authenticate bids.

1. Sign this token with an address from your Lotus wallet:

    lotus wallet sign [address] %s

2. Start listening for deal auctions using the wallet address and signature from step 1:

    bidbot daemon [address] [signature]

Note: In the event you win an auction, you must use this wallet address to make the deal(s).

Good luck!
`, signingToken)
	},
}

var daemonCmd = &cobra.Command{
	Use:   "daemon [address] [signature]",
	Short: "Run a network-connected bidding bot",
	Long:  "Run a network-connected bidding bot that listens for and bids on storage deal auctions.",
	Args:  cobra.ExactArgs(2),
	PersistentPreRun: func(c *cobra.Command, args []string) {
		common.ExpandEnvVars(v, v.AllSettings())
		err := common.ConfigureLogging(v, []string{
			"bidbot",
			"bidbot/service",
			"mpeer",
			"pubsub",
		})
		common.CheckErrf("setting log levels: %v", err)
	},
	Run: func(c *cobra.Command, args []string) {
		if v.ConfigFileUsed() == "" {
			fmt.Printf("Configuration file not found. Run '%s init'.", cliName)
			return
		}

		fin := finalizer.NewFinalizer()

		settings, err := marketpeer.MarshalConfig(v)
		common.CheckErrf("marshaling config: %v", err)
		log.Infof("loaded config: %s", string(settings))

		err = common.SetupInstrumentation(v.GetString("metrics.addr"))
		common.CheckErrf("booting instrumentation: %v", err)

		pconfig, err := marketpeer.GetConfig(v, false)
		common.CheckErrf("getting peer config: %v", err)

		walletAddr := args[0]
		walletAddrSig, err := hex.DecodeString(args[1])
		common.CheckErrf("decoding wallet address signature: %v", err)

		ch, err := chain.New()
		common.CheckErrf("creating chain client: %v", err)
		fin.Add(ch)

		config := service.Config{
			RepoPath: pconfig.RepoPath,
			Peer:     pconfig,
			BidParams: service.BidParams{
				WalletAddr:       walletAddr,
				WalletAddrSig:    walletAddrSig,
				AskPrice:         v.GetInt64("ask-price"),
				VerifiedAskPrice: v.GetInt64("verified-ask-price"),
				FastRetrieval:    v.GetBool("fast-retrieval"),
				DealStartWindow:  v.GetUint64("deal-start-window"),
			},
			AuctionFilters: service.AuctionFilters{
				DealDuration: service.MinMaxFilter{
					Min: v.GetUint64("deal-duration-min"),
					Max: v.GetUint64("deal-duration-max"),
				},
				DealSize: service.MinMaxFilter{
					Min: v.GetUint64("deal-size-min"),
					Max: v.GetUint64("deal-size-max"),
				},
			},
		}
		serv, err := service.New(config, ch)
		common.CheckErrf("starting service: %v", err)
		fin.Add(serv)

		err = serv.Subscribe(true)
		common.CheckErrf("subscribing to deal auction feed: %v", err)

		common.HandleInterrupt(func() {
			common.CheckErr(fin.Cleanupf("closing service: %v", nil))
		})
	},
}

func main() {
	common.CheckErr(rootCmd.Execute())
}

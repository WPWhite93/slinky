package mexc

import (
	"time"

	"github.com/skip-mev/slinky/oracle/config"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
)

const (
	// Please refer to the following link for the MEXC websocket documentation:
	// https://mexcdevelop.github.io/apidocs/spot_v3_en/#websocket-market-streams.

	// Name is the name of the MEXC provider.
	Name = "mexc"

	// WSS is the public MEXC Websocket URL.
	WSS = "wss://wbs.mexc.com/ws"

	// DefaultPingInterval is the default ping interval for the MEXC websocket. The documentation
	// specifies that this should be done every 30 seconds, however, the actual threshold should be
	// slightly lower than this to account for network latency.
	DefaultPingInterval = 20 * time.Second

	// MaxSubscriptionsPerConnection is the maximum number of subscriptions that can be made
	// per connection.
	//
	// ref: https://mexcdevelop.github.io/apidocs/spot_v3_en/#websocket-market-streams
	MaxSubscriptionsPerConnection = 30
)

var (
	// DefaultWebSocketConfig is the default configuration for the MEXC Websocket.
	DefaultWebSocketConfig = config.WebSocketConfig{
		Name:                          Name,
		Enabled:                       true,
		MaxBufferSize:                 1000,
		ReconnectionTimeout:           config.DefaultReconnectionTimeout,
		WSS:                           WSS,
		ReadBufferSize:                config.DefaultReadBufferSize,
		WriteBufferSize:               config.DefaultWriteBufferSize,
		HandshakeTimeout:              config.DefaultHandshakeTimeout,
		EnableCompression:             config.DefaultEnableCompression,
		ReadTimeout:                   config.DefaultReadTimeout,
		WriteTimeout:                  config.DefaultWriteTimeout,
		PingInterval:                  DefaultPingInterval,
		MaxReadErrorCount:             config.DefaultMaxReadErrorCount,
		MaxSubscriptionsPerConnection: config.DefaultMaxSubscriptionsPerConnection,
	}

	// DefaultMarketConfig is the default market configuration for the MEXC Websocket.
	DefaultMarketConfig = config.MarketConfig{
		Name: Name,
		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
			"ATOM/USDC": {
				Ticker:       "ATOMUSDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("ATOM", "USDC"),
			},
			"ATOM/USDT": {
				Ticker:       "ATOMUSDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("ATOM", "USDT"),
			},
			"AVAX/USDC": {
				Ticker:       "AVAXUSDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("AVAX", "USDC"),
			},
			"AVAX/USDT": {
				Ticker:       "AVAXUSDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("AVAX", "USDT"),
			},
			"BITCOIN/USDC": {
				Ticker:       "BTCUSDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("BITCOIN", "USDC"),
			},
			"BITCOIN/USDT": {
				Ticker:       "BTCUSDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("BITCOIN", "USDT"),
			},
			"ETHEREUM/BITCOIN": {
				Ticker:       "ETHBTC",
				CurrencyPair: slinkytypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
			},
			"ETHEREUM/USDC": {
				Ticker:       "ETHUSDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("ETHEREUM", "USDC"),
			},
			"ETHEREUM/USDT": {
				Ticker:       "ETHUSDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("ETHEREUM", "USDT"),
			},
			"SOLANA/USDC": {
				Ticker:       "SOLUSDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("SOLANA", "USDC"),
			},
			"SOLANA/USDT": {
				Ticker:       "SOLUSDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("SOLANA", "USDT"),
			},
			"USDC/USDT": {
				Ticker:       "USDCUSDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("USDC", "USDT"),
			},
		},
	}
)

package kucoin

import (
	"time"

	"github.com/skip-mev/slinky/oracle/config"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
)

const (
	// Name is the name of the KuCoin provider.
	Name = "kucoin"

	// WSSEndpoint contains the endpoint format for Kucoin websocket API. Specifically
	// this inputs the dynamically generated token from the user and the endpoint.
	WSSEndpoint = "%s?token=%s"

	// WSS is the websocket URL for Kucoin. Note that this may change as the URL is
	// dynamically generated. A token is required to connect to the websocket feed.
	WSS = "wss://ws-api-spot.kucoin.com/"

	// URL is the Kucoin websocket URL. This URL specifically points to the public
	// spot and maring REST API.
	URL = "https://api.kucoin.com"

	// DefaultPingInterval is the default ping interval for the KuCoin websocket.
	DefaultPingInterval = 10 * time.Second
)

var (
	// DefaultWebSocketConfig defines the default websocket config for Kucoin.
	DefaultWebSocketConfig = config.WebSocketConfig{
		Enabled:                       true,
		MaxBufferSize:                 config.DefaultMaxBufferSize,
		ReconnectionTimeout:           config.DefaultReconnectionTimeout,
		WSS:                           WSS, // Note that this may change as the URL is dynamically generated.
		Name:                          Name,
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

	// DefaultAPIConfig defines the default API config for KuCoin. This is
	// only utilized on the initial connection to the websocket feed.
	DefaultAPIConfig = config.APIConfig{
		Enabled:    false,
		Timeout:    5 * time.Second, // KuCoin recommends a timeout of 5 seconds.
		Interval:   1 * time.Minute, // This is not used.
		MaxQueries: 1,               // This is not used.
		URL:        URL,
		Name:       Name,
	}

	// DefaultMarketConfig defines the default market config for Kucoin.
	DefaultMarketConfig = config.MarketConfig{
		Name: Name,
		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
			"ATOM/USDC": {
				Ticker:       "ATOM-USDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("ATOM", "USDC"),
			},
			"ATOM/USDT": {
				Ticker:       "ATOM-USDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("ATOM", "USDT"),
			},
			"AVAX/USDC": {
				Ticker:       "AVAX-USDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("AVAX", "USDC"),
			},
			"AVAX/USDT": {
				Ticker:       "AVAX-USDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("AVAX", "USDT"),
			},
			"BITCOIN/USDC": {
				Ticker:       "BTC-USDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("BITCOIN", "USDC"),
			},
			"BITCOIN/USDT": {
				Ticker:       "BTC-USDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("BITCOIN", "USDT"),
			},
			"CELESTIA/USDT": {
				Ticker:       "TIA-USDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("CELESTIA", "USDT"),
			},
			"DYDX/USDT": {
				Ticker:       "DYDX-USDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("DYDX", "USDT"),
			},
			"ETHEREUM/BITCOIN": {
				Ticker:       "ETH-BTC",
				CurrencyPair: slinkytypes.NewCurrencyPair("ETHEREUM", "BITCOIN"),
			},
			"ETHEREUM/USDC": {
				Ticker:       "ETH-USDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("ETHEREUM", "USDC"),
			},
			"ETHEREUM/USDT": {
				Ticker:       "ETH-USDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("ETHEREUM", "USDT"),
			},
			"OSMOSIS/USDT": {
				Ticker:       "OSMO-USDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("OSMOSIS", "USDT"),
			},
			"SOLANA/USDC": {
				Ticker:       "SOL-USDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("SOLANA", "USDC"),
			},
			"SOLANA/USDT": {
				Ticker:       "SOL-USDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("SOLANA", "USDT"),
			},
			"USDC/USDT": {
				Ticker:       "USDC-USDT",
				CurrencyPair: slinkytypes.NewCurrencyPair("USDC", "USDT"),
			},
		},
	}
)

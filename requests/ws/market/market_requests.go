package market

import "github.com/ggquant/okex"

type (
	Tickers struct {
		InstID string `json:"instId"`
	}
	Candlesticks struct {
		InstID  string                    `json:"instId"`
		Channel okex.CandleStickWsBarSize `json:"channel"`
	}
	Trades struct {
		InstID string `json:"instId"`
	}
	OrderBook struct {
		InstID  string `json:"instId"`
		Channel string `json:"channel"`
	}
)

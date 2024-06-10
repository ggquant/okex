package market

import (
	"github.com/amir-the-h/okex/events"
	"github.com/amir-the-h/okex/models/market"
)

type (
	Tickers struct {
		Arg     *events.Argument `json:"arg"`
		Tickers []*market.Ticker `json:"data"`
	}
	Candlesticks struct {
		Arg     *events.Argument `json:"arg"`
		Candles []*market.Candle `json:"data"`
	}
	Trades struct {
		Arg    *events.Argument `json:"arg"`
		Trades []*market.Trade  `json:"data"`
	}
	OrderBook struct {
		Arg    *events.Argument      `json:"arg"`
		InstID string                `json:"instId"`
		Books  []*market.OrderBookWs `json:"data"`
		Action string                `json:"action"`
	}
)

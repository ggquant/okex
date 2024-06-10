package public

import (
	"github.com/ggquant/okex/events"
	"github.com/ggquant/okex/models/market"
	"github.com/ggquant/okex/models/publicdata"
)

type (
	Instruments struct {
		Arg         *events.Argument         `json:"arg"`
		Instruments []*publicdata.Instrument `json:"data"`
	}
	OpenInterest struct {
		Arg           *events.Argument           `json:"arg"`
		OpenInterests []*publicdata.OpenInterest `json:"data"`
	}
	EstimatedDeliveryExercisePrice struct {
		Arg                             *events.Argument                             `json:"arg"`
		EstimatedDeliveryExercisePrices []*publicdata.EstimatedDeliveryExercisePrice `json:"data"`
	}
	MarkPrice struct {
		Arg    *events.Argument        `json:"arg"`
		Prices []*publicdata.MarkPrice `json:"data"`
	}
	MarkPriceCandlesticks struct {
		Arg    *events.Argument      `json:"arg"`
		Prices []*market.IndexCandle `json:"data"`
	}
	PriceLimit struct {
		Arg   *events.Argument         `json:"arg"`
		Limit []*publicdata.LimitPrice `json:"data"`
	}
	OPTIONSummary struct {
		Arg     *events.Argument               `json:"arg"`
		Options []*publicdata.OptionMarketData `json:"data"`
	}
	FundingRate struct {
		Arg   *events.Argument          `json:"arg"`
		Rates []*publicdata.FundingRate `json:"data"`
	}
	IndexCandlesticks struct {
		Arg   *events.Argument      `json:"arg"`
		Rates []*market.IndexCandle `json:"data"`
	}
	IndexTickers struct {
		Arg     *events.Argument      `json:"arg"`
		Tickers []*market.IndexTicker `json:"data"`
	}
)

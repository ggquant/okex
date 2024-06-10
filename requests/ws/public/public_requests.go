package public

import "github.com/ggquant/okex"

type (
	Instruments struct {
		InstType okex.InstrumentType `json:"instType"`
	}
	OpenInterest struct {
		InstID string `json:"instId"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly,omitempty"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
	}
	MarkPrice struct {
		InstID string `json:"instId"`
	}
	MarkPriceCandlesticks struct {
		InstID  string                    `json:"instId"`
		Channel okex.CandleStickWsBarSize `json:"channel"`
	}
	PriceLimit struct {
		InstID string `json:"instId"`
	}
	OPTIONSummary struct {
		InstID string `json:"instId"`
		Uly    string `json:"uly"`
	}
	FundingRate struct {
		InstID string `json:"instId"`
	}
	IndexCandlesticks struct {
		InstID  string `json:"instId"`
		Channel string `json:"channel"`
	}
	IndexTickers struct {
		InstID string `json:"instId"`
	}
	LiquidationOrders struct {
		InstType string                    `json:"instType"`
		Channel  okex.CandleStickWsBarSize `json:"channel"`
	}
)

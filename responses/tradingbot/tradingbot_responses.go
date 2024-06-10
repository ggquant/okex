package tradingbot

import (
	"github.com/amir-the-h/okex/models/trade"
	"github.com/amir-the-h/okex/responses"
)

type (
	PlaceAlgoOrder struct {
		responses.Basic
		PlaceAlgoOrders []*trade.PlaceAlgoOrder `json:"data"`
	}
)

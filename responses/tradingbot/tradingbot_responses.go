package tradingbot

import (
	"github.com/ggquant/okex/models/trade"
	"github.com/ggquant/okex/responses"
)

type (
	PlaceAlgoOrder struct {
		responses.Basic
		PlaceAlgoOrders []*trade.PlaceAlgoOrder `json:"data"`
	}
)

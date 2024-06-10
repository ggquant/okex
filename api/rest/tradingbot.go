package rest

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/rest/tradingbot"
	responses "github.com/amir-the-h/okex/responses/tradingbot"
	"net/http"
)

// TradingBot
//
// https://www.okx.com/docs-v5/zh/#order-book-trading-grid-trading
type TradingBot struct {
	client *ClientRest
}

// NewTradingBot returns a pointer to a fresh Account
func NewTradingBot(c *ClientRest) *TradingBot {
	return &TradingBot{c}
}

// PlaceAlgoOrder
// Retrieve a list of instruments with open contracts.
//
// https://www.okx.com/docs-v5/zh/#order-book-trading-grid-trading-post-place-grid-algo-order
func (c *TradingBot) PlaceAlgoOrder(req requests.PlaceAlgoOrder) (response responses.PlaceAlgoOrder, err error) {
	p := "/api/v5/tradingBot/grid/order-algo"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

package rest

import (
	"encoding/json"
	"github.com/ggquant/okex"
	requests "github.com/ggquant/okex/requests/rest/copytrading"
	responses "github.com/ggquant/okex/responses/copytrading"
	"net/http"
)

// CopyTrading
//
// https://www.okx.com/docs-v5/zh/#order-book-trading-copy-trading
type CopyTrading struct {
	client *ClientRest
}

// NewCopyTrading returns a pointer to a fresh Account
func NewCopyTrading(c *ClientRest) *Account {
	return &Account{c}
}

// GetCurrentSubpositions
// Retrieve a list of instruments with open contracts.
//
// https://www.okx.com/docs-v5/zh/#order-book-trading-copy-trading-get-existing-lead-or-copy-positions
func (c *CopyTrading) GetCurrentSubpositions(req requests.GetCurrentSubPositions) (response responses.GetCurrentSubPositions, err error) {
	p := "/api/v5/copytrading/current-subpositions"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

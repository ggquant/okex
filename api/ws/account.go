package ws

import (
	"encoding/json"
	"github.com/ggquant/okex"
	"github.com/ggquant/okex/events"
	"github.com/ggquant/okex/events/account"
	requests "github.com/ggquant/okex/requests/ws/private"
)

// Account
//
// https://www.okx.com/docs-v5/zh/#trading-account-websocket
type Account struct {
	*ClientWs
	aCh   chan *account.Account
	pCh   chan *account.Position
	bnpCh chan *account.BalanceAndPosition
	oCh   chan *account.Order
}

// NewAccount returns a pointer to a fresh Private
func NewAccount(c *ClientWs) *Account {
	return &Account{ClientWs: c}
}

// Account
// Retrieve account information. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-account-channel
func (c *Account) Account(req requests.Account, ch ...chan *account.Account) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.aCh = ch[0]
	}
	return c.Subscribe(true, []okex.ChannelName{"account"}, m)
}

// UAccount
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-account-channel
func (c *Account) UAccount(req requests.Account, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.aCh = nil
	}
	return c.Unsubscribe(true, []okex.ChannelName{"account"}, m)
}

// Position
// Retrieve position information. Initial snapshot will be pushed according to subscription granularity. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-positions-channel
func (c *Account) Position(req requests.Position, ch ...chan *account.Position) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.pCh = ch[0]
	}
	return c.Subscribe(true, []okex.ChannelName{"positions"}, m)
}

// UPosition
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-positions-channel
func (c *Account) UPosition(req requests.Position, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.pCh = nil
	}
	return c.Unsubscribe(true, []okex.ChannelName{"positions"}, m)
}

// BalanceAndPosition
// Retrieve account balance and position information. Data will be pushed when triggered by events such as filled order, funding transfer.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
func (c *Account) BalanceAndPosition(ch ...chan *account.BalanceAndPosition) error {
	m := make(map[string]string)
	if len(ch) > 0 {
		c.bnpCh = ch[0]
	}
	return c.Subscribe(true, []okex.ChannelName{"balance_and_position"}, m)
}

// UBalanceAndPosition unsubscribes a position channel
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
func (c *Account) UBalanceAndPosition(rCh ...bool) error {
	m := make(map[string]string)
	if len(rCh) > 0 && rCh[0] {
		c.bnpCh = nil
	}
	return c.Unsubscribe(true, []okex.ChannelName{"balance_and_position"}, m)
}

// Order
// Retrieve position information. Initial snapshot will be pushed according to subscription granularity. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-order-channel
func (c *Account) Order(req requests.Order, ch ...chan *account.Order) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.oCh = ch[0]
	}
	return c.Subscribe(true, []okex.ChannelName{"orders"}, m)
}

// UOrder
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-order-channel
func (c *Account) UOrder(req requests.Order, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.oCh = nil
	}
	return c.Unsubscribe(true, []okex.ChannelName{"orders"}, m)
}

func (c *Account) Process(data []byte, e *events.Basic) bool {
	if e.Event == "" && e.Arg != nil && e.Data != nil && len(e.Data) > 0 {
		ch, ok := e.Arg.Get("channel")
		if !ok {
			return false
		}
		switch ch {
		case "account":
			e := account.Account{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.aCh != nil {
					c.aCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		case "positions":
			e := account.Position{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.pCh != nil {
					c.pCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		case "balance_and_position":
			e := account.BalanceAndPosition{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.bnpCh != nil {
					c.bnpCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		case "orders":
			e := account.Order{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.oCh != nil {
					c.oCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		}
	}
	return false
}

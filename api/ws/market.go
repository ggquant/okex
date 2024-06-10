package ws

import (
	"encoding/json"
	"fmt"
	"github.com/ggquant/okex"
	"github.com/ggquant/okex/events"
	"github.com/ggquant/okex/events/market"
	requests "github.com/ggquant/okex/requests/ws/market"
	"strings"
)

// Market
//
// https://www.okx.com/docs-v5/zh/#order-book-trading-market-data
type Market struct {
	*ClientWs
	tCh  chan *market.Tickers
	cCh  chan *market.Candlesticks
	trCh chan *market.Trades
	obCh chan *market.OrderBook
}

// NewMarket returns a pointer to a fresh Public
func NewMarket(c *ClientWs) *Market {
	return &Market{ClientWs: c}
}

// Tickers
// Retrieve the last traded price, bid price, ask price and 24-hour trading volume of instruments. Data will be pushed every 100 ms.
//
// https://www.okx.com/docs-v5/zh/#order-book-trading-market-data-ws-tickers-channel
func (c *Market) Tickers(req requests.Tickers, ch ...chan *market.Tickers) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.tCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"tickers"}, m)
}

// UTickers
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-tickers-channel
func (c *Market) UTickers(req requests.Tickers, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.tCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"tickers"}, m)
}

// Candlesticks
// Retrieve the open interest. Data will by pushed every 3 seconds.
//
// https://www.okx.com/docs-v5/zh/#order-book-trading-market-data-ws-candlesticks-channel
func (c *Market) Candlesticks(req requests.Candlesticks, ch ...chan *market.Candlesticks) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.cCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{}, m)
}

// UCandlesticks
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (c *Market) UCandlesticks(req requests.Candlesticks, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.cCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{}, m)
}

// Trades
// Retrieve the recent trades data. Data will be pushed whenever there is a trade.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-trades-channel
func (c *Market) Trades(req requests.Trades, ch ...chan *market.Trades) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.trCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"trades"}, m)
}

// UTrades
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-trades-channel
func (c *Market) UTrades(req requests.Trades, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.trCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"trades"}, m)
}

// TradesAll
// Retrieve the recent trades data. Data will be pushed whenever there is a trade.
//
// https://www.okx.com/docs-v5/zh/#order-book-trading-market-data-ws-all-trades-channel
func (c *Market) TradesAll(req requests.Trades, ch ...chan *market.Trades) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.trCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{"trades-all"}, m)
}

// UTrades
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-trades-channel
func (c *Market) UTradesAll(req requests.Trades, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.trCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{"trades"}, m)
}

// OrderBook
// Retrieve order book data.
//
// Use books for 400 depth levels, book5 for 5 depth levels, books50-l2-tbt tick-by-tick 50 depth levels, and books-l2-tbt for tick-by-tick 400 depth levels.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-order-book-channel
func (c *Market) OrderBook(req requests.OrderBook, ch ...chan *market.OrderBook) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.obCh = ch[0]
	}
	return c.Subscribe(false, []okex.ChannelName{}, m)
}

// UOrderBook
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-order-book-channel
func (c *Market) UOrderBook(req requests.OrderBook, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.obCh = nil
	}
	return c.Unsubscribe(false, []okex.ChannelName{okex.ChannelName(req.Channel)}, m)
}

func (c *Market) Process(data []byte, e *events.Basic) bool {
	if e.Event == "" && e.Arg != nil && e.Data != nil && len(e.Data) > 0 {
		ch, ok := e.Arg.Get("channel")
		if !ok {
			return false
		}
		switch ch {
		case "tickers":
			e := market.Tickers{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.tCh != nil {
					c.tCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		case "trades":
			e := market.Trades{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.trCh != nil {
					c.trCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		default:
			// special cases
			// market price candlestick channel
			chName := fmt.Sprint(ch)
			// candlestick channels
			if strings.Contains(chName, "candle") {
				e := market.Candlesticks{}
				err := json.Unmarshal(data, &e)
				if err != nil {
					return false
				}
				go func() {
					if c.cCh != nil {
						c.cCh <- &e
					}
					c.StructuredEventChan <- e
				}()
				return true
			}
			// order book channels
			if strings.Contains(chName, "books") {
				e := market.OrderBook{}
				err := json.Unmarshal(data, &e)
				if err != nil {
					return false
				}
				go func() {
					if c.obCh != nil {
						c.obCh <- &e
					}
					c.StructuredEventChan <- e
				}()
				return true
			}
		}
	}
	return false
}

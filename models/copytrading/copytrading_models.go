package copytrading

import "github.com/ggquant/okex"

type (
	SubPosition struct {
		InstID      string            `json:"instId"`
		SubPosId    string            `json:"subPosId"`
		PosSide     okex.PositionSide `json:"posSide"`
		MgnMode     okex.MarginMode   `json:"mgnMode"`
		Lever       string            `json:"lever"`
		OpenOrdId   string            `json:"openOrdId"`
		OpenAvgPx   string            `json:"openAvgPx"`
		OpenTime    okex.JSONTime     `json:"openTime"`
		SubPos      string            `json:"subPos"`
		TpTriggerPx string            `json:"tpTriggerPx"`
		SlTriggerPx string            `json:"slTriggerPx"`
		AlgoId      string            `json:"algoId"`
		InstType    string            `json:"instType"`
		TpOrdPx     string            `json:"tpOrdPx"`
		SlOrdPx     string            `json:"slOrdPx"`
		Margin      string            `json:"margin"`
		Upl         string            `json:"upl"`
		UplRatio    string            `json:"uplRatio"`
		MarkPx      string            `json:"markPx"`
		UniqueCode  string            `json:"uniqueCode"`
		Ccy         string            `json:"ccy"`
		AvailSubPos string            `json:"availSubPos"`
	}
	LeadTrader struct {
		PortLink           string           `json:"portLink"`
		NickName           string           `json:"nickName"`
		Margin             okex.JSONFloat64 `json:"margin"`
		CopyTotalAmt       okex.JSONFloat64 `json:"copyTotalAmt"`
		CopyTotalPnl       okex.JSONFloat64 `json:"copyTotalPnl"`
		UniqueCode         string           `json:"uniqueCode"`
		Ccy                string           `json:"ccy"`
		ProfitSharingRatio okex.JSONFloat64 `json:"profitSharingRatio"`
		BeginCopyTime      okex.JSONTime    `json:"beginCopyTime"`
		Upl                okex.JSONFloat64 `json:"upl"`
		TodayPnl           okex.JSONFloat64 `json:"todayPnl"`
		LeadMode           okex.LeadMode    `json:"leadMode"`
	}
)

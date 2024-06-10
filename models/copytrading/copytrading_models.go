package copytrading

import "github.com/amir-the-h/okex"

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
)

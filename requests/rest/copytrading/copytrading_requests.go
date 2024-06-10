package copytrading

import "github.com/ggquant/okex"

type (
	GetCurrentSubPositions struct {
		InstType   okex.InstrumentType `json:"instType"`
		InstID     string              `json:"instId,omitempty"`
		UniqueCode string              `json:"uniqueCode,omitempty"`
		SubPosType okex.SubPosType     `json:"subPosType,omitempty"`
		After      string              `json:"after,omitempty"`
		Before     string              `json:"before,omitempty"`
		Limit      string              `json:"limit,omitempty"`
	}
)

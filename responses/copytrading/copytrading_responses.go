package copytrading

import (
	models "github.com/ggquant/okex/models/copytrading"
	"github.com/ggquant/okex/responses"
)

type (
	GetCurrentSubPositions struct {
		responses.Basic
		SubPositions []*models.SubPosition `json:"data,omitempty"`
	}
)

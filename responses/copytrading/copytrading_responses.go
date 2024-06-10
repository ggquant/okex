package copytrading

import (
	models "github.com/amir-the-h/okex/models/copytrading"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetCurrentSubPositions struct {
		responses.Basic
		SubPositions []*models.SubPosition `json:"data,omitempty"`
	}
)

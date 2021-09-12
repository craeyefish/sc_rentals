package items

import (
	"time"

	"github.com/shopspring/decimal"
)

// Item defines the qualities of the rental objects.
type Item struct {
	Name     string
	Code     int64
	Type     Type
	Value    decimal.Decimal
	InitDate time.Time
}

// Type of rental object.
type Type int

const (
	TypeUnknown  Type = 0
	TypeChair    Type = 1
	typeSentinel Type = 2
)

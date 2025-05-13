package order

import (
	"time"

	"github.com/EdwardLuthor/pizzaman-terminal/pizza"
)

type OrderState string

const (
	OrderStateNew        OrderState = "new"
	OrderStateInProgress OrderState = "in_progress"
	OrderStateDone       OrderState = "done"
)

type Order struct {
	Items     []pizza.Pizza
	TotalCost float64
	State     OrderState
	CreatedAt time.Time
	UpdatedAt time.Time
}

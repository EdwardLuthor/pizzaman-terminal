package order

import "time"

type OrderState string

const (
	OrderStateNew        OrderState = "new"
	OrderStateInProgress OrderState = "in_progress"
	OrderStateDone       OrderState = "done"
)

type Order struct {
	Items     []pizza.pizza
	TotalCost float64
	State     OrderState
	CreatedAt time.Time
	UpdatedAt time.Time
}

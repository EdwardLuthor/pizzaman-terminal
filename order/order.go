package order

type OrderState string

const (
	OrderStateNew        OrderState = "new"
	OrderStateInProgress OrderState = "in_progress"
	OrderStateDone       OrderState = "done"
)

type Order map[string]interface{}

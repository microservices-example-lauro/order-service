package services

type PlaceOrderInput struct {
	UserId    string
	ProductId string
	Quantity  int
}

type PlaceOrderOutput struct {
	OrderId string
	Value   float64
}

type GetOrderInput struct {
	OrderId string
}

type GetOrderOutput struct {
	OrderId   string
	UserId    string
	ProductId string
	Quantity  int
	Value     float64
}

type OrderService interface {
	PlaceOrder(input PlaceOrderInput) (PlaceOrderOutput, error)
	GetOrder(input GetOrderInput) (GetOrderOutput, error)
}

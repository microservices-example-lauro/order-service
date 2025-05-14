package repositories

type Order struct {
	OrderId   string
	UserId    string
	ProductId string
	Quantity  int
	Value     float64
}

type OrderRepository interface {
	CreateOrder(order Order) error
	GetOrder(orderId string) (Order, error)
}

package repositories

import "errors"

const FailedToGetOrderFromDatabaseError = "failed to get order from database"

type FakeOrderRepository struct {
	orders []Order
}

func (r *FakeOrderRepository) CreateOrder(order Order) error {
	r.orders = append(r.orders, order)
	return nil
}

func (r *FakeOrderRepository) GetOrder(orderId string) (Order, error) {
	for _, order := range r.orders {
		if order.OrderId == orderId {
			return order, nil
		}
	}
	return Order{}, errors.New(FailedToGetOrderFromDatabaseError)
}

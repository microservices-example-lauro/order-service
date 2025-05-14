package services

import (
	"errors"
	"microservices-example-lauro/order-service/internal/repositories"
	"regexp"

	"github.com/google/uuid"
)

const InvalidUserIdError = "invalid user id"
const InvalidProductIdError = "invalid product id"
const InvalidQuantityError = "invalid quantity"

type OrderServiceImpl struct {
	orderRepository repositories.OrderRepository
}

func isUUIDv4(s string) bool {
	var uuidv4Regex = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return uuidv4Regex.MatchString(s)
}

func generateUUIDv4() string {
	return uuid.New().String()
}

func (o *OrderServiceImpl) PlaceOrder(input PlaceOrderInput) (PlaceOrderOutput, error) {
	if !isUUIDv4(input.UserId) {
		return PlaceOrderOutput{}, errors.New(InvalidUserIdError)
	}

	if !isUUIDv4(input.ProductId) {
		return PlaceOrderOutput{}, errors.New(InvalidProductIdError)
	}

	if input.Quantity <= 0 {
		return PlaceOrderOutput{}, errors.New(InvalidQuantityError)
	}

	orderId := generateUUIDv4()
	orderValue := float64(5 * input.Quantity)
	order := repositories.Order{OrderId: orderId, UserId: input.UserId, ProductId: input.ProductId, Quantity: input.Quantity, Value: orderValue}
	createOrderErr := o.orderRepository.CreateOrder(order)
	if createOrderErr != nil {
		return PlaceOrderOutput{}, createOrderErr
	}

	return PlaceOrderOutput{OrderId: orderId, Value: orderValue}, nil
}

func (o *OrderServiceImpl) GetOrder(input GetOrderInput) (GetOrderOutput, error) {
	order, getOrderErr := o.orderRepository.GetOrder(input.OrderId)

	if getOrderErr != nil {
		return GetOrderOutput{}, getOrderErr
	}

	return GetOrderOutput{OrderId: order.OrderId, UserId: order.UserId, ProductId: order.ProductId, Quantity: order.Quantity, Value: order.Value}, nil
}

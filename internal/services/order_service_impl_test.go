package services

import (
	"errors"
	"microservices-example-lauro/order-service/internal/repositories"
	"testing"
)

func TestPlaceOrder(t *testing.T) {
	tests := map[string]struct {
		userId         string
		productId      string
		quantity       int
		expectedOutput PlaceOrderOutput
		expectedErr    error
	}{
		"valid_order": {
			userId:         "91d27161-8b36-46de-b5eb-9fbe6af2356c",
			productId:      "073fa958-fe19-41bc-ae99-0974c9cfe9f7",
			quantity:       1,
			expectedOutput: PlaceOrderOutput{OrderId: "", Value: 5},
			expectedErr:    nil,
		},
		"invalid_order_user_id": {
			userId:         "",
			productId:      "073fa958-fe19-41bc-ae99-0974c9cfe9f7",
			quantity:       1,
			expectedOutput: PlaceOrderOutput{},
			expectedErr:    errors.New(InvalidUserIdError),
		},
		"invalid_order_product_id": {
			userId:         "91d27161-8b36-46de-b5eb-9fbe6af2356c",
			productId:      "",
			quantity:       1,
			expectedOutput: PlaceOrderOutput{},
			expectedErr:    errors.New(InvalidProductIdError),
		},
		"invalid_order_quantity": {
			userId:         "91d27161-8b36-46de-b5eb-9fbe6af2356c",
			productId:      "073fa958-fe19-41bc-ae99-0974c9cfe9f7",
			quantity:       -1,
			expectedOutput: PlaceOrderOutput{},
			expectedErr:    errors.New(InvalidQuantityError),
		},
	}

	var orderRepository repositories.OrderRepository = &repositories.FakeOrderRepository{}
	var orderService OrderService = &OrderServiceImpl{orderRepository: orderRepository}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			input := PlaceOrderInput{UserId: test.userId, ProductId: test.productId, Quantity: test.quantity}
			output, err := orderService.PlaceOrder(input)
			if (err == nil && test.expectedErr != nil) || (err != nil && test.expectedErr == nil) || (err != nil && err.Error() != test.expectedErr.Error()) {
				t.Fatalf("errors are not equal. expected: %v - got: %v", test.expectedErr, err)
			}
			if err == nil {
				if !isUUIDv4(output.OrderId) {
					t.Fatalf("invalid UUID format for order id: %s", output.OrderId)
				}

				if output.Value != test.expectedOutput.Value {
					t.Fatalf("order values are not equal. expected: %f - got: %f", test.expectedOutput.Value, output.Value)
				}
			}
		})
	}
}

func TestGetOrder(t *testing.T) {
	tests := map[string]struct {
		orderId        string
		userId         string
		productId      string
		quantity       int
		value          float64
		expectedOutput GetOrderOutput
		expectedErr    error
	}{
		"valid_order": {
			orderId:        "91d27161-8b36-46de-b5eb-9fbe6af2356c",
			userId:         "073fa958-fe19-41bc-ae99-0974c9cfe9f7",
			productId:      "91d27161-8b36-46de-b5eb-9fbe6af2356c",
			quantity:       1,
			value:          5,
			expectedOutput: GetOrderOutput{OrderId: "91d27161-8b36-46de-b5eb-9fbe6af2356c", UserId: "073fa958-fe19-41bc-ae99-0974c9cfe9f7", ProductId: "91d27161-8b36-46de-b5eb-9fbe6af2356c", Quantity: 1, Value: 5},
			expectedErr:    nil,
		},
	}

	var orderRepository repositories.OrderRepository = &repositories.FakeOrderRepository{}
	var orderService OrderService = &OrderServiceImpl{orderRepository: orderRepository}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			order := repositories.Order{OrderId: test.orderId, UserId: test.userId, ProductId: test.productId, Quantity: test.quantity, Value: test.value}
			orderRepository.CreateOrder(order)

			input := GetOrderInput{OrderId: test.orderId}
			output, err := orderService.GetOrder(input)

			if (err == nil && test.expectedErr != nil) || (err != nil && test.expectedErr == nil) || (err != nil && err.Error() != test.expectedErr.Error()) {
				t.Fatalf("errors are not equal. expected: %v - got: %v", test.expectedErr, err)
			}
			if err == nil {
				if output.OrderId != test.expectedOutput.OrderId || output.UserId != test.expectedOutput.UserId || output.ProductId != test.expectedOutput.ProductId || output.Quantity != test.expectedOutput.Quantity || output.Value != test.expectedOutput.Value {
					t.Fatalf("orders are not equal. expected: %v - got: %v", test.expectedOutput, output)
				}
			}
		})
	}
}

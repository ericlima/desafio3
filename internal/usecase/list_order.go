package usecase

import "desafio3/internal/entity"

type OrderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type OrdersOutputDTO struct {
	Orders []OrderDTO
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute() (OrdersOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		return OrdersOutputDTO{}, err
	}

	var ordersOutput []OrderDTO
	for _, order := range orders {
		o := OrderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		ordersOutput = append(ordersOutput, o)
	}

	return OrdersOutputDTO{Orders: ordersOutput}, nil
}

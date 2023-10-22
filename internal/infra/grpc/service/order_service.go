package service

import (
	"context"
	"desafio3/internal/usecase"

	pb "desafio3/internal/infra/grpc/pb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUC usecase.CreateOrderUseCase, listOrderUC usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUC,
		ListOrderUseCase:   listOrderUC,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, blank *pb.Blank) (*pb.OrderResponseList, error) {
	orders, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.CreateOrderResponse
	for _, order := range orders.Orders {
		ordersResponse = append(ordersResponse,
			&pb.CreateOrderResponse{
				Id:         order.ID,
				Price:      float32(order.Price),
				Tax:        float32(order.Tax),
				FinalPrice: float32(order.FinalPrice),
			})
	}

	return &pb.OrderResponseList{Orders: ordersResponse}, nil
}

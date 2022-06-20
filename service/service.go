package service

import (
	"assignment-dua/entity"
	"assignment-dua/repository"
	"context"
	"fmt"
)

type ServicesInterface interface {
	Create(ctx context.Context, orders entity.Orders) (*entity.Orders, error)
	Update(ctx context.Context, order entity.Orders, id string) (*entity.Orders, error)
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*entity.Orders, error)
}

type Service struct {
	itemRepo  repository.ItemsRepositoryInterface
	orderRepo repository.OrdersRepositoryInterface
}

func NewService(itemRepo repository.ItemsRepositoryInterface, orderRepo repository.OrdersRepositoryInterface) ServicesInterface {
	return &Service{
		itemRepo:  itemRepo,
		orderRepo: orderRepo,
	}
}

func (s Service) GetAll(ctx context.Context) ([]*entity.Orders, error) {
	//TODO implement me
	orderData, err := s.orderRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	fmt.Println("cek-data-service", orderData)

	return orderData, nil
}

func (s Service) Create(ctx context.Context, orders entity.Orders) (*entity.Orders, error) {
	//TODO implement me
	//var itemTemp []*entity.Items
	order, err := s.orderRepo.CreateOrders(ctx, orders)

	if err != nil {
		return nil, err
	}

	for _, val := range orders.Items {
		_, errItem := s.itemRepo.CreateItems(ctx, val, order.ID)
		if errItem != nil {
			return nil, errItem
		}
	}

	//fmt.Println("order-service", orders)

	if err != nil {
		return nil, err
	}

	return order, nil

}

func (s Service) Update(ctx context.Context, orders entity.Orders, id string) (*entity.Orders, error) {
	//TODO implement me
	//order, err := s.orderRepo.UpdateOrders(ctx, order, id)

	order, err := s.orderRepo.UpdateOrders(ctx, orders, id)

	fmt.Println("ORDERS", orders)
	fmt.Println("ORDER", order)

	if err != nil {
		fmt.Println("errUpdate", err)
		return nil, err
	}

	for _, val := range order.Items {
		errItem := s.itemRepo.UpdateItems(ctx, val, id)
		if errItem != nil {
			return nil, errItem
		}
	}

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return order, nil
}

func (s Service) Delete(ctx context.Context, id string) error {
	//TODO implement me
	err := s.itemRepo.DeleteItems(ctx, id)
	err = s.orderRepo.DeleteOrders(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

package repository

import (
	"assignment-dua/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type OrdersRepositoryInterface interface {
	CreateOrders(ctx context.Context, orders entity.Orders) (*entity.Orders, error)
	UpdateOrders(ctx context.Context, orders entity.Orders, id string) (*entity.Orders, error)
	DeleteOrders(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*entity.Orders, error)
}

type OrdersRepository struct {
	pgpool *pgxpool.Pool
}

func NewOrdersRepository(pgpool *pgxpool.Pool) OrdersRepositoryInterface {
	return &OrdersRepository{pgpool: pgpool}
}

func (o OrdersRepository) GetAll(ctx context.Context) ([]*entity.Orders, error) {
	queryString := `select
		o.id as order_id
		,o.customer_name
		,o.ordered_at
		,json_agg(json_build_object(
			'item_id',i.id
			,'item_code',i.item_code
			,'description',i.description
			,'quantity',i.quantity
			,'order_id',i.order_id
		)) as items
	from orders o join items i
	on o.id = i.order_id
	group by o.id`

	rows, err := o.pgpool.Query(ctx, queryString)
	if err != nil {
		fmt.Println("query row error", err)
	}
	defer rows.Close()

	var orders []*entity.Orders
	for rows.Next() {
		var o entity.Orders
		var itemsStr string

		if serr := rows.Scan(&o.ID, &o.CustomerName, &o.OrderedAt, &itemsStr); serr != nil {
			fmt.Println("Scan error", serr)
		}
		var items []entity.Items
		if err := json.Unmarshal([]byte(itemsStr), &items); err != nil {
			fmt.Errorf("error when parsing items")
		} else {
			o.Items = append(o.Items, items...)
		}
		orders = append(orders, &o)
	}

	return orders, nil
}

func (o OrdersRepository) CreateOrders(ctx context.Context, orders entity.Orders) (*entity.Orders, error) {
	//TODO implement me
	SQL := "insert into orders(customer_name, ordered_at) values($1, $2) returning id"
	rows, err := o.pgpool.Query(ctx, SQL, orders.CustomerName, orders.OrderedAt)

	for rows.Next() {
		err := rows.Scan(&orders.ID)

		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	fmt.Println("items-save", orders)

	return &orders, nil
}

func (o OrdersRepository) UpdateOrders(ctx context.Context, orders entity.Orders, id string) (*entity.Orders, error) {
	//TODO implement me
	SQL := "update orders set customer_name=$1, ordered_at=$2 where id=$3"
	_, err := o.pgpool.Query(ctx, SQL, orders.CustomerName, orders.OrderedAt, id)

	if err != nil {
		fmt.Println("err-orders", err)
		return nil, err
	}

	return &orders, nil
}

func (o OrdersRepository) DeleteOrders(ctx context.Context, id string) error {
	//TODO implement me
	SQL := "delete from items where id=$1"
	_, err := o.pgpool.Exec(ctx, SQL, id)

	if err != nil {
		return err
	}

	return nil
}

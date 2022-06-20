package repository

import (
	"assignment-dua/entity"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

type ItemsRepositoryInterface interface {
	CreateItems(ctx context.Context, items entity.Items, id int) (*entity.Items, error)
	UpdateItems(ctx context.Context, items entity.Items, id string) error
	DeleteItems(ctx context.Context, id string) error
}

type ItemsRepository struct {
	pgpool *pgxpool.Pool
}

func NewItemRepository(pgpool *pgxpool.Pool) ItemsRepositoryInterface {
	return &ItemsRepository{pgpool: pgpool}
}

func (i ItemsRepository) CreateItems(ctx context.Context, items entity.Items, id int) (*entity.Items, error) {
	//TODO implement me
	fmt.Println("ID", id)
	SQL := "insert into items(item_code, description, quantity, order_id) values($1, $2, $3, $4)"
	_, err := i.pgpool.Exec(ctx, SQL, items.ItemCode, items.Description, items.Quantity, id)

	if err != nil {
		return nil, err
	}

	fmt.Println("items-save", items)

	return &items, nil
}

func (i ItemsRepository) UpdateItems(ctx context.Context, items entity.Items, id string) error {
	//TODO implement me
	fmt.Println("ITEMS", items)
	SQL := "update items set item_code=$1, description=$2, quantity=$3 where id=$4"
	_, err := i.pgpool.Exec(ctx, SQL, items.ItemCode, items.Description, items.Quantity, items.ID)

	if err != nil {
		fmt.Println("err-items", err)
		return err
	}

	return nil
}

func (i ItemsRepository) DeleteItems(ctx context.Context, id string) error {
	//TODO implement me
	SQL := "delete from items where order_id=$1"
	_, err := i.pgpool.Exec(ctx, SQL, id)

	if err != nil {
		return err
	}

	return nil
}

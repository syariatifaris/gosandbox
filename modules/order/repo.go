package order

import (
	"fmt"
	"log"

	rdb "github.com/syariatifaris/gosandbox/core/db"
)

type OrderRepo interface {
	GetOrderById(id int64) Order
	UpdateOrder(order Order) error
}

type orderRepo struct {
	db rdb.DB
}

func NewOrderRepo(rdb rdb.DB) OrderRepo {
	return &orderRepo{
		db: rdb,
	}
}

func (o *orderRepo) GetOrderById(id int64) Order {
	mq := `SELECT id, status, driver_id FROM orders WHERE id = ?`
	var orders []Order

	err := o.db.Select(&orders, mq, fmt.Sprintf("%d", id))
	if err != nil {
		log.Println(fmt.Printf("unable to select data with err: %s, order_id %d", err.Error(), id))
	}

	if len(orders) > 0 {
		return orders[0]
	}

	return Order{}
}

func (o *orderRepo) UpdateOrder(order Order) error {
	mq := `UPDATE orders SET status = ? WHERE id = ?`
	tx := o.db.BeginTransaction()
	sqlResult := o.db.Execute(tx, mq, order.Status, order.Id)

	if sqlResult != nil {
		if row, err := sqlResult.RowsAffected(); err != nil || row == 0 {
			return err
		}
	}

	err := o.db.CommitTransaction(tx)
	if err != nil {
		return err
	}
	return nil
}

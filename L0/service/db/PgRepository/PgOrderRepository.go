package PgRepository

import (
	"database/sql"
	"service/model/representation"
	"strings"

	_ "github.com/lib/pq"
)

type PgOrderRepository struct {
	db *sql.DB
}

func NewPgOrderRepository() PgOrderRepository {
	return PgOrderRepository{}
}

func (p *PgOrderRepository) Open(dataSourceName string) error {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *PgOrderRepository) Close() {
	p.db.Close()
}

func (p *PgOrderRepository) CreateOrder(order representation.Order) error {
	_, err := p.db.Exec("insert into orders (order_uid, data) values ($1, $2)", order.Uid, order.Body)

	return err
}

func (p *PgOrderRepository) GetOrders() ([]representation.Order, error) {
	rows, err := p.db.Query("select  * from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]representation.Order, 0)
	for rows.Next() {
		order := representation.Order{}
		err := rows.Scan(&order.Uid, &order.Body)
		if err != nil {
			return nil, err
		}

		order.Uid = strings.Replace(order.Uid, "-", "", -1)
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

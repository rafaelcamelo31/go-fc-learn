package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rafaelcamelo31/go-fc-learn/go-intensive/internal/order/entity"
)

// CREATE TABLE `orders` (
//   `id` varchar(255) NOT NULL,
//   `price` float NOT NULL,
//   `tax` float NOT NULL,
//   `final_price` float NOT NULL,
//   PRIMARY KEY (`id`))

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (orderRepository *OrderRepository) Save(order *entity.Order) error {
	statement, err := orderRepository.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

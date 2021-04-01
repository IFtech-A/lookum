package sqlstore

import (
	"fmt"
	"strings"

	"database/sql"

	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/sirupsen/logrus"
)

type OrderRepo struct {
	store *Store
}

func (r *OrderRepo) Create(order *model.Order) (int, error) {

	createOrderSql := `INSERT INTO orders(user_id, status) VALUES ($1, $2) RETURNING id`
	createOrderItemsSql := `INSERT INTO order_items(order_id, product_id, quantity, at_price) VALUES ($1, $2, $3, $4)`
	// productCheckSql := `SELECT quantity FROM products WHERE $1`

	transaction, err := r.store.db.Begin()
	if err != nil {
		return 0, err
	}

	// res := transaction.QueryRow(productCheckSql, order.OrderItems[0].ProductID)
	// if res.Err() != nil {
	// 	transaction.Rollback()
	// 	return 0, res.Err()
	// }

	// var realQuantity float32
	// err = res.Scan(&realQuantity)
	// if err != nil {
	// 	transaction.Rollback()
	// 	return 0, err
	// }

	// if realQuantity < order.OrderItems[0].Quantity {
	// 	return 0, errors.New("insufficient quantity")
	// }

	res := transaction.QueryRow(createOrderSql, 0, order.Status)
	if res.Err() != nil {
		transaction.Rollback()
		return 0, res.Err()
	}
	err = res.Scan(&order.ID)
	if err != nil {
		logrus.Error(err.Error())

		row := transaction.QueryRow("SELECT max(id) FROM orders")
		if row.Err() != nil {
			transaction.Rollback()
			return 0, row.Err()
		}
		err = row.Scan(&order.ID)
		if err != nil {
			transaction.Rollback()
			return 0, err
		}
	}

	for _, v := range order.OrderItems {
		_, err = transaction.Exec(createOrderItemsSql, order.ID, v.ProductID, v.Quantity, v.AtPrice)
		if err != nil {
			transaction.Rollback()
			return 0, err
		}
	}

	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return 0, nil
	}

	return order.ID, nil
}

func (r *OrderRepo) GetOrders(limit int) ([]*model.Order, error) {

	var sql strings.Builder

	sql.WriteString("SELECT id, user_id, status, created_at FROM orders")

	if limit == 0 {
		limit = 20
	}
	sql.WriteString(fmt.Sprintf(" LIMIT %v ", limit))

	rows, err := r.store.db.Query(sql.String())
	if err != nil {
		return nil, err
	}

	var order *model.Order
	var orders []*model.Order
	for rows.Next() {

		order = &model.Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		if orders == nil {
			orders = make([]*model.Order, 0)
		}
		orders = append(orders, order)
	}

	return orders, nil
}
func (r *OrderRepo) GetOrder(orderID int) (*model.Order, error) {

	querySql := `SELECT id, user_id, status, created_at FROM orders WHERE id=$1`

	row := r.store.db.QueryRow(querySql, orderID)

	order := &model.Order{}
	err := row.Scan(&order.ID, &order.UserID, &order.Status, &order.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return order, nil
}

func (r *OrderRepo) GetOrderWithItems(orderID int) (*model.Order, error) {

	getOrderItemsSql := `SELECT product_id, quantity, at_price FROM order_items WHERE id=$1`

	order, err := r.GetOrder(orderID)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	rows, err := r.store.db.Query(getOrderItemsSql, orderID)
	if err != nil {
		return order, err
	}

	var orderItem *model.OrderItem
	for rows.Next() {

		orderItem = &model.OrderItem{OrderID: orderID}
		err = rows.Scan(&orderItem.ProductID, &orderItem.Quantity, &orderItem.AtPrice)
		if err != nil {
			break
		}

		if order.OrderItems == nil {
			order.OrderItems = make([]*model.OrderItem, 0)
		}

		order.OrderItems = append(order.OrderItems, orderItem)
	}

	return order, err
}

func (r *OrderRepo) DeleteOrder(orderID int) error {

	deleteOrderSql := `DELETE FROM orders WHERE id=$1`
	deleteOrderItemsSql := `DELETE FROM order_items WHERE order_id=$1`

	tr, err := r.store.db.Begin()
	if err != nil {
		return err
	}

	_, err = tr.Exec(deleteOrderSql, orderID)
	if err != nil {
		tr.Rollback()
		return err
	}

	_, err = tr.Exec(deleteOrderItemsSql, orderID)
	if err != nil {
		tr.Rollback()
		return err
	}

	err = tr.Commit()
	if err != nil {
		tr.Rollback()
		return err
	}

	return nil
}

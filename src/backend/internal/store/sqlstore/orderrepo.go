package sqlstore

import (
	"strings"

	"database/sql"

	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/sirupsen/logrus"
)

type OrderRepo struct {
	store *Store
}

func (r *OrderRepo) Create(o *model.Order) (int, error) {

	createOrderSql := `INSERT INTO order(
		user_id,
		address_id,
		token,
		status,
		sub_total,
		item_discount,
		tax,
		shipping,
		total,
		promo,
		total_discount,
		grand_total)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING id`
	createOrderItemsSql := `INSERT INTO order_item(
		product_id,
		order_id,
		sku,
		price,
		discount,
		quantity)
	VALUES ($1, $2, $3, $4, $5, $6)`
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

	err = transaction.QueryRow(createOrderSql,
		o.UserID,
		o.AddressID,
		o.Token,
		o.Status,
		o.SubTotal,
		o.ItemDiscount,
		o.Tax,
		o.Shipping,
		o.Total,
		o.Promo,
		o.TotalDiscount,
		o.GrandTotal,
	).Scan(&o.ID)
	if err != nil {
		logrus.Error(err.Error())
		err = transaction.QueryRow("SELECT max(id) FROM order").Scan(&o.ID)
		if err != nil {
			transaction.Rollback()
			return 0, err
		}
	}

	for _, v := range o.OrderItems {
		_, err = transaction.Exec(createOrderItemsSql,
			v.ProductID,
			o.ID,
			v.SKU,
			v.Price,
			v.Discount,
			v.Quantity)
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

	return o.ID, nil
}

func (r *OrderRepo) GetOrders(userID, limit int) ([]*model.Order, error) {

	var querySql strings.Builder

	querySql.WriteString(`SELECT 
		id, 
		user_id,
		address_id,
		token,
		status,
		sub_total,
		item_discount,
		tax,
		shipping,
		total,
		promo,
		total_discount,
		grand_total,
		created_at,
		updated_at
	FROM order WHERE user_id=$1 LIMIT $2`)

	if limit == 0 {
		limit = 20
	}

	rows, err := r.store.db.Query(querySql.String(), userID, limit)
	if err != nil {
		return nil, err
	}

	var o *model.Order
	var orders []*model.Order
	var updatedAtNullable sql.NullTime
	for rows.Next() {
		o = &model.Order{}
		err := rows.Scan(
			&o.ID,
			&o.UserID,
			&o.AddressID,
			&o.Token,
			&o.Status,
			&o.SubTotal,
			&o.ItemDiscount,
			&o.Tax,
			&o.Shipping,
			&o.Total,
			&o.Promo,
			&o.TotalDiscount,
			&o.GrandTotal,
			&o.CreatedAt,
			&updatedAtNullable,
		)
		if err != nil {
			return nil, err
		}
		if updatedAtNullable.Valid {
			o.UpdatedAt = updatedAtNullable.Time
		}
		orders = append(orders, o)
	}

	return orders, nil
}

func (r *OrderRepo) GetOrder(orderID int) (*model.Order, error) {

	querySql := `SELECT 
		id, 
		user_id,
		address_id,
		token,
		status,
		sub_total,
		item_discount,
		tax,
		shipping,
		total,
		promo,
		total_discount,
		grand_total,
		created_at,
		updated_at
	FROM order 
	WHERE id=$1`

	row := r.store.db.QueryRow(querySql, orderID)

	o := &model.Order{}
	var updatedAtNullable sql.NullTime
	err := row.Scan(&o.ID,
		&o.UserID,
		&o.AddressID,
		&o.Token,
		&o.Status,
		&o.SubTotal,
		&o.ItemDiscount,
		&o.Tax,
		&o.Shipping,
		&o.Total,
		&o.Promo,
		&o.TotalDiscount,
		&o.GrandTotal,
		&o.CreatedAt,
		&updatedAtNullable,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if updatedAtNullable.Valid {
		o.UpdatedAt = updatedAtNullable.Time
	}

	return o, nil
}

func (r *OrderRepo) GetOrderWithItems(orderID int) (*model.Order, error) {

	getOrderItemsSql := `SELECT
		product_id,
		order_id,
		sku,
		price,
		discount,
		quantity,
		created_at,
		updated_at
	FROM order_item WHERE id=$1`

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

	var oi *model.OrderItem
	var updatedAtNullable sql.NullTime
	for rows.Next() {

		oi = &model.OrderItem{}
		err = rows.Scan(
			&oi.ProductID,
			&oi.OrderID,
			&oi.SKU,
			&oi.Price,
			&oi.Discount,
			&oi.Quantity,
			&oi.CreatedAt,
			&updatedAtNullable,
		)
		if err != nil {
			break
		}
		if updatedAtNullable.Valid {
			oi.UpdatedAt = updatedAtNullable.Time
		}

		order.OrderItems = append(order.OrderItems, oi)
	}

	return order, err
}

func (r *OrderRepo) DeleteOrder(orderID int) error {

	deleteOrderSql := `DELETE FROM order WHERE id=$1`
	deleteOrderItemsSql := `DELETE FROM order_item WHERE order_id=$1`

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

package sqlstore

import (
	"strings"

	"database/sql"

	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/sirupsen/logrus"
)

type CartRepo struct {
	store *Store
}

func (r *CartRepo) Create(c *model.Cart) (int, error) {

	createCartSql := `INSERT INTO cart(
		user_id,
		address_id,
		token,
		status,
		content)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	createCartItemsSql := `INSERT INTO cart_item(
		product_id,
		cart_id,
		sku,
		price,
		discount,
		quantity,
		content)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err := r.store.db.QueryRow(createCartSql,
		c.UserID,
		c.AddressID,
		c.Token,
		c.Status,
		c.Content,
	).Scan(&c.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = r.store.db.QueryRow("SELECT max(id) FROM cart").Scan(&c.ID)
		} else {
			logrus.Error(err.Error())
		}
		if err != nil {
			return 0, err
		}
	}

	for _, v := range c.CartItems {
		_, err = r.store.db.Exec(createCartItemsSql,
			v.ProductID,
			c.ID,
			v.SKU,
			v.Price,
			v.Discount,
			v.Quantity,
			v.Content,
		)
		if err != nil {
			return 0, err
		}
	}

	return c.ID, nil
}

func (r *CartRepo) GetCarts(userID, limit int) ([]*model.Cart, error) {

	var querySql strings.Builder

	querySql.WriteString(`SELECT 
		id, 
		user_id,
		address_id,
		token,
		status,
		created_at,
		updated_at,
		content
	FROM cart WHERE user_id=$1 LIMIT $2`)

	if limit == 0 {
		limit = 20
	}

	rows, err := r.store.db.Query(querySql.String(), userID, limit)
	if err != nil {
		return nil, err
	}

	var c *model.Cart
	var carts []*model.Cart
	var updatedAtNullable sql.NullTime
	for rows.Next() {
		c = &model.Cart{}
		err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.AddressID,
			&c.Token,
			&c.Status,
			&c.CreatedAt,
			&updatedAtNullable,
			&c.Content,
		)
		if err != nil {
			return nil, err
		}
		if updatedAtNullable.Valid {
			c.UpdatedAt = updatedAtNullable.Time
		}
		carts = append(carts, c)
	}

	return carts, nil
}

func (r *CartRepo) GetCart(cartID int) (*model.Cart, error) {

	querySql := `SELECT 
		id, 
		user_id,
		address_id,
		token,
		status,
		created_at,
		updated_at,
		content
	FROM cart 
	WHERE id=$1`

	row := r.store.db.QueryRow(querySql, cartID)

	c := &model.Cart{}
	var updatedAtNullable sql.NullTime
	err := row.Scan(&c.ID,
		&c.UserID,
		&c.AddressID,
		&c.Token,
		&c.Status,
		&c.CreatedAt,
		&updatedAtNullable,
		&c.Content,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if updatedAtNullable.Valid {
		c.UpdatedAt = updatedAtNullable.Time
	}

	return c, nil
}

func (r *CartRepo) GetCartWithItems(cartID int) (*model.Cart, error) {

	getCartItemsSql := `SELECT
		product_id,
		cart_id,
		sku,
		price,
		discount,
		quantity,
		created_at,
		updated_at,
		content
	FROM cart_item WHERE id=$1`

	cart, err := r.GetCart(cartID)
	if err != nil {
		return nil, err
	}

	if cart == nil {
		return nil, nil
	}

	rows, err := r.store.db.Query(getCartItemsSql, cartID)
	if err != nil {
		return cart, err
	}

	var ci *model.CartItem
	var updatedAtNullable sql.NullTime
	for rows.Next() {

		ci = &model.CartItem{}
		err = rows.Scan(
			&ci.ProductID,
			&ci.CartID,
			&ci.SKU,
			&ci.Price,
			&ci.Discount,
			&ci.Quantity,
			&ci.CreatedAt,
			&updatedAtNullable,
			&ci.Content,
		)
		if err != nil {
			break
		}
		if updatedAtNullable.Valid {
			ci.UpdatedAt = updatedAtNullable.Time
		}

		cart.CartItems = append(cart.CartItems, ci)
	}

	return cart, err
}

func (r *CartRepo) DeleteCart(cartID int) error {

	deleteCartSql := `DELETE FROM cart WHERE id=$1`
	deleteCartItemsSql := `DELETE FROM cart_item WHERE cart_id=$1`

	_, err := r.store.db.Exec(deleteCartSql, cartID)
	if err != nil {
		return err
	}

	_, err = r.store.db.Exec(deleteCartItemsSql, cartID)
	if err != nil {
		return err
	}

	return nil
}

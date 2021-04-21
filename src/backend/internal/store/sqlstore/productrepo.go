package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/iftech-a/lookum/src/backend/internal/model"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

//ProductRepo handler for database operations on product
type ProductRepo struct {
	store *Store
}

//Create inserts new product to database, returns error on fail
func (r *ProductRepo) Create(p *model.Product) error {

	sql := `INSERT INTO product(
			user_id,
			title,
			meta_title,
			slug,
			summary,
			type,
			sku,
			price,
			discount,
			quantity,
			available,
			content)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
	RETURNING id`

	err := r.store.db.QueryRow(sql,
		p.UserID,
		p.Title,
		p.MetaTitle,
		p.Slug,
		p.Summary,
		p.Type,
		p.SKU,
		p.Price,
		p.Discount,
		p.Quantity,
		p.Available,
		p.Content).Scan(&p.ID)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

//UpdateProduct updates attributes in the program
func (r *ProductRepo) UpdateProduct(p *model.Product) error {

	sql := `UPDATE product
	SET
		user_id=$1,
		title=$2,
		meta_title=$3,
		slug=$4,
		summary=$5,
		type=$6,
		sku=$7,
		price=$8,
		discount=$9,
		quantity=$10,
		available=$11,
		content=$12,
		updated_at=$13
	WHERE id=$14`

	_, err := r.store.db.Exec(sql,
		p.UserID,
		p.Title,
		p.MetaTitle,
		p.Slug,
		p.Summary,
		p.Type,
		p.SKU,
		p.Price,
		p.Discount,
		p.Quantity,
		p.Available,
		p.Content,
		time.Now(),
		p.ID)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

//DeleteProduct deletes product using product ID
func (r *ProductRepo) DeleteProduct(id int) error {

	sql := `DELETE FROM product WHERE id=$1`

	_, err := r.store.db.Exec(sql, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

//GetProducts retrieve all products from database
//@limit int - limit the resulting rows to given limit parameter, if 0 the default 20 will be used
//@category int - retrieve products from only one category, if 0 all categories will be used
//@@return - returns array of products
func (r *ProductRepo) GetProducts(limit int, category int) ([]*model.Product, error) {

	var querySql strings.Builder

	querySql.WriteString(`SELECT 
		product.id, 
		product.user_id,
		product.title,
		product.meta_title,
		product.slug,
		product.summary,
		product.type,
		product.sku,
		product.price,
		product.discount,
		product.quantity,
		product.available,
		product.content,
		product.created_at,
		product.updated_at
	FROM product `)

	if category != 0 {
		querySql.WriteString(fmt.Sprintf(`
		INNER JOIN product_category 
			ON product.id = product_category.product_id 
		WHERE product_category.category_id = %v `, category))
	}

	if limit == 0 {
		limit = 20
	}
	querySql.WriteString(fmt.Sprintf(" LIMIT %v ", limit))

	rows, err := r.store.db.Query(querySql.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]*model.Product, 0, limit)
	var updatedAtNullable sql.NullTime
	for rows.Next() {
		p := model.NewProduct()
		if err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Title,
			&p.MetaTitle,
			&p.Slug,
			&p.Summary,
			&p.Type,
			&p.SKU,
			&p.Price,
			&p.Discount,
			&p.Quantity,
			&p.Available,
			&p.Content,
			&p.CreatedAt,
			&updatedAtNullable); err != nil {
			return nil, err
		}

		if updatedAtNullable.Valid {
			p.UpdatedAt = updatedAtNullable.Time
		}

		products = append(products, p)
	}

	for _, p := range products {
		p.Images, err = r.GetImages(p.ID)
		if err != nil {
			logrus.Errorf(err.Error())
		}
	}

	return products, nil
}

//GetProduct retrieve product using the product ID
//@ID int - Identification number of product on database
//@@*model.Product - Product struct with for a given. It is nil if ID was not found on Database
//@@error - error sturcture to show error on database request. It is nil if no error has occured
func (r *ProductRepo) GetProduct(id int) (*model.Product, error) {

	querySql := `SELECT 
		id, 
		user_id,
		title,
		meta_title,
		slug,
		summary,
		type,
		sku,
		price,
		discount,
		quantity,
		available,
		content,
		created_at,
		updated_at
	FROM product WHERE id=$1`

	p := &model.Product{}
	var updatedAtNullable sql.NullTime
	err := r.store.db.QueryRow(querySql, id).
		Scan(
			&p.ID,
			&p.UserID,
			&p.Title,
			&p.MetaTitle,
			&p.Slug,
			&p.Summary,
			&p.Type,
			&p.SKU,
			&p.Price,
			&p.Discount,
			&p.Quantity,
			&p.Available,
			&p.Content,
			&p.CreatedAt,
			&updatedAtNullable,
		)
	if err != nil {
		return nil, err
	}

	if updatedAtNullable.Valid {
		p.UpdatedAt = updatedAtNullable.Time
	}

	p.Images, err = r.GetImages(id)
	if err != nil {
		logrus.Errorln(err.Error())
	}

	return p, nil
}

//GetImages ...
func (r *ProductRepo) GetImages(productID int) ([]*model.Image, error) {
	sql := `SELECT 
		id,
		file_uri,
		filename,
		main
	FROM image WHERE product_id=$1`

	rows, err := r.store.db.Query(sql, productID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var image *model.Image = nil
	var images []*model.Image = nil

	for rows.Next() {
		image = &model.Image{}
		err := rows.Scan(&image.ID, &image.FileURI, &image.Filename, &image.Main)
		if err != nil {
			break
		}

		images = append(images, image)
	}

	return images, err
}

//AddImage ...
func (r *ProductRepo) AddImage(productID int, filename, fileuri string) error {
	sql := `
	INSERT INTO image(
		product_id,
		filename,
		file_uri) 
	VALUES ($1, $2, $3)`

	_, err := r.store.db.Exec(sql, productID, filename, fileuri)
	if err != nil {
		return err
	}

	return nil
}

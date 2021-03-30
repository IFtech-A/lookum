package sqlstore

import (
	"fmt"
	"strings"

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

	sql := `INSERT INTO products(name, "desc", price, discount, category_id)
	VALUES ($1,$2,$3,$4,$5)`

	_, err := r.store.db.Exec(sql, p.Name, p.Description, p.Price, p.Discount, p.CategoryID)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

//UpdateProduct updates attributes in the program
func (r *ProductRepo) UpdateProduct(p *model.Product) error {

	sql := `UPDATE products
	SET
		name=$1,
		"desc"=$2,
		price=$3,
		discount=$4,
		category_id=$5
	WHERE id=$6`

	_, err := r.store.db.Exec(sql, p.Name, p.Description, p.Price, p.Discount, p.CategoryID, p.ID)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

//DeleteProduct deletes product using product ID
func (r *ProductRepo) DeleteProduct(id int) error {

	sql := `DELETE FROM products WHERE id=$1`

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

	var sql strings.Builder

	sql.WriteString("SELECT id, name, \"desc\", price, discount, status, likes, category_id FROM products ")

	if category != 0 {
		sql.WriteString(fmt.Sprintf("WHERE category_id = %v ", category))
	}

	if limit == 0 {
		limit = 20
	}
	sql.WriteString(fmt.Sprintf(" LIMIT %v ", limit))

	rows, err := r.store.db.Query(sql.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]*model.Product, 0, limit)
	for rows.Next() {
		p := model.NewProduct()
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Discount, &p.Status, &p.Likes, &p.CategoryID); err != nil {
			return nil, err
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

	sql := `SELECT id, name, "desc", price, discount, likes, status, created_at, category_id
			FROM products
			WHERE id=$1`

	product := &model.Product{}
	err := r.store.db.QueryRow(sql, id).
		Scan(&product.ID, &product.Name, &product.Description,
			&product.Price, &product.Discount, &product.Likes,
			&product.Status, &product.CreatedAt, &product.CategoryID)
	if err != nil {
		return nil, err
	}

	product.Images, err = r.GetImages(id)
	if err != nil {
		logrus.Errorln(err.Error())
	}

	return product, nil
}

//GetImages ...
func (r *ProductRepo) GetImages(productID int) ([]*model.Image, error) {
	sql := `SELECT id, filename, file_uri
	FROM images
	WHERE product_id=$1`

	rows, err := r.store.db.Query(sql, productID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var image *model.Image = nil
	var images []*model.Image = nil

	for rows.Next() {
		image = &model.Image{}
		err := rows.Scan(&image.ID, &image.Filename, &image.FileURI)
		if err != nil {
			break
		}

		if images == nil {
			images = make([]*model.Image, 0)
		}
		images = append(images, image)
	}

	return images, err
}

//AddImage ...
func (r *ProductRepo) AddImage(ID int, filename, fileuri string) error {
	sql := `INSERT INTO images(product_id, filename, file_uri) 
	VALUES ($1, $2, $3)
	`

	_, err := r.store.db.Exec(sql, ID, filename, fileuri)
	if err != nil {
		return err
	}

	return nil
}

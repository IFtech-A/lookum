package apiserver

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/iftech-a/lookum/src/backend/internal/model"

	"github.com/labstack/echo/v4"
)

const imagesPrefixPath = "images"

func (s *Server) getProduct(c echo.Context) error {

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	product, err := s.s.Product().GetProduct(productID)
	if err != nil {
		s.e.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	c.JSONPretty(http.StatusOK, product, " ")
	return nil
}

func (s *Server) getProducts(c echo.Context) error {

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 0
	}

	category, err := strconv.Atoi(c.QueryParam("category"))
	if err != nil {
		category = 0
	}

	products, err := s.s.Product().GetProducts(limit, category)
	if err != nil {
		log.Println(err.Error())
		c.NoContent(http.StatusInternalServerError)
		return err
	}

	c.Response().Writer.Header().Add("Access-Control-Allow-Origin", "*")
	err = c.JSONPretty(http.StatusOK, products, " ")

	return err
}

func parseRequestProduct(c echo.Context) (*model.Product, error) {

	badRequestErr := errors.New("bad_request")
	name := c.FormValue("name")
	if name == "" {
		return nil, badRequestErr
	}
	desc := c.FormValue("desc")

	price, err := strconv.ParseFloat(c.FormValue("price"), 32)
	if err != nil {
		price = 0
	}
	discount, err := strconv.ParseFloat(c.FormValue("discount"), 32)
	if err != nil {
		discount = 0
	}
	instock, err := strconv.Atoi(c.FormValue("instock"))
	if err != nil {
		instock = 0
	}

	return &model.Product{
		Name:        name,
		Description: desc,
		Price:       float32(price),
		Discount:    float32(discount),
		Stock:       instock,
		CategoryID:  1,
	}, nil

}
func (s *Server) createProduct(c echo.Context) error {

	badRequestErr := map[string]string{"error": "bad_request"}
	product, err := parseRequestProduct(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, badRequestErr)
	}

	err = s.s.Product().Create(product)
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}

func (s *Server) uploadProductImages(c echo.Context) error {
	form, err := c.MultipartForm()

	if err != nil {
		fmt.Print(err)
		return err
	}
	product := form.Value["id"]
	if len(product) == 0 {
		return fmt.Errorf("no product id")
	}

	productID, err := strconv.Atoi(product[0])
	if err != nil {
		return fmt.Errorf("no product id")
	}

	files := form.File["images"]

	for _, file := range files {
		fmt.Print(file.Filename)
		s.e.Logger.Debugf("%v", file.Filename)

		filePath := fmt.Sprintf("%v/%v_%v", imagesPrefixPath, productID, file.Filename)
		err := s.s.Product().AddImage(productID, file.Filename, filePath)
		if err != nil {
			return err
		}
		// Source
		src, err := file.Open()
		if err != nil {
			s.e.Logger.Error(err)
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(filePath)
		if err != nil {
			s.e.Logger.Error(err)
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}
	return nil
}

package apiserver

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func (s *Server) getOrder(c echo.Context) error {

	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	order, err := s.s.Order().GetOrder(orderID)
	if err != nil {
		s.e.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	if order != nil {
		c.JSONPretty(http.StatusOK, order, " ")
	} else {
		c.JSONBlob(http.StatusNotFound, []byte("{ \"error\": \"not found\"}"))

	}
	return nil
}

func (s *Server) getOrders(c echo.Context) error {

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 0
	}

	orders, err := s.s.Order().GetOrders(limit)
	if err != nil {
		log.Println(err.Error())
		c.NoContent(http.StatusInternalServerError)
		return err
	}

	c.Response().Writer.Header().Add("Access-Control-Allow-Origin", "*")

	if orders != nil {
		err = c.JSONPretty(http.StatusOK, orders, " ")
	} else {
		err = c.JSONBlob(http.StatusOK, []byte("[]"))
	}

	return err
}

func (s *Server) getOrderWithItems(c echo.Context) error {

	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	order, err := s.s.Order().GetOrderWithItems(orderID)
	if err != nil {
		s.e.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	if order != nil {
		c.JSONPretty(http.StatusOK, order, " ")
	} else {
		c.JSONBlob(http.StatusNotFound, []byte("{ \"error\": \"not found\"}"))

	}
	return nil
}

func parseRequestOrder(c echo.Context) (*model.Order, error) {

	badRequestErr := errors.New("bad_request")

	productID, err := strconv.Atoi(c.FormValue("product_id"))
	if err != nil {
		return nil, badRequestErr
	}
	quantity := c.FormValue("quantity")
	if quantity == "" {
		return nil, badRequestErr
	}
	quantityVal, err := strconv.ParseFloat(quantity, 32)
	if err != nil {
		return nil, badRequestErr
	}

	atPrice := c.FormValue("price")
	if atPrice == "" {
		return nil, badRequestErr
	}
	atPriceVal, err := strconv.ParseFloat(atPrice, 32)
	if err != nil {
		return nil, badRequestErr
	}

	order := model.NewOrder()
	order.OrderItems = make([]*model.OrderItem, 0)
	order.OrderItems = append(order.OrderItems, model.NewOrderItem(productID, float32(quantityVal), float32(atPriceVal)))

	return order, nil

}
func (s *Server) createOrder(c echo.Context) error {

	badRequestErr := map[string]string{"error": "bad_request"}
	order, err := parseRequestOrder(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, badRequestErr)
	}

	orderID, err := s.s.Order().Create(order)
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
	}

	s.e.Logger.Debugf("order created %v", orderID)
	return c.NoContent(http.StatusCreated)
}

package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"strconv"

	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func (s *Server) getCart(c echo.Context) error {

	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	cart, err := s.s.Cart().GetCart(cartID)
	if err != nil {
		s.e.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	if cart != nil {
		c.JSONPretty(http.StatusOK, cart, " ")
	} else {
		c.JSONBlob(http.StatusNotFound, []byte("{ \"error\": \"not found\"}"))

	}
	return nil
}

func (s *Server) getCarts(c echo.Context) error {

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 0
	}

	carts, err := s.s.Cart().GetCarts(limit, 1)
	if err != nil {
		log.Println(err.Error())
		c.NoContent(http.StatusInternalServerError)
		return err
	}

	c.Response().Writer.Header().Add("Access-Control-Allow-Origin", "*")

	if carts != nil {
		err = c.JSONPretty(http.StatusOK, carts, " ")
	} else {
		err = c.JSONBlob(http.StatusOK, []byte("[]"))
	}

	return err
}

func (s *Server) getCartWithItems(c echo.Context) error {

	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	cart, err := s.s.Cart().GetCartWithItems(cartID)
	if err != nil {
		s.e.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	if cart != nil {
		c.JSONPretty(http.StatusOK, cart, " ")
	} else {
		c.JSONBlob(http.StatusNotFound, []byte("{ \"error\": \"not found\"}"))

	}
	return nil
}

func (s *Server) createCart(c echo.Context) error {

	badRequestErr := map[string]string{"error": "bad_request"}

	contentType, _, err := mime.ParseMediaType(c.Request().Header.Get("Content-Type"))
	if err != nil {
		s.e.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, badRequestErr)
	}
	if contentType == "" || contentType != "application/json" {
		s.e.Logger.Error("createProduct: content type is not application/json")
		return c.JSON(http.StatusBadRequest, badRequestErr)
	}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, badRequestErr)
	}

	cart := model.NewCart()
	err = json.Unmarshal(body, cart)
	if err != nil {
		s.e.Logger.Errorf("createCart: %v", err.Error())
		return c.JSON(http.StatusBadRequest, badRequestErr)
	}

	cartID, err := s.s.Cart().Create(cart)
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
	}

	s.e.Logger.Debugf("cart created %v", cartID)
	return c.NoContent(http.StatusCreated)
}

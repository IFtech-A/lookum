package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"mime"
	"net/http"
	"strconv"

	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func (s *Server) getUser(c echo.Context) error {

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	user, err := s.s.User().GetUser(userID)
	if err != nil {
		s.e.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, map[string]string{"error": "bad_request"})
		return err
	}

	if user != nil {
		c.JSONPretty(http.StatusOK, user, " ")
	} else {
		c.JSONBlob(http.StatusNotFound, []byte("{ \"error\": \"not found\"}"))

	}
	return nil
}

func (s *Server) createUser(c echo.Context) error {

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

	user := model.NewUser()
	err = json.Unmarshal(body, user)
	if err != nil {
		s.e.Logger.Errorf("createUser: %v", err.Error())
		return c.JSON(http.StatusBadRequest, badRequestErr)
	}

	//TODO implement user existance check

	//TODO implement password validation
	//TODO implement password hashing

	userID, err := s.s.User().Create(user)
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
	}

	s.e.Logger.Debugf("user created %v", userID)
	return c.NoContent(http.StatusCreated)
}

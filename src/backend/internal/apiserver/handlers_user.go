package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"mime"
	"net/http"
	"strconv"
	"time"

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
	userAlreadyExistsErr := map[string]string{"error": "user_already_exists"}

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
	dbUser, err := s.s.User().GetUserByEmail(user.Email)
	if err != nil {
		s.e.Logger.Errorf("GetUserByEmail failed: %v", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	if dbUser != nil {
		return c.JSON(http.StatusConflict, userAlreadyExistsErr)
	}

	//TODO implement password validation
	// min: 8, max: 40, mixed case, special characters included, numbers included, no repitations, no overlap with email
	//TODO implement password hashing
	passwordHash, err := user.GeneratePasswordHash()
	if err != nil {
		s.e.Logger.Errorf("GeneratePasswordHash: %v", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	user.Password = string(passwordHash)

	userID, err := s.s.User().Create(user)
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
	}

	s.e.Logger.Debugf("user created %v", userID)
	return c.NoContent(http.StatusCreated)
}

func (s *Server) Login(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	user := model.NewUser()
	err = json.Unmarshal(body, user)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	dbUser, err := s.s.User().GetUserByEmail(user.Email)
	if err != nil {
		s.e.Logger.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := dbUser.CheckPassword(user.Password); err != nil {
		s.e.Logger.Warnf("Login failed %v", user.Email)
		return c.NoContent(http.StatusBadRequest)
	}

	token, err := dbUser.GenerateToken([]byte(s.c.SessionKey))
	if err != nil {
		s.e.Logger.Errorf("jwt claim signing failed: %v", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

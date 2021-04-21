package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/labstack/echo/v4"
)

func (s *Server) Register(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	user := model.NewUser()
	err = json.Unmarshal(body, user)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	hashedPassword, err := user.GeneratePasswordHash()
	user.Password = string(hashedPassword)
	userStore[user.Email] = user

	return c.NoContent(http.StatusOK)
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

	info, ok := userStore[user.Email]
	if !ok {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := info.CheckPassword(user.Password); err != nil {
		s.l.Warnf("Login failed %v", user.Email)
		return c.NoContent(http.StatusBadRequest)
	}
	user.Password = ""
	info.Password = ""

	token, err := info.GenerateToken([]byte(s.c.SessionKey))
	if err != nil {
		s.l.Errorf("jwt claim signing failed: %v", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, user)
}

func (s *Server) Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "jwt",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}

	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

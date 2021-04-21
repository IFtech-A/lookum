package admin

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/iftech-a/lookum/src/backend/internal/config"
	"github.com/iftech-a/lookum/src/backend/internal/model"
	"github.com/iftech-a/lookum/src/backend/internal/store/sqlstore"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	e *echo.Echo
	l *logrus.Logger
	c *config.Config
	s *sqlstore.Store
}

var userStore map[string]*model.User

func NewAdminServer(conf *config.Config) *Server {

	s := &Server{
		e: echo.New(),
		l: logrus.New(),
		c: conf,
	}

	return s
}

func (s *Server) Start() error {
	db, err := s.newDB()
	if err != nil {
		return err
	}
	s.s = sqlstore.New(db)

	s.configureRoutes()
	s.l.SetLevel(logrus.TraceLevel)
	/* Test code */
	userStore = make(map[string]*model.User)
	/* Test code  end */

	return s.e.Start(fmt.Sprintf("%v:%v", s.c.BindAddr, s.c.PortHTTP))
}

func (s *Server) newDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", s.c.GetDatabaseURL())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (s *Server) configureRoutes() {

	s.e.POST("/register", s.Register)
	s.e.POST("/login", s.Login)
	s.e.GET("/", func(c echo.Context) error {

		cookie, err := c.Cookie("jwt")
		if err != nil {
			s.l.Error(err.Error())
			return c.NoContent(http.StatusUnauthorized)
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			s.l.Error(err.Error())
			return c.NoContent(http.StatusUnauthorized)
		}

		claim := token.Claims.(*jwt.StandardClaims)

		user, ok := userStore[claim.Subject]
		if !ok {
			return c.NoContent(http.StatusUnauthorized)
		}

		return c.JSON(http.StatusOK, user)

	})
	s.e.GET("/logout", s.Logout)

}

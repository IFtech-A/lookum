package apiserver

import (
	"database/sql"
	"lookum/internal/store/sqlstore"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
	s *sqlstore.Store
}

//Start starts API server using the given configuration parameters
func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	store := sqlstore.New(db)
	srv := newServer(store)
	return srv.e.Start(config.BindAddr)

}

func newServer(store *sqlstore.Store) *Server {
	server := &Server{
		s: store,
	}
	e := echo.New()
	e.GET("/product/:id", server.getProduct)
	e.GET("/products", server.getProducts)
	e.POST("/product", server.createProduct)
	e.POST("/fileUpload", server.uploadProductImages)
	e.Static("/", "web")
	e.Static("/images/", "images")
	server.e = e
	return server

}

func newDB(databaseURL string) (*sql.DB, error) {

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

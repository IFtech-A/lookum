package apiserver

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/iftech-a/lookum/src/backend/internal/store/sqlstore"

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

	if config.CertificatePath != "" && config.PrivateKeyPath != "" {
		cert, err := tls.LoadX509KeyPair(config.CertificatePath, config.PrivateKeyPath)
		if err != nil {
			return err
		}
		httpsServer := &http.Server{
			Addr: fmt.Sprintf("%v:%v", config.BindAddr, config.PortHTTPS),
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{
					cert,
				},
			},
			ReadHeaderTimeout: time.Second * 10,
			IdleTimeout:       time.Second * 60,
		}
		go srv.e.StartServer(httpsServer)
	}

	httpServer := &http.Server{
		Addr:              fmt.Sprintf("%v:%v", config.BindAddr, config.PortHTTP),
		ReadHeaderTimeout: time.Second * 10,
		IdleTimeout:       time.Second * 60,
	}

	return srv.e.StartServer(httpServer)
}

func newServer(store *sqlstore.Store) *Server {
	server := &Server{
		s: store,
	}
	e := echo.New()
	e.GET("/product/:id", server.getProduct)
	e.GET("/products", server.getProducts)
	e.POST("/product", server.createProduct)
	e.GET("/order/:id", server.getOrder)
	e.GET("/orders", server.getOrders)
	e.POST("/order", server.createOrder)
	e.GET("/cart/:id", server.getCart)
	e.GET("/carts", server.getCarts)
	e.POST("/cart", server.createCart)
	e.GET("/user/:id", server.getUser)
	e.POST("/user", server.createUser)

	e.POST("/fileUpload", server.uploadProductImages)
	e.Static("/images", "images")
	e.Static("/", "web")
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

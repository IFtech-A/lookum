package apiserver

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/iftech-a/lookum/src/backend/internal/config"
	"github.com/iftech-a/lookum/src/backend/internal/store/sqlstore"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
	s *sqlstore.Store
	c *config.Config
}

//Start starts API server using the given configuration parameters
func Start(conf *config.Config) error {

	db, err := newDB(conf.GetDatabaseURL())

	if err != nil {
		return err
	}

	store := sqlstore.New(db)

	srv := newServer(store, conf)

	srv.configureRoutes()

	if conf.CertificatePath != "" && conf.PrivateKeyPath != "" {
		cert, err := tls.LoadX509KeyPair(conf.CertificatePath, conf.PrivateKeyPath)
		if err != nil {
			return err
		}
		httpsServer := &http.Server{
			Addr: fmt.Sprintf("%v:%v", conf.BindAddr, conf.PortHTTPS),
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
		Addr:              fmt.Sprintf("%v:%v", conf.BindAddr, conf.PortHTTP),
		ReadHeaderTimeout: time.Second * 10,
		IdleTimeout:       time.Second * 60,
	}

	return srv.e.StartServer(httpServer)
}

func newServer(store *sqlstore.Store, conf *config.Config) *Server {
	return &Server{
		s: store,
		e: echo.New(),
		c: conf,
	}
}

func (s *Server) configureRoutes() {

	s.e.GET("/product/:id", s.getProduct)
	s.e.GET("/products", s.getProducts)
	s.e.POST("/product", s.createProduct)
	s.e.GET("/order/:id", s.getOrder)
	s.e.GET("/orders", s.getOrders)
	s.e.POST("/order", s.createOrder)
	s.e.GET("/cart/:id", s.getCart)
	s.e.GET("/carts", s.getCarts)
	s.e.POST("/cart", s.createCart)
	s.e.GET("/user/:id", s.getUser)
	s.e.POST("/user", s.createUser)
	s.e.POST("/login", s.Login)

	s.e.POST("/fileUpload", s.uploadProductImages)
	s.e.Static("/images", "images")
	s.e.Static("/", "web")
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

package gql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/iftech-a/lookum/src/backend/internal/config"
	"github.com/iftech-a/lookum/src/backend/internal/store/sqlstore"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type GQLServer struct {
	s      *sqlstore.Store
	c      *config.Config
	schema graphql.Schema
}

func NewGQLServer(conf *config.Config) *GQLServer {
	return &GQLServer{
		c: conf,
	}
}

type GQuery struct {
	Query string `json:"query"`
}

func (s *GQLServer) ServeGQ(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.Body == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	query := &GQuery{}
	err = json.Unmarshal(body, query)
	if err != nil {
		logrus.Error(err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        s.schema,
		RequestString: query.Query,
	})

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)

}

func (s *GQLServer) Start() error {

	db, err := newDB(s.c.GetDatabaseURL())

	if err != nil {
		return err
	}

	s.s = sqlstore.New(db)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    s.newQuery(),
		Mutation: s.newMutation(),
	})
	if err != nil {
		return err
	}
	s.schema = schema

	e := echo.New()
	e.POST("/graphql", echo.WrapHandler(handler.New(&handler.Config{
		Pretty:   true,
		GraphiQL: true,
		Schema:   &s.schema,
	})))

	// mux := http.NewServeMux()
	// mux.Handle("/graphql", http.HandlerFunc(s.ServeGQ))

	// httpServer := &http.Server{
	// 	Addr:              fmt.Sprintf("%v:%v", s.c.BindAddr, s.c.PortHTTP),
	// 	ReadHeaderTimeout: time.Second * 10,
	// 	IdleTimeout:       time.Second * 60,
	// 	Handler:           mux,
	// }

	// return httpServer.ListenAndServe()
	return e.Start(fmt.Sprintf("%v:%v", s.c.BindAddr, s.c.PortHTTP))
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

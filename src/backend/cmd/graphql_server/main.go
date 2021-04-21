package main

import (
	env "github.com/Netflix/go-env"
	"github.com/iftech-a/lookum/src/backend/internal/config"
	gql "github.com/iftech-a/lookum/src/backend/internal/graphql"
	"github.com/sirupsen/logrus"
)

func main() {

	config := config.NewConfig()
	_, err := env.UnmarshalFromEnviron(config)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	server := gql.NewGQLServer(config)
	err = server.Start()
	logrus.Fatal(err.Error())
}

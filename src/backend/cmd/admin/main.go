package main

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/iftech-a/lookum/src/backend/internal/admin"
	"github.com/iftech-a/lookum/src/backend/internal/config"
)

func main() {

	config := config.NewConfig()
	_, err := env.UnmarshalFromEnviron(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := admin.NewAdminServer(config)

	err = server.Start()
	if err != nil {
		log.Fatal(err.Error())
	}

}

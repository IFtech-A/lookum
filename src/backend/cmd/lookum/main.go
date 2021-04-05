package main

import (
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
	"github.com/iftech-a/lookum/src/backend/internal/apiserver"
)

func main() {

	config := apiserver.NewConfig()
	_, err := env.UnmarshalFromEnviron(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	config.DatabaseURL = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName,
		config.DbSSLMode,
	)

	err = apiserver.Start(config)
	if err != nil {
		log.Fatal(err.Error())
	}

}

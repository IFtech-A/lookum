package main

import (
	"fmt"

	"github.com/iftech-a/lookum/src/backend/internal/apiserver"
)

func main() {

	config := apiserver.NewConfig()

	err := apiserver.Start(config)
	if err != nil {
		fmt.Print(err)
		return
	}

}

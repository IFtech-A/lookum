package main

import (
	"fmt"
	"lookum/internal/apiserver"
)

func main() {

	config := apiserver.NewConfig()

	err := apiserver.Start(config)
	if err != nil {
		fmt.Print(err)
		return
	}

}

package main

import (
	. "GoWebProxy/server/config"
	"fmt"
)

func main() {
	var cnf Config
	config,err:= cnf.Load("config/config.yaml")
	if err!=nil {
		panic(err)
	}
	fmt.Printf("config: %v  error: %v",config,err)
}

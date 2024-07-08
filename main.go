package main

import (
	"db-kv-go/routes"
	"db-kv-go/server"
	"fmt"
)

func main() {
	fmt.Println("run")
	s := server.NewServer()
	r := routes.RunRoutes(s)
	r.Run(":3010")
}

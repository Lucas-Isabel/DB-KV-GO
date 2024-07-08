package main

import (
	"fmt"

	"github.com/Lucasbyte/DB-KV-GO/routes"
	"github.com/Lucasbyte/DB-KV-GO/server"
)

func main() {
	fmt.Println("run")
	s := server.NewServer()
	r := routes.RunRoutes(s)
	r.Run(":3010")
}

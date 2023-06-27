package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "mark-server/docs"
	"mark-server/pkg/server"
)

func main() {
	server.Start()
}

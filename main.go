package main

import (
	"price/internal/server"
)

func main() {
	server := server.NewApp()
	server.Run()
}

package main

import (
	"flag"
	"fmt"
	"forum/internal/delivery"
	"forum/internal/repository"
	"forum/internal/server"
	"forum/internal/service"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := repository.OpenSqliteDB("store.db")
	if err != nil {
		log.Fatalf("error while opening db: %s", err)
	}

	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := delivery.NewHandler(service)
	server := new(server.Server)

	fmt.Printf("Starting server at addr %s\nhttp://localhost%s/\n", *addr, *addr)

	if err := server.Run(*addr, handler.InitRoutes()); err != nil {
		log.Fatalf("error while running the server: %s", err.Error())
	}
}

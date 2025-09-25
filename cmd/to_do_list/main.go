package main

import (
	"fmt"
	"log"
	"toDoList/internal"
	"toDoList/internal/repository/db"
	"toDoList/internal/repository/inmemory"
	"toDoList/internal/server"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// конфигураця приложения
	fmt.Println("To-do-list Api is starting")
	cfg := internal.ReadConfig()
	
	var database server.Storage

	// конфигураця и создание хранилища
	postgresDB, err := db.NewStorage(cfg.DNS)
	if err != nil {
		log.Printf("Postgres недоступен (%v), используем in-memory storage", err)
		inmemoryDB := inmemory.NewInMemoryStorage()
		database = inmemoryDB
	} else {
		database = postgresDB
		// запуск миграции
		if err := db.Migrations(cfg.DNS, cfg.MigratePath); err != nil {
			log.Fatal(err)
		}
	}

	// конфигурация и запуск веб-сервера
	srv := server.NewServer(cfg, database)
	if err := srv.Run(); err != nil {
		panic(err)
	}
}

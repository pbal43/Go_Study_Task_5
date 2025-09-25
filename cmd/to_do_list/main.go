package main

import (
	"fmt"
	"log"
	"toDoList/internal"
	"toDoList/internal/repository/db"
	"toDoList/internal/server"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// конфигураця приложения
	fmt.Println("To-do-list Api is starting")
	cfg := internal.ReadConfig()

	// конфигураця и создание хранилища
	database, err := db.NewStorage(cfg.DNS)
	if err != nil {
		log.Fatal(err)
	}

	// запуск миграции
	if err := db.Migrations(cfg.DNS, cfg.MigratePath); err != nil {
		log.Fatal(err)
	}

	// конфигурация и запуск веб-сервера
	srv := server.NewServer(cfg, database)
	if err := srv.Run(); err != nil {
		panic(err)
	}
}

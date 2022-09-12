package main

import (
	"embed"
	"fmt"

	"github.com/nekitvand/to_do_service/internal/config"
	"github.com/nekitvand/to_do_service/internal/pkg/db"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

var embedMigrations embed.FS

func main() {
	if err := config.ReadConfigYaml("config.yaml"); err != nil {
		fmt.Println(err)
	}
	cfg := config.GetConfig()

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	goose.SetBaseFS(embedMigrations)

	err = goose.Up(conn.DB, "migrations")
	if err != nil {
		fmt.Println(err)
	}
}
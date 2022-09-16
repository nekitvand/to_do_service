package main

import (
	"fmt"
	"github.com/nekitvand/to_do_service/internal/config"
	"github.com/nekitvand/to_do_service/internal/pkg/db"
	"github.com/nekitvand/to_do_service/migrations"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

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

	goose.SetBaseFS(migrations.EmbedFS)

	err = goose.Up(conn.DB, ".")
	if err != nil {
		fmt.Println(err)
	}
}
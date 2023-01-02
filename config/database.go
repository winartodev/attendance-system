package config

import (
	"context"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/winartodev/attencande-system/ent"
)

func NewDatabase(cfg Config) (*ent.Client, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	client, err := ent.Open(cfg.Database.Driver, dsn)
	if err != nil {
		return client, err
	}

	return client, err
}

func AutoMigrate(client *ent.Client) error {
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}

	return nil
}

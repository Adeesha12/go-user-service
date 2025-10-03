package tests

import (
	"testing"

	"go-user-service/config"
	"go-user-service/data_base"
)

func TestDb(t *testing.T) {

	cfg := config.LoadConfig()

	db := data_base.ConnectDb(cfg)
	defer db.Close()

}

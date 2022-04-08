package data

import (
	"realworld_02/internal/conf"
	"testing"
)

func TestNewDB(t *testing.T) {
	c := &conf.Data{
		Database: &conf.Data_Database{
			Dsn: "root:123456@tcp(127.0.0.1:3305)/realworld?charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
	NewDB(c)
}
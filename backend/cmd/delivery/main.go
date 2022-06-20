package main

import (
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/cfg"
	"github.com/RuslanNadyrshyn/GoLangCourse/backend/internal/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := cfg.NewConfig(false)
	server.Start(config)
}

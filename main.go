package main

import (
	"embed"
	"fmt"

	"github.com/HamzaDLM/go_vue/config"
	"github.com/HamzaDLM/go_vue/container"
	"github.com/HamzaDLM/go_vue/database"
	"github.com/HamzaDLM/go_vue/logger"
	"github.com/HamzaDLM/go_vue/middleware"
	"github.com/HamzaDLM/go_vue/router"

	"github.com/gofiber/fiber/v2"
)

//go:embed config/app.*.yaml
var AppConfigFile embed.FS

//go:embed public/*
var staticFile embed.FS

const banner = `Started server
-------------------------------------
WEB: http://127.0.0.1:%s
API: http://127.0.0.1:%s/api/v1
DOC: http://127.0.0.1:%s/api/v1/docs
-------------------------------------
`

// @title Doc
// @version 0.0.1
// description API specs

// @license.name MIT
// @license.url https://opensource.org/licenses/mit-license.php

// @host localhost:6969
// @BasePath /api/v1
func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	conf, env := config.Load(AppConfigFile)
	logger := logger.Get()

	Banner := fmt.Sprintf(banner, conf.App.Port, conf.App.Port, conf.App.Port)
	logger.Info(Banner)

	db, err := database.New(&database.DatabaseConfig{
		Driver:   conf.Database.Driver,
		Host:     conf.Database.Host,
		Username: conf.Database.Username,
		Password: conf.Database.Password,
		Port:     conf.Database.Port,
		Database: conf.Database.Dbname,
	}, logger)

	if err != nil || db == nil {
		logger.Fatal("Couldn't connect to database")
	}

	container := container.NewContainer(conf, env, logger, db)

	middleware.InitLoggerMiddleware(app, container)
	middleware.StaticContentsMiddleware(app, container, staticFile)

	api := app.Group("/api")

	v1 := api.Group("/v1")
	router.Init(v1, container)

	app.Listen(":" + conf.App.Port)
}

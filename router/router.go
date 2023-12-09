package router

import (
	"github.com/HamzaDLM/go_vue/container"
	controller "github.com/HamzaDLM/go_vue/controller"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// When router is initialized, call the endpoints
func Init(app fiber.Router, container container.Container) {
	// Should be kept at the start of the stack
	setRecover(app)

	setHealthController(app, container)
	setMetrics(app)
	setPing(app, container)
	setPanic(app)
	setSwagger(app, container)
	// Grouped business logic endpoints
	registerEmployees(app, container)

	// Should be kept at the end of the stack
	setNotFound(app)
}

// Business logic grouped endpoints

func registerEmployees(app fiber.Router, container container.Container) {
	api := app.Group("/employees")

	db := container.GetDb()

	api.Get("/", controller.GetAllEmployees(db))
	api.Get("/:id", controller.GetEmployeeById(db))
}

// Utility endpoints

func setHealthController(app fiber.Router, container container.Container) {
	health := controller.NewHealthController(container)
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return health.GetHealthCheck(ctx)
	})
}

func setMetrics(app fiber.Router) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics"}))
}

func setPing(app fiber.Router, container container.Container) {
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Clugie")
	})
}

func setNotFound(app fiber.Router) {
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusNotFound)
	})
}

// Panic test route, this brings up an error
func setPanic(app fiber.Router) {
	app.Get("/panic", func(ctx *fiber.Ctx) error {
		panic("Hi, I'm a panic error!")
	})
}

func setRecover(app fiber.Router) {
	app.Use(recover.New())
}

func setSwagger(app fiber.Router, container container.Container) {
	if container.GetConfig().Swagger.Enabled {
		cfg := swagger.Config{
			BasePath: "/api/v1/",
			FilePath: "./docs/swagger.json",
			Path:     "docs",
			Title:    "Swagger API Docs",
		}
		app.Use(swagger.New(cfg))
	}
}

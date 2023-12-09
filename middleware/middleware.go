package middleware

import (
	"embed"
	"github.com/HamzaDLM/go_vue/container"
	"github.com/HamzaDLM/go_vue/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strconv"
	// "strings"
	"time"
)

// Initializes the listed middlewares.
func InitLoggerMiddleware(app fiber.Router, container container.Container) {
	app.Use(SetCORSConfig(app))

	app.Use(RequestLoggerMiddleware(container))
	// app.Use(ActionLoggerMiddleware(container))
}

func SetCORSConfig(app fiber.Router) fiber.Handler {
	return cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     "*",
		AllowMethods: "DEL",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

// Middleware for loading static files.
func StaticContentsMiddleware(app fiber.Router, container container.Container, staticFile embed.FS) {
	conf := container.GetConfig()
	logger := logger.Get()

	if conf.StaticContents.Enabled {
		app.Static("/", "./public", fiber.Static{
			Compress:      true,
			ByteRange:     true,
			Browse:        true,
			Index:         "index.html",
			CacheDuration: 10 * time.Second,
			MaxAge:        3600,
		})
		logger.Info("Serving static contents from ./public folder.")
	}
}

// Middleware for logging requests
func RequestLoggerMiddleware(container container.Container) fiber.Handler {
	return func(c *fiber.Ctx) error {
		startTime := time.Now()

		// Handle request
		err := c.Next()

		req := c.Request()
		res := c.Response()
		logger := logger.Get()

		logger.Info(c.IP() + " - " + string(req.Header.Method()) + " " + string(req.RequestURI()) + " " + strconv.Itoa(res.StatusCode()) + " " + time.Since(startTime).String()) // zap.String("ip", c.IP()),
		// zap.String("method", string(req.Header.Method())),
		// zap.String("uri", string(req.RequestURI())),
		// zap.Int("status", res.StatusCode()),
		// zap.String("latency", time.Since(startTime).String()),

		return err
	}
}

func ActionLoggerMiddleware(container container.Container) {
	println("action logger middlewear")
}



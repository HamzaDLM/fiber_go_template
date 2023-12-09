package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HamzaDLM/go_vue/config"
	"github.com/HamzaDLM/go_vue/container"
	"github.com/HamzaDLM/go_vue/database"
	"github.com/HamzaDLM/go_vue/logger"
	"github.com/HamzaDLM/go_vue/router"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func prepTestRouter() *fiber.App {
	app := fiber.New()

	conf := &config.Config{
		Database: struct {
			Driver    string `default:"sqlite3"`
			Host      string `default:"book.db"`
			Port      string
			Dbname    string
			Username  string
			Password  string
			Migration bool `default:"false"`
		}{
			Driver:    "sqlite3",
			Host:      "book.db",
			Port:      "",
			Dbname:    "",
			Username:  "",
			Password:  "",
			Migration: false,
		},
		Swagger: struct {
			Enabled bool "default:\"false\""
			Path    string
		}{
			Enabled: false,
			Path:    "",
		},
	}
	logger := logger.Get()

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
	container := container.NewContainer(conf, "test", logger, db)

	api := app.Group("/api")

	v1 := api.Group("/v1")
	router.Init(v1, container)

	return app
}

func performRequest(method, target string, app *fiber.App) *http.Response {
	r := httptest.NewRequest(method, target, nil)
	resp, _ := app.Test(r, -1)
	return resp
}

func TestAPIV1(t *testing.T) {
	app := prepTestRouter()
	baseRoute := "/api/v1"

	tests := []struct {
		description string

		// Test input
		method string
		route  string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "Health check",
			method:        "GET",
			route:         "/health",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "Ok",
		},
		{
			description:   "Ping check",
			method:        "GET",
			route:         "/ping",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "Clugie",
		},
		{
			description:   "Metrics check",
			method:        "GET",
			route:         "/metrics",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "Ok",
		},
		{
			description:   "Panic check",
			method:        "GET",
			route:         "/panic",
			expectedError: false,
			expectedCode:  500,
			expectedBody:  "Ok",
		},
		{
			description:   "404 Check",
			method:        "GET",
			route:         "/arbitraryroute",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Not found",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			resp := performRequest(test.method, baseRoute+test.route, app)
			require.Equal(t, test.expectedCode, resp.StatusCode)

			// if test.expectedError {
			// 	continue
			// }
			_, err := io.ReadAll(resp.Body)
			require.Nilf(t, err, test.description)

			// require.Equal(t, test.expectedBody, string(body))
		})
	}
}

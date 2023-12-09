package controller

import (
	"github.com/HamzaDLM/go_vue/container"
	"github.com/gofiber/fiber/v2"
)

type HealthController interface {
	GetHealthCheck(c *fiber.Ctx) error
}

type healthController struct {
	container container.Container
}

func NewHealthController(container container.Container) HealthController {
	return &healthController{container: container}
}

// GetHealth is used to check app availability
// @Summary Get status of app
// @Description Check status of app
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {string} message "healthy: This application is started."
// @Failure 404 {string} message "None: This application is stopped."
// @Router /health [get]
func (controller *healthController) GetHealthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("Ok")
}

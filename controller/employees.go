package controller

import (
	"github.com/HamzaDLM/go_vue/database"
	"github.com/gofiber/fiber/v2"
)

type Employee struct {
	EmployeeId int
	LastName   string
	FirstName  string
}

// GetAllEmployees
// @Summary Get all employees
// @Description Get all employees
// @Tags Employees
// @Accept  json
// @Produce  json
// @Success 200 {string} message ""
// @Failure 404 {string} message ""
// @Router /employees [get]
func GetAllEmployees(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var employee []Employee

		db.Raw("SELECT EmployeeId, LastName, FirstName FROM employees").Scan(&employee)

		return ctx.JSON(employee)
	}
}

// GetEmployeeById is used to get an employee by id
// @Summary Get employee by id
// @Description Get employee by id
// @Tags Employees
// @Accept  json
// @Produce  json
// @Success 200 {string} message ""
// @Failure 404 {string} message ""
// @Router /employees/{id} [get]
func GetEmployeeById(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		id := ctx.Params("id")
		var employee Employee

		db.Raw("SELECT EmployeeId, LastName, FirstName FROM employees WHERE EmployeeId = ?", id).Scan(&employee)

		return ctx.JSON(employee)
	}
}

package main

import "github.com/gofiber/fiber/v2"

func registerRoutes(app *fiber.App) {
	// Healthcheck endpoint
	// @Summary     Healthcheck
	// @Accept       json
	// @Produce      json
	// @Success      200  {string}  "Healthy"
	// @Failure      500
	// @Router       /healthz [get]
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Healthy")
	})

	app.Get("/v1/foo", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "foo"})
	})

	app.Get("/v1/bar", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "bar"})
	})

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("The requested path does not exist on this resource")
	})
}

package main

import (
	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	port := GetEnvOrDefaultString("PORT", "3000")
	return ":" + port
}

func main() {
	app := fiber.New()
	app.Use(limiterConfig())

	app.Get("/", handleIndex)
	api := app.Group("api", validateApiKey)
	api.Get("staff", handleSyncRajoGetStaff)

	if err := app.Listen(getPort()); err != nil {
		panic(err)
	}
}

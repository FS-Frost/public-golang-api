package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func getPort() string {
	port := GetEnvOrDefaultString("PORT", "3000")
	return ":" + port
}

func main() {
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{

		Max:        GetEnvOrDefaultInt("MAX_REQUESTS_PER_SECOND", 10),
		Expiration: time.Duration(GetEnvOrDefaultInt("SECONDS_EXPIRATION_RECORD", 30)) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			c.Status(fiber.StatusTooManyRequests)
			return c.JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
		Storage: limiter.ConfigDefault.Storage,
	}))

	app.Get("/", handleIndex)
	api := app.Group("api", validateApiKey)
	syncrajo := api.Group("syncrajo", validateSyncRajoKey)
	syncrajo.Get("", handleSyncRajoGetIndex)
	syncrajo.Get("staff", handleSyncRajoGetStaff)

	app.Listen(getPort())
}

func handleIndex(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

func handleSyncRajoGetIndex(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"mensaje":     "¡Bienvenido a la API oficial de SyncRajo Fansub!",
		"visitanosEn": "http://www.syncrajo.net",
	})
}

func handleSyncRajoGetStaff(c *fiber.Ctx) error {
	type miembro struct {
		Apodo string `json:"apodo"`
		Pais  string `json:"pais"`
	}

	miembros := []miembro{
		{
			Apodo: "[FS] Frost",
			Pais:  "Chile",
		},
	}

	return c.JSON(fiber.Map{
		"staff": miembros,
	})
}

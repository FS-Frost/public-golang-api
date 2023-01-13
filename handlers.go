package main

import "github.com/gofiber/fiber/v2"

func handleIndex(c *fiber.Ctx) error {
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

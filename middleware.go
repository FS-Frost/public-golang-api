package main

import "github.com/gofiber/fiber/v2"

func validateApiKey(c *fiber.Ctx) error {
	keyName := GetEnvOrDefaultString("API_KEY_NAME", "some-api-key")
	keyValue := GetEnvOrDefaultString(keyName, "some-api-key-value")
	reqKey := c.Get(keyName)
	if keyValue != reqKey {
		c.Status(401)
		return c.JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next()
}

func validateSyncRajoKey(c *fiber.Ctx) error {
	keyName := GetEnvOrDefaultString("SYNCRAJO_KEY_NAME", "some-syncrajo-key")
	keyValue := GetEnvOrDefaultString(keyName, "some-syncrajo-key-value")
	reqKey := c.Get(keyName)
	if keyValue != reqKey {
		c.Status(401)
		return c.JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next()
}

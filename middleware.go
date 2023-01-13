package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func validateApiKey(c *fiber.Ctx) error {
	keyName := GetEnvOrDefaultString("API_KEY_NAME", "SOME_API_KEY")
	keyValue := GetEnvOrDefaultString(keyName, "some-api-key-value")
	reqKeyName := strings.ReplaceAll(keyName, "_", "-")
	reqKey := c.Get(reqKeyName)
	if keyValue != reqKey {
		c.Status(401)
		return c.JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next()
}

func validateSyncRajoKey(c *fiber.Ctx) error {
	keyName := GetEnvOrDefaultString("SYNCRAJO_KEY_NAME", "SOME_SYNCRAJO_KEY")
	keyValue := GetEnvOrDefaultString(keyName, "some-syncrajo-key-value")
	reqKeyName := strings.ReplaceAll(keyName, "_", "-")
	reqKey := c.Get(reqKeyName)
	if keyValue != reqKey {
		c.Status(401)
		return c.JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next()
}

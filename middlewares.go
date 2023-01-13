package main

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func limiterConfig() func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
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
	})
}

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

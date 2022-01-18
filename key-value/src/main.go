package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	kv := make(map[string]string)

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/:key", func(c *fiber.Ctx) error {
		key := c.Params("key")
		value := kv[key]
		if value == "" {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Status(fiber.StatusOK).JSON(value)
	})

	app.Post("/:key", func(c *fiber.Ctx) error {
		key := c.Params("key")
		value := c.Body()
		kv[key] = string(value)

		fmt.Println(kv)

		return c.SendStatus(fiber.StatusAccepted)
	})

	app.Listen(":3000")
}

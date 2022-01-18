package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	keyValueService := os.Getenv("KEY_VALUE_SERVICE")
	if keyValueService == "" {
		panic("KEY_VALUE_SERVICE is not set")
	}

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/:key", func(c *fiber.Ctx) error {
		key := c.Params("key")
		res, err := http.Get("http://" + keyValueService + "/" + key)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		bodyString := string(body)

		if res.StatusCode == fiber.StatusNotFound {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.Status(fiber.StatusOK).SendString(bodyString)
	})

	app.Listen(":3000")
}

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	type GetNameTimeRequest struct {
		Name string `json:"name"`
	}

	app.Get("/", func(c fiber.Ctx) error {
		now := time.Now()
		return c.JSON(fiber.Map{
			"message":        "My name is Olivia",
			"timestamp":      now.UnixMilli(),
			"formatted_time": now.Format("15:04:05 01-02-2006"),
		})
	})

	app.Post("/", func(c fiber.Ctx) error {
		var body GetNameTimeRequest

		if len(c.Body()) > 0 {
			if err := c.Bind().Body(&body); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"error": "invalid request body",
					"example": fiber.Map{
						"name": "Olivia",
					},
				})
			}
		}

		if body.Name == "" {
			body.Name = "Olivia"
		}

		now := time.Now()
		return c.JSON(fiber.Map{
			"message":        fmt.Sprintf("My name is %s", body.Name),
			"timestamp":      now.UnixMilli(),
			"formatted_time": now.Format("15:04:05 01-02-2006"),
		})
	})

	log.Fatal(app.Listen(":3000"))
}

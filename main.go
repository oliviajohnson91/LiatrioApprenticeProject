package main

import (
    "log"
	"time"
    "github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
		now := time.Now()
        return c.JSON(fiber.Map{
			"message": "My name is Olivia",
			"timestamp": now.Unix(),
		})
    })

    log.Fatal(app.Listen(":3000"))
}
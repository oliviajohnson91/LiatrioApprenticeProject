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
        Name  string `json:"name"`
    }

    app.Get("/", func(c fiber.Ctx) error {
        var body GetNameTimeRequest

        if err := c.Bind().Body(&body); err != nil {
            return c.Status(400).JSON(fiber.Map{
                "error": "invalid request body",
                "example": fiber.Map{
                    "name": "Olivia",
                },
            })
        }

        if body.Name == "" {
            return c.Status(400).JSON(fiber.Map{
                "error": "name is required",
            })
        }

		now := time.Now()
        return c.JSON(fiber.Map{
			"message": fmt.Sprintf("My name is %s", body.Name),
			"timestamp": now.Unix(),
		})
    })

    log.Fatal(app.Listen(":3000"))
}
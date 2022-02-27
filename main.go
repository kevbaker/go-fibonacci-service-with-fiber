package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevbaker/go-service-rest-fib-two/responses"
	"github.com/kevbaker/go-service-rest-fib-two/utils"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!")
	})

	app.Get("/fibonacci", func(c *fiber.Ctx) error {
		// var count = int c.Query("count")
		// fmt.Println(string(c.Request().URI().QueryString()))
		fl := utils.GenerateFibonacciList(200)
		response := responses.Int64ListResponse{Data: fl, Status: 200}
		return c.SendString(string(response.Json()))
	})

	app.Listen(":3000")
}

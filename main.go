package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevbaker/go-service-rest-fib-two/lib/responses"
	"github.com/kevbaker/go-service-rest-fib-two/lib/utils"
)

func getContextQueryInt(c *fiber.Ctx, queryKey string, defaultInt int) int {
	valueString := string(c.Query(queryKey))
	value := defaultInt
	if valueString != "" {
		value, _ = strconv.Atoi(valueString)
	}
	return value
}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!")
	})

	app.Get("/fibonacci", func(c *fiber.Ctx) error {
		count := getContextQueryInt(c, "count", 10)

		fmt.Println(count)
		fl := utils.GenerateFibonacciList(count)
		response := responses.Int64ListResponse{Data: fl, Status: 200}
		return c.SendString(string(response.Json()))
	})

	app.Listen(":3000")
}

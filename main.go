package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kevbaker/go-service-rest-fib-two/responses"
	fib "github.com/t-pwk/go-fibonacci"
)

func GenerateFibonacciList(numberCount int) []uint64 {
	var fibonacciList []uint64
	fmt.Println(fibonacciList)

	for i := 0; i < numberCount; i++ {
		fibvalue := fib.Fibonacci(uint(i))
		fibonacciList = append(fibonacciList, fibvalue)
	}
	return fibonacciList
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/fibonacci", func(c *fiber.Ctx) error {
		// fmt.Printf("Fibonacci:" + string(fib.Fibonacci(10)))
		// out := "Fibonacci:"
		// fibval := fib.Fibonacci(1)
		// fmt.Print(fibval)
		// return c.SendString(out)
		fl := GenerateFibonacciList(100)
		response := responses.Int64ListResponse{Data: fl, Status: 200}
		fmt.Println(string(response.Json()))
		return c.SendString(string(response.Json()))
	})

	app.Listen(":3000")
}

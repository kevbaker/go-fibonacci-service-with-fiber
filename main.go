package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	fib "github.com/t-pwk/go-fibonacci"
)

// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}

// 	return fib(n-1) + fib(n-2)
// }

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
		out := "Fibonacci:"
		fibval := fib.Fibonacci(1)
		fmt.Print(fibval)
		return c.SendString(out)
	})

	app.Listen(":3000")
}

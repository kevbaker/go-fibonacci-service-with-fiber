package utils

import (
	"fmt"

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

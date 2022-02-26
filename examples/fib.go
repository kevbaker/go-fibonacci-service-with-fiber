package main

import (
	"fmt"

	"github.com/kevbaker/go-service-rest-fib-two/responses"

	fib "github.com/t-pwk/go-fibonacci"
)

// type Int64ListResponse struct {
// 	Data   []uint64
// 	Status int
// 	hideMe string
// }

// func (r responses.Int64ListResponse) Json() []byte {
// 	rj, err := json.MarshalIndent(r, "", "  ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return rj
// }

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
	fl := GenerateFibonacciList(100)
	response := responses.Int64ListResponse{Data: fl, Status: 200}
	fmt.Println(string(response.Json()))
}

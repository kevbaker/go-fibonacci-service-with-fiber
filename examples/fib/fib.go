package main

import (
	"fmt"

	"github.com/kevbaker/go-service-rest-fib-two/responses"
	"github.com/kevbaker/go-service-rest-fib-two/utils"
)

func main() {
	fl := utils.GenerateFibonacciList(100)
	response := responses.Int64ListResponse{Data: fl, Status: 200}
	fmt.Println(string(response.Json()))
}

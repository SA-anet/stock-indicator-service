package main

import (
	"fmt"

	"example.com/server"
)

func main() {
	fmt.Println("Hello, World!")
	server.Hello(13)
	resultMin := server.CalculateIndicator(13, "MIN")
	resultMax := server.CalculateIndicator(13, "MAX")
	resultAvg := server.CalculateIndicator(13, "AVG")
	fmt.Println(resultMax, resultMin, resultAvg)
}

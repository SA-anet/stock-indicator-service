package main

import (
	"fmt"

	"example.com/server"
)

func main() {
	fmt.Println("Hello, World!")
	server.Hello(13)
	resultMin := server.CumulativeOperation(13, "MIN")
	resultMax := server.CumulativeOperation(13, "MAX")
	resultAvg := server.CumulativeOperation(13, "AVG")
	fmt.Println(resultMax, resultMin, resultAvg)
}

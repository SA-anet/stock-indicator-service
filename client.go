package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"example.com/example-app/stock_service"
	"google.golang.org/grpc"
)

var (
	stockName    = flag.String("stockName", "ANET", "The name of the stock (default: ANET)")
	days         = flag.Int64("days", 2, "the number of days to look back")
	average      = flag.Bool("average", false, "Get the average of the prices")
	min          = flag.Bool("min", false, "Get the minimum value of the prices")
	max          = flag.Bool("max", false, "Get the maximum value of the prices")
	stdDeviation = flag.Bool("std", false, "Get the standard deviation of the prices")
)

func getIndicator() (stock_service.StockRequest_Indicator, error) {
	flagCount := 0
	var indicator stock_service.StockRequest_Indicator
	flags := []*bool{average, min, max, stdDeviation}
	indicators := []stock_service.StockRequest_Indicator{
		stock_service.StockRequest_AVERAGE,
		stock_service.StockRequest_MIN,
		stock_service.StockRequest_MAX,
		stock_service.StockRequest_STD_DEVIATION,
	}

	for ind, flag := range flags {
		if *flag {
			flagCount++
			indicator = indicators[ind]
		}
		if flagCount > 1 {
			return stock_service.StockRequest_AVERAGE, errors.New("Cannot do more than one action at once")
		}

	}

	return indicator, nil
}

func main() {
	flag.Parse()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5003", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error creating a connection")
	}
	defer conn.Close()

	client := stock_service.NewStockServiceClient(conn)
	indicator, err := getIndicator()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	req := stock_service.StockRequest{
		StockName: *stockName,
		NDays:     *days,
		Indicator: indicator,
	}
	res, err := client.GetStockDetails(context.Background(), &req)
	if err != nil {
		fmt.Println("Error getting response")
	}

	fmt.Printf("Response price: %v\n", res.Price)
}

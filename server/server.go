package server

import (
	"fmt"
	"strconv"

	"math"

	goyhfin "github.com/svarlamov/goyhfin"
)

func Hello(n int) {
	resp, err := goyhfin.GetTickerData("ANET", strconv.Itoa(n)+"d", goyhfin.OneDay, false)
	if err != nil {
		// NOTE: For library-specific errors, you can check the err against the errors exposed in goyhfin/errors.go
		fmt.Println("Error fetching Yahoo Finance data:", err)
		panic(err)
	}
	for ind := range resp.Quotes {
		fmt.Println("The day's close was", resp.Quotes[ind].Close, "on the", resp.Quotes[ind].OpensAt.Day(), "day of", resp.Quotes[ind].OpensAt.Month(), "of", resp.Quotes[ind].OpensAt.Year())
	}
}

type Tickers struct {
	tickers []float64
}

func getNDayData(n int) Tickers {
	ticker := Tickers{tickers: make([]float64, n)}
	resp, err := goyhfin.GetTickerData("ANET", strconv.Itoa(n)+"d", goyhfin.OneDay, false)
	if err != nil {
		fmt.Println("Error fetching Yahoo Finance data:", err)
		panic(err)
	}
	for i := range resp.Quotes {
		ticker.tickers[i] = resp.Quotes[i].Close
	}
	return ticker
}

func CumulativeOperation(n int, operation string) float64 {
	ticker := getNDayData(n).tickers
	switch operation {
	case "MAX":
		maxTicker := -math.MaxFloat64
		for i := 0; i < len(ticker); i++ {
			maxTicker = math.Max(maxTicker, ticker[i])
		}
		return maxTicker
	case "MIN":
		minTicker := math.MaxFloat64
		for i := 0; i < len(ticker); i++ {
			fmt.Println(i, ticker[i])
			minTicker = math.Min(minTicker, ticker[i])
		}
		return minTicker
	case "AVG":
		var avg float64 = 0
		for i := 0; i < len(ticker); i++ {
			avg += ticker[i]
		}
		avg /= float64(n)
		return avg
	default:
		return -1
	}
}

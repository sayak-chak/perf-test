package main

import (
	"fmt"
	"perf-test/collector"
	"time"
)

func main() {
	start := time.Now()
	collectr := collector.New(4)
	collectr.AddTest("GET", "http://google.com", "", "", nil, 200)

	collectr.PrintTestResults()
	fmt.Println("Time taken =", time.Since(start).Seconds())
}

package main

import (
	"perf-test/collector"
)

func main() {
	collectr := collector.New()
	collectr.AddTest("GET", "http://google.com", "", "", nil, 200)

	collectr.PrintTestResults()
}

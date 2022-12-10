package collector

import (
	"fmt"
	testRunner "perf-test/runner"
)

type testSummary struct {
	runNo                  int
	latency                float64
	totalNumberOfTests     int
	totalNnumberOfFailures int
	numberOfSetupFailures  int
}

type testReport struct {
	summary         testSummary
	detailedResults []testRunner.Result
}

func (t *testReport) print() {
	fmt.Printf("\n----------------------Report for run %d is ----------------------\n", t.summary.runNo)
	fmt.Printf("Total number of tests = %d\n", t.summary.totalNumberOfTests)
	fmt.Printf("Total number of failures = %d\n", t.summary.totalNnumberOfFailures)
	fmt.Printf("Total number of setup failures = %d\n", t.summary.numberOfSetupFailures)
	fmt.Printf("Average latency = %f seconds\n", t.summary.latency)

	fmt.Printf("\n------------------------------------------------------------------\n")
}

func generateTestReport(detailedResults []testRunner.Result, run int) testReport { //TODO: break into smaller methods
	summary := testSummary{
		totalNumberOfTests: len(detailedResults),
		runNo:              run,
	}
	cumulativeLatency := 0.0
	for _, result := range detailedResults {
		if result.IsSetupSuccesful() {
			cumulativeLatency += result.Latency()
			if result.Error() != nil {
				summary.totalNnumberOfFailures++
			}
		} else {
			summary.numberOfSetupFailures++
			summary.totalNnumberOfFailures++
		}
	}

	if summary.totalNnumberOfFailures != summary.totalNumberOfTests {
		summary.latency = cumulativeLatency / float64(summary.totalNumberOfTests-summary.numberOfSetupFailures)
	}

	return testReport{
		summary:         summary,
		detailedResults: detailedResults,
	}
}

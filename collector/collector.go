package collector

import (
	testRunner "perf-test/runner"
)

const (
	defaultRunnerCount = 5
)

type runner interface {
	AddTest(methodType, linkAddress, headers, reqBody string, reqHeaders map[string][]string, expectedResp int)
	RunTests() []testRunner.Result
}

type collector struct {
	runners     []runner
	testReports []testReport
}

//Returns a new collector with five test runners
func New(noOfTimesEachTestIsToBeExecutedInARun int) *collector {
	cllctr := collector{}
	cllctr.addRunners(defaultRunnerCount, noOfTimesEachTestIsToBeExecutedInARun)
	return &cllctr
}

// Adds this test to all teest runners
func (c *collector) AddTest(methodType, linkAddress, headers, reqBody string, reqHeaders map[string][]string, expectedResp int) { //TODO: refactor
	for idx := range c.runners {
		c.runners[idx].AddTest(methodType, linkAddress, headers, reqBody, reqHeaders, expectedResp)
	}
}

//Prints the test results from all runners
func (c *collector) PrintTestResults() {
	// prints the test results
	c.executeRunners()
	for idx := range c.testReports {
		c.testReports[idx].print()
	}
}

func (c *collector) executeRunners() {
	for idx := range c.runners {
		c.executeRunner(idx)
	}
}

func (c *collector) executeRunner(idx int) {
	c.testReports = append(c.testReports, generateTestReport(c.runners[idx].RunTests(), idx))
}

func (c *collector) addRunners(runnerCount, noOfTimesEachTestIsToBeExecutedInARun int) {
	for i := 0; i < runnerCount; i++ {
		c.addRunner(noOfTimesEachTestIsToBeExecutedInARun)
	}
}

func (c *collector) addRunner(noOfTimesEachTestIsToBeExecutedInARun int) {
	c.runners = append(c.runners, testRunner.New(noOfTimesEachTestIsToBeExecutedInARun))
}

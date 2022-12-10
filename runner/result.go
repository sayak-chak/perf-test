package runner

import (
	runner_errors "perf-test/runner/errors"
	"time"
)

type Result struct {
	err              error
	latency          float64
	isSetupSuccesful bool
}

func (r *Result) IsSetupSuccesful() bool {
	return r.isSetupSuccesful
}

func (r *Result) Latency() float64 {
	return r.latency
}

func (r *Result) Error() error {
	return r.err
}

func getResultOnSetupFailure(setupFailureErr error) *Result {
	return &Result{
		err: setupFailureErr,
	}
}

func getResultPostSetupSuccess(timeTaken time.Duration, actualRspCode, expectedRspCode int) *Result {
	if actualRspCode != expectedRspCode {
		return &Result{
			latency:          timeTaken.Seconds(),
			err:              runner_errors.StatusCodeMismatchErr(actualRspCode, expectedRspCode),
			isSetupSuccesful: true,
		}
	}

	return &Result{
		latency:          timeTaken.Seconds(),
		isSetupSuccesful: true,
	}
}

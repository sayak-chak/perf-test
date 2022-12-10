package runner

import (
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

type runner struct {
	testList     []testAttributes
	client       *http.Client
	noOfTestRuns int
}

func New(noOfTestRuns int) *runner {
	return &runner{
		client:       &http.Client{},
		noOfTestRuns: noOfTestRuns,
	}
}

func (w *runner) AddTest(methodType, linkAddress, headers, reqBody string, reqHeaders map[string][]string, expectedResp int) { //TODO: refactor
	w.testList = append(w.testList, *newTestArributes(methodType, linkAddress, headers, reqBody, reqHeaders, expectedResp))
}

func (w *runner) RunTests() []Result {
	var results []Result
	var wg sync.WaitGroup
	for i := 0; i < w.noOfTestRuns; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for _, testDetails := range w.testList {
				results = append(results, w.runTestFor(testDetails))
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()

	return results
}

func (w *runner) runTestFor(testDetails testAttributes) Result {
	var timeTaken time.Duration
	req, err := http.NewRequest(testDetails.methodType, testDetails.link, testDetails.reqBody)
	if err != nil {
		return *getResultOnSetupFailure(err)
	}
	req.Header = testDetails.headers
	startTime := time.Now()
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			timeTaken = time.Since(startTime)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	res, err := w.client.Do(req)
	if err != nil {
		return *getResultOnSetupFailure(err)
	}

	return *getResultPostSetupSuccess(timeTaken, res.StatusCode, testDetails.expectedRespCode)
}

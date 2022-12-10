package runner

import (
	"io"
	"strings"
)

type testAttributes struct {
	link             string
	methodType       string
	headers          map[string][]string
	reqBody          io.Reader
	expectedRespCode int
}

func newTestArributes(methodType, linkAddress, headers, reqBody string, reqHeaders map[string][]string, expectedResp int) *testAttributes { //TODO: refactor
	return &testAttributes{
		link:             linkAddress,
		methodType:       methodType,
		headers:          reqHeaders,
		reqBody:          strings.NewReader(reqBody),
		expectedRespCode: expectedResp,
	}
}

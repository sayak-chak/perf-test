package errors

import "fmt"

func StatusCodeMismatchErr(actualStatus, expectedStatus int) error {
	return fmt.Errorf("status code mismatch, expected status code is %d, but actual status code is %d", expectedStatus, actualStatus)
}

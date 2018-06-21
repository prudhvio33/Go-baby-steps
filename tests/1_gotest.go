/*
	The Go language comes with a lightweight testing framework called testing, and we can use the go test command to execute
	unit and performance tests. Go's testing framework works similarly to testing frameworks in other languages. You can develop
	all sorts of test suites with them, which may include tests for unit testes, benchmarking, stress tests, etc. Let's
	learn about testing in Go, step by step.

*/

package gotest

import (
"errors"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisor can not be 0")
	}
	return a / b, nil
}
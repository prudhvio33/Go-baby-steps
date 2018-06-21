package gotest

import "testing"

/*

Gotest_test.go: This is our unit test file. Keep in mind the following principles for test files:
File names must end in _test.go so that go test can find and execute the appropriate code
You have to import the testing package
All test case functions begin with Test
Test cases follow the source code order
Test functions of the form TestXxx() take a testing.T argument; we can use this type to record errors or to get the testing status
In functions of the form func TestXxx(t * testing.T), the Xxx section can be any alphanumeric combination, but the first letter cannot be a lowercase letter [az]. For example, Testintdiv would be an invalid function name.
By calling one of the Error, Errorf, FailNow, Fatal or FatalIf methods of testing.T on our testing functions, we can fail the test. In addition, we can call the Log method of testing.T to record the information in the error log.
*/

func Test_Division_1(t *testing.T) {
	// try a unit test on function
	if i, e := Division(6, 2); i != 3 || e != nil {
		// If it is not as expected, then the test has failed
		t.Error("division function tests do not pass ")
	} else {
		// record the expected information
		t.Log("first test passed ")
	}
}

func Test_Division_2(t *testing.T) {
	// try a unit test on function
	if _, e := Division(6, 0); e == nil {
		// If it is not as expected, then the error
		t.Error("Division did not work as expected.")
	} else {
		// record some of the information you expect to record
		t.Log("one test passed.", e)
	}
}
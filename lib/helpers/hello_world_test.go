package helpers

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Alfy")
	if result != "Hello Alfy" {
		// errorkan unit test
		panic("Result is not Hello Alfy")
	}
}

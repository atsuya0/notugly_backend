package random

import (
	"testing"
)

func TestRandomStringLength(t *testing.T) {
	request := [3]int{2, 4, 6}

	for _, v := range request {
		randomString := RandomString(v)
		output := len(randomString)
		if output != v {
			t.Errorf("The length is different. request: %d, output: %d", v, output)
		}
	}
}

func TestRandomString(t *testing.T) {
	const (
		length = 10
		max    = 1000
	)
	for i := 0; i < max; i++ {
		if RandomString(length) == RandomString(length) {
			t.Errorf("Not a random character string")
		}
	}
}

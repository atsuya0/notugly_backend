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

	firstString := RandomString(request[0])
	secondString := RandomString(request[0])
	if firstString == secondString {
		t.Errorf("Not a random character string. first: %s, second: %s",
			firstString, secondString)
	}
}

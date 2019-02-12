package random

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	request := [2]int{3, 5}

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

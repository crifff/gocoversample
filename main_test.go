package main

import "testing"

func mainTest(t *testing.T) {
	str := hello()
	if str != "hello" {
		t.Error()
	}

}

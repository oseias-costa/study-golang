package main

import (
	"testing"
	"os"
)

func TestGreeting(t *testing.T){
	result := Greeting("Chris")
	expect := "Hello, Chris"

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}
}


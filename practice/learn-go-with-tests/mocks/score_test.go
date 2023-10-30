package main

import (
	"bytes"
	"testing"
)

func TestScore(t *testing.T) {
	buffer := bytes.Buffer{}
	Score(buffer)

	result := buffer.String()
	expect := `3
	2
	1
	Go!`

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}
}

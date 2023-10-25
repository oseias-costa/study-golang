package main

import "testing"

func TestHello(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, result, expect string) {
		t.Helper()
		if result != expect {
			t.Errorf("result '%s', expect '%s', ", result, expect)
		}
	}

	t.Run("say hello for the people", func(t *testing.T) {
		result := Hello("Oséias")
		expect := "Hello, Oséias"

		verifyCorrectMessage(t, result, expect)
	})

	t.Run("say hello with string empty", func(t *testing.T) {
		result := Hello("")
		expect := "Hello, Oséias"

		verifyCorrectMessage(t, result, expect)
	})

}

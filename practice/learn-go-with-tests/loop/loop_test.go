package loop

import "testing"

func TestLoop(t *testing.T) {
	loop := Loop("a")
	expect := "aaaa"

	if loop != expect {
		t.Errorf("expect '%s', loop '%s' ", expect, loop)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop("a")
	}
}

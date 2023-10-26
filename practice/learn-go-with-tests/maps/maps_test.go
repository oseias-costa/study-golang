package maps

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("word Search", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		result := Search(dictionary, "test")
		expect := "this is a test"

		compareStrings(t, result, expect)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		result := Search(dictionary, "other")
		fmt.Printf("result %s", result)
	})
}
func compareStrings(t *testing.T, result string, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("result %s, expect %s ", result, expect)
	}
}

// utility methods for tests
package testutil

import (
	"testing"
)

// utility for asserting equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

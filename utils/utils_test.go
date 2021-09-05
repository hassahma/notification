// tests for utility functions.
package utils

import (
	"testing"
	"github.com/marvel/constant"
	"github.com/marvel/testutils"
)

// tests the generation of md5 hash.
func TestGetMD5Hash(t *testing.T) {
	testutil.AssertEqual(t, GetMD5Hash("dummy"), "275876e34cf609db118f3d84b799a790")
}

// tests the generation of url
func TestGetCharacterIdUrl(t *testing.T) {
	testutil.AssertEqual(t, GetCharacterIdUrl("1234"), constant.MARVEL_URL + "/1234")
}
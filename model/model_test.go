// tests for model package.
package model

import (
	"testing"
	"github.com/marvel/testutils"
)

// tests the struct representing the http error 400
func TestHTTPError400(t *testing.T) {
	hTTPError400 := HTTPError400{400, "bad"}
	testutil.AssertEqual(t, hTTPError400.Code, 400)
	testutil.AssertEqual(t, hTTPError400.Message, "bad")
}

// tests the struct representing the http error 404
func TestHTTPError404(t *testing.T) {
	hTTPError404 := HTTPError404{404, "not found"}
	testutil.AssertEqual(t, hTTPError404.Code, 404)
	testutil.AssertEqual(t, hTTPError404.Message, "not found")
}

// tests the struct representing the http error 500
func TestHTTPError500(t *testing.T) {
	hTTPError500 := HTTPError400{500, "error"}
	testutil.AssertEqual(t, hTTPError500.Code, 500)
	testutil.AssertEqual(t, hTTPError500.Message, "error")
}


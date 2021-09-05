package model

import (
	"testing"
	"github.com/marvel/testutils"
)

func TestHTTPError400(t *testing.T) {
	hTTPError400 := HTTPError400{400, "bad"}
	testutil.AssertEqual(t, hTTPError400.Code, 400)
	testutil.AssertEqual(t, hTTPError400.Message, "bad")
}

func TestHTTPError404(t *testing.T) {
	hTTPError404 := HTTPError404{404, "bad"}
	testutil.AssertEqual(t, hTTPError404.Code, 404)
	testutil.AssertEqual(t, hTTPError404.Message, "bad")
}

func TestHTTPError500(t *testing.T) {
	hTTPError500 := HTTPError400{500, "bad"}
	testutil.AssertEqual(t, hTTPError500.Code, 500)
	testutil.AssertEqual(t, hTTPError500.Message, "bad")
}


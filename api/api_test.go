package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type ApiSuite struct {
	handler http.Handler
}

var _ = check.Suite(&ApiSuite{})

func (s *ApiSuite) SetUpSuite(c *check.C) {
	s.handler = &CreateDestinyUser{}
}

func (s *ApiSuite) TestOK(c *check.C) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "/user/blah", nil)

	s.handler.ServeHTTP(w, r)

	c.Assert(w.Code, check.Equals, http.StatusOK)
	c.Assert(w.Body.String(), check.Equals, "Hello warriors")
}
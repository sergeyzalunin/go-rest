package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI_HelloHandleFunc(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleFunc().ServeHTTP(rec, req)

	assert.Equal(t, "hello", rec.Body.String())
}

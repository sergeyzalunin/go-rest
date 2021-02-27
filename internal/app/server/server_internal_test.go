package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sergeyzalunin/go-rest/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}

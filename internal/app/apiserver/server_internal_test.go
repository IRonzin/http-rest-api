package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IRonzin/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleHealthCheck(t *testing.T) {
	s := newServer(teststore.New())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Healthy")
	assert.Equal(t, rec.Code, http.StatusOK)
}

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
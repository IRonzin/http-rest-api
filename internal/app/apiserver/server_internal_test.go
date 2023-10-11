package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IRonzin/http-rest-api/internal/app/model"
	"github.com/IRonzin/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleHealthCheck(t *testing.T) {
	s := newServer(teststore.New())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, "Healthy", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestServer_HandleUsersCreate(t *testing.T) {

	testCases := []struct {
		name         string
		payload      any
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "12345678",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "userexample.org",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := bytes.Buffer{}
			json.NewEncoder(&b).Encode(tc.payload)

			req, _ := http.NewRequest(http.MethodPost, "/users", &b)
			s := newServer(teststore.New())
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}

}

func TestServer_HandleSessionsCreate(t *testing.T) {

	user := model.TestUser(t)
	store := teststore.New()
	store.User().Create(user)

	testCases := []struct {
		name         string
		payload      any
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    user.Email,
				"password": user.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    user.Email,
				"password": user.Password + "1",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "email not found",
			payload: map[string]string{
				"email":    user.Email + "1",
				"password": user.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := bytes.Buffer{}
			json.NewEncoder(&b).Encode(tc.payload)

			req, _ := http.NewRequest(http.MethodPost, "/sessions", &b)
			s := newServer(store)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

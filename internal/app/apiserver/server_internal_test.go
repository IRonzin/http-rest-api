package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IRonzin/http-rest-api/internal/app/model"
	"github.com/IRonzin/http-rest-api/internal/app/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)

	testCases := []struct {
		name         string
		cookieValue  map[any]any
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[any]any{
				"user_id": u.Id,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	secretKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleHealthCheck(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
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
			s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
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
			s := newServer(store, sessions.NewCookieStore([]byte("secret")))
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

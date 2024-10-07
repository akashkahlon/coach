package api

import (
	"bytes"
	"coach/testutils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	testutils.SetupTestEnv(m)
	testutils.RunTests(m)
}

func TestLoginHandler_Success(t *testing.T) {
	email := "a@b.com"
	password := "password"
	testutils.CreateUser(email, password)

	reqBody := map[string]string{
		"email":    email,
		"password": password,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Could not marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := LoginHandler(testutils.TestDB)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned notOK status code: got %v want %v", status, http.StatusOK)
	}

	cookie := rr.Result().Cookies()[0]

	if cookie.Name != "token" {
		t.Errorf("Handler did not set the token cookie")
	}

	if cookie.HttpOnly != true {
		t.Errorf("Cookie is not HttpOnly")
	}
	
	if cookie.Value == "" {
		t.Errorf("Token cookie is empty")
	}

	if cookie.Expires.IsZero() {
		t.Errorf("Token cookie does not have an expiration time")
	}

	testutils.CleanUpDatabase()
}

func TestLoginHandler_InvalidRequest(t *testing.T) {
	invalidJSON := `{"emailID": "a@b.com", "password": }` // Malformed JSON (missing value for "password")

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(invalidJSON)))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := LoginHandler(testutils.TestDB)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestLoginHandler_InvalidRequestFieldMissing(t *testing.T) {
	reqBody := map[string]string{
		"emailID":    "a@b.com",
		"password": "password",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Could not marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := LoginHandler(testutils.TestDB)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned notOK status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	email := "a@b.com"
	password := "password"
	testutils.CreateUser(email, password)

	reqBody := map[string]string{
		"email":    email,
		"password": "wrongpassword",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Could not marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := LoginHandler(testutils.TestDB)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned notOK status code: got %v want %v", status, http.StatusUnauthorized)
	}
	
	testutils.CleanUpDatabase()
}
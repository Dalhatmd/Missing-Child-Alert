package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dalhatmd/Missing-Child-Alert/db"
	"github.com/dalhatmd/Missing-Child-Alert/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// mockUser is a helper to create a user.User pointer for testing.
func mockUser(id, username, email, password, location string) *user.User {
	return &user.User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
		Location: location,
	}
}

// mockDB is a helper to mock db.DB.Create
type mockDB struct {
	err error
}

func (m *mockDB) Create(value interface{}) *gorm.DB {
	return &gorm.DB{Error: m.err}
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/users", CreateUser)
	return r
}

func TestCreateUser_Success(t *testing.T) {
	// Mock dependencies
	db.DB = &mockDB{err: nil}
	user.NewUser = func(id, username, email, password, location string) (*user.User, error) {
		return mockUser(id, username, email, password, location), nil
	}

	router := setupRouter()
	email := "test@example.com"
	password := "password123"
	username := "testuser"
	location := "Test City"
	body := map[string]interface{}{
		"email":    email,
		"password": password,
		"username": username,
		"location": location,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp user.User
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, email, resp.Email)
	assert.Equal(t, username, resp.Username)
	assert.Equal(t, location, resp.Location)
}

func TestCreateUser_MissingEmail(t *testing.T) {
	router := setupRouter()
	body := map[string]interface{}{
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "email is required")
}

func TestCreateUser_MissingPassword(t *testing.T) {
	router := setupRouter()
	body := map[string]interface{}{
		"email": "test@example.com",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password is required")
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestCreateUser_NewUserError(t *testing.T) {
	db.DB = &mockDB{err: nil}
	user.NewUser = func(id, username, email, password, location string) (*user.User, error) {
		return nil, errors.New("new user error")
	}

	router := setupRouter()
	body := map[string]interface{}{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "new user error")
}

func TestCreateUser_DBError(t *testing.T) {
	db.DB = &mockDB{err: errors.New("db error")}
	user.NewUser = func(id, username, email, password, location string) (*user.User, error) {
		return mockUser(id, username, email, password, location), nil
	}

	router := setupRouter()
	body := map[string]interface{}{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "db error")
}

// We recommend installing an extension to run go tests.

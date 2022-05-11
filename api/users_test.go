package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/users", nil)
	if err != nil {
		log.Fatal(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetUsers(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		log.Fatal(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

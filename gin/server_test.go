package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testRouter = setupRouter()

func TestRoute_Data(t *testing.T) {	
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/data", nil)
	if err != nil {
		t.Errorf("Error: http.NewRequest")
	}

	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Some binary data here.", w.Body.String())
}

func TestRoute_JSON(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/json", nil)
	if err != nil {
		t.Errorf("Error: http.NewRequest")
	}

	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"hello\":\"json\"}", w.Body.String())
}

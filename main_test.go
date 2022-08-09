package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestWithNoQueryParamReturnAllData(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRoot)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestRequestWithInvalidParameter(t *testing.T) {
	req, err := http.NewRequest("GET", "/?id=xxx", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRoot)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusBadRequest)
}

func TestRequestWithSingleID(t *testing.T) {
	req, err := http.NewRequest("GET", "/?id=1", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRoot)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestRequestWithSingleIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/?id=7", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRoot)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusNotFound)
}

func TestRequestWithMultipleID(t *testing.T) {
	req, err := http.NewRequest("GET", "/?id=1,2,3,4,5", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRoot)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusOK)
}

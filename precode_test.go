package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getResponse(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(resp, req)

	return resp
}

func TestMainHandlerOk(t *testing.T) {
	response := getResponse("/cafe?city=moscow&count=3")
	targetCode := http.StatusOK
	targetCount := 3

	assert.Equal(t, targetCode, response.Code)
	require.NotEmpty(t, response.Body)
	body := strings.Split(response.Body.String(), ",")
	assert.Len(t, body, targetCount)
}

func TestMainHandlerWrongCity(t *testing.T) {
	response := getResponse("/cafe?city=kolomna&count=2")
	targetBody := "wrong city value"
	targetCode := http.StatusBadRequest

	assert.Equal(t, targetCode, response.Code)
	require.NotEmpty(t, response.Body)
	body := response.Body.String()
	assert.Equal(t, targetBody, body)
}

func TestMainHandlerIncorrectCount(t *testing.T) {
	response := getResponse("/cafe?city=moscow&count=5")
	targetCode := http.StatusOK
	targetCount := 4

	assert.Equal(t, targetCode, response.Code)
	require.NotEmpty(t, response.Body)
	body := strings.Split(response.Body.String(), ",")
	assert.Len(t, body, targetCount)
}

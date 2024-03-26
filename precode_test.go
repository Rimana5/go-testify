package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=7&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=7&city=london", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	require.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	totalCount := 4

	// Создаем запрос с count больше, чем количество кафе
	req, err := http.NewRequest("GET", "/?city=moscow&count=5", nil)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
}

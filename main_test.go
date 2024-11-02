package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	// Обработаем запрос
	handler.ServeHTTP(responseRecorder, req)
	// Проверка статуса ответа
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	// Проверка тела ответа
	expectedCafes := strings.Join(cafeList["moscow"], ",")
	assert.Equal(t, expectedCafes, responseRecorder.Body.String())
	// Убедитесь, что тело не пустое
	assert.NotEmpty(t, responseRecorder.Body.String())
	// Проверка длины ответа
	assert.Len(t, responseRecorder.Body.String(), len(expectedCafes))
}
func TestMainHandlerWithIncorrectCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=notarealcity&count=2", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	// Обработаем запрос
	handler.ServeHTTP(responseRecorder, req)
	// Проверка статуса ответа
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	// Проверка тела ответа
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}
func TestMainHandlerWithMissingCount(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	// Обработаем запрос
	handler.ServeHTTP(responseRecorder, req)
	// Проверка статуса ответа
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	// Проверка тела ответа
	assert.Equal(t, "count missing", responseRecorder.Body.String())
}

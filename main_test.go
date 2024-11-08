package main

import (
	"net/http"
	"net/http/httptest"
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
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	// Ожидаемое значение ответа
	expectedCafes := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedCafes, responseRecorder.Body.String())
}

func TestMainHandlerWithIncorrectCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=notarealcity&count=2", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	// Обработаем запрос
	handler.ServeHTTP(responseRecorder, req)
	// Проверка статуса ответа
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
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
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	// Проверка тела ответа
	assert.Equal(t, "count missing", responseRecorder.Body.String())
}

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	// Обработаем запрос
	handler.ServeHTTP(responseRecorder, req)
	// Проверка статуса ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	// Проверка, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body.String())
}

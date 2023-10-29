package main

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "server is running", w.Body.String())
}

func TestIsXml(t *testing.T) {
	type testCase struct {
		xml         string
		status      int
		expectedErr error
	}

	testCases := []testCase{
		{
			xml:    "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<note>\n  <to>Tove</to>\n  <from>Jani</from>\n  <heading>Reminder</heading>\n  <body>Don't forget to buy milk on your way home!</body>\n</note>\n",
			status: http.StatusOK,
		},
		{
			xml:    "<root>\n    <element1>Value 1</element1>\n    <element2>Value 2\n</root>",
			status: http.StatusBadRequest,
		},
	}
	router := setupRouter()

	for _, s := range testCases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/xml/isValid", strings.NewReader(s.xml))
		router.ServeHTTP(w, req)

		assert.Equal(t, s.status, w.Code)
	}
}

func TestXmlToYaml(t *testing.T) {
	router := setupRouter()

	xml := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<note>\n  <to>Tove</to>\n  <from>Jani</from>\n  <heading>Reminder</heading>\n  <body>Don't forget to buy milk on your way home!</body>\n</note>\n"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/xml/yaml", strings.NewReader(xml))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/x-yaml", w.Header().Get("content-type"))
}

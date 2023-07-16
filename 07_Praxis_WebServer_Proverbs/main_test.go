package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleQuote(t *testing.T) {
	// 1. HTTP Recorder erstellen
	recorder := httptest.NewRecorder()

	// 2. Request erstellen
	req, _ := http.NewRequest("GET", "/api", nil)

	// 3. Handler Funktion aufrufen
	HandleProverbJson(recorder, req)

	// 4. Ergebnis prüfen
	// a) Return Code prüfen
	if recorder.Code != http.StatusOK {
		t.Errorf("Wrong status: got %v expected %v", recorder.Code, http.StatusOK)
	}

	// b) Body prüfen
	body := recorder.Body.String()
	if !strings.Contains(body, "link") {
		t.Errorf("Could not find link in json:\n%v", body)
	}

}

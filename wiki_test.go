package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestViewHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/view/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeHandler(viewHandler))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

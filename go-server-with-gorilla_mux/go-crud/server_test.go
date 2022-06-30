package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Parallel()
	r, err := http.NewRequest("GET", "/person", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	getPersons(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler error got:%v want:%v", status, http.StatusOK)
	}
}

package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(userAgent string) (*httptest.ResponseRecorder, *http.Request) {
	r, _ := http.NewRequest("GET", "/", nil)
	r.Header["User-Agent"] = []string{userAgent}
	w := httptest.NewRecorder()
	return w, r
}

func TestTrickyHandler(t *testing.T) {
	th := &trickyHandler{uaTargets: []string{"target useragent"}, imageData: []byte("img"), attackData: []byte("attack")}
	w, r := testRequest("target useragent")
	th.ServeHTTP(w, r)
	if !bytes.Equal(w.Body.Bytes(), []byte("img")) {
		t.Fatal("expected uaTarget to be given an image")
	}
	w, r = testRequest("other useragent")
	th.ServeHTTP(w, r)
	if !bytes.Equal(w.Body.Bytes(), []byte("attack")) {
		t.Fatal("expected non uaTarget to be given attack")
	}
}

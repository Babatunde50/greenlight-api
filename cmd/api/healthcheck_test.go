package main

import (
	"net/http"
	"testing"
)

func Test_healthcheckHandler(t *testing.T) {

	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.send(t, "/v1/healthcheck", map[string]string{})

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

}

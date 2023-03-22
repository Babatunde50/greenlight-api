package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Babatunde50/green-light/internal/data"
	"github.com/Babatunde50/green-light/internal/jsonlog"
	"github.com/Babatunde50/green-light/internal/mailer"
)

func newTestApplication(t *testing.T) *application {

	return &application{
		config: config{
			port: 8080,
			env:  "development",
			db: struct {
				dsn          string
				maxOpenConns int
				maxIdleConns int
				maxIdleTime  string
			}{
				dsn:          "postgres://user:password@localhost:5432/mydb",
				maxOpenConns: 10,
				maxIdleConns: 5,
				maxIdleTime:  "30m",
			},
			limiter: struct {
				rps     float64
				burst   int
				enabled bool
			}{
				rps:     2,
				burst:   4,
				enabled: false,
			},
			smtp: struct {
				host     string
				port     int
				username string
				password string
				sender   string
			}{
				host:     "",
				port:     2525,
				username: "",
				password: "",
				sender:   "Greenlight",
			},
			cors: struct{ trustedOrigins []string }{
				trustedOrigins: []string{"http://localhost:3000", "https://example.com"},
			},
		},
		logger: jsonlog.New(os.Stdout, jsonlog.LevelInfo),
		models: data.NewMockModels(),
		mailer: mailer.New("", 2525, "", "", "Greenlight"),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

func (ts *testServer) send(t *testing.T, urlPath string, headers map[string]string) (int, http.Header, []byte) {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL+urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	return resp.StatusCode, resp.Header, body
}

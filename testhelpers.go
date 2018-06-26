package main

import (
	"bytes"
	"fmt"
	"github.com/asdine/storm"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"time"
)

func createTestDatabase() *storm.DB {
	Dir := path.Join("/tmp/gin-rest-api", fmt.Sprintf("%d", time.Now().UnixNano()))
	err := os.MkdirAll(Dir, os.FileMode(0700))
	if err != nil {
		panic(err)
	}

	Dir = path.Join(Dir, "test.db")
	db, _ := storm.Open(Dir)

	return db
}

func createTestApplication() *httptest.Server {
	db := createTestDatabase()
	app := &Application{Database: db}
	api := newAPIServer(app)
	return api
}

func newAPIServer(app *Application) *httptest.Server {
	apiEngine := setupRouter(app)
	return httptest.NewServer(apiEngine)
}

// ParseResponseBody parses response body into a string
func ParseResponseBody(resp *http.Response) string {
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(resp.Body)
	mustNotErr(resp.Body.Close())
	return buf.String()
}

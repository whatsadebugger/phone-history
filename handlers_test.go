package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const appjson = "application/json"

func TestCreateAddress(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	address := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
	new := bytes.NewBufferString(address)
	res, err := BasicPost(url+"/address", appjson, new)

	body := ParseResponseBody(res)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, address, body)
}

func TestFindAddress(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	address := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
	new := bytes.NewBufferString(address)
	res, err := BasicPost(url+"/address", appjson, new)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	res, err = BasicGet(url + "/address/1")
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body := ParseResponseBody(res)

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, address, body)
}

func TestFindAllAddress(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	address := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
	new := bytes.NewBufferString(address)
	res, err := BasicPost(url+"/address", appjson, new)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body := ParseResponseBody(res)

	res, err = BasicGet(url + "/address/")
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, address, body)
}

func TestUpdateAddress(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	address := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
	new := bytes.NewBufferString(address)
	res, err := BasicPost(url+"/address", appjson, new)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	update := `{"id":1,"firstname":"newname","lastname":"newlastname","email":"iii@winterfell.com", "phone": "133-333-1313"}`
	new = bytes.NewBufferString(update)
	res, err = BasicPut(url+"/address/", appjson, new)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body := ParseResponseBody(res)

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, update, body)
}

func TestDeleteAddress(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	address := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
	new := bytes.NewBufferString(address)
	res, err := BasicPost(url+"/address", appjson, new)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	res, err = BasicDelete(url+"/address/1", appjson, nil)
	assert.NoError(t, err)

	body := ParseResponseBody(res)

	expected := "Success"
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expected, body)
}

func TestImportCSV(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	data, err := os.Open("test.csv")
	assert.NoError(t, err)

	res, err := PostTextCSV(url+"/address/upload", data)
	assert.NoError(t, err)

	body := ParseResponseBody(res)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "imported 2 records", body)
}

func TestExportCSV(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	address := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
	new := bytes.NewBufferString(address)
	res, err := BasicPost(url+"/address", appjson, new)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	res, err = BasicGet(url + "/addressbook")

	expected := "john,snow,snow@winterfell.com,133-333-1313\n"
	body := ParseResponseBody(res)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expected, body)
}

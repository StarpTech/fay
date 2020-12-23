package controller

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/assert"
)

func newFileUploadRequest(uri string, params map[string]string, files map[string]string) (*http.Request, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	defer writer.Close()

	for key, val := range files {
		file, err := os.Open(val)
		if err != nil {
			return nil, err
		}
		part, err := writer.CreateFormFile(key, filepath.Base(val))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
		file.Close()
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	request := httptest.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request, nil
}

func TestDifferentFilename(t *testing.T) {
	pw, err := playwright.Run()
	assert.NoError(t, err)

	browser, err := pw.Chromium.Launch()
	assert.NoError(t, err)
	defer pw.Stop()
	defer browser.Close()

	e := echo.New()

	params := map[string]string{
		"url":      "https://google.com",
		"filename": "foo.pdf",
	}
	files := map[string]string{}
	req, err := newFileUploadRequest("/", params, files)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := Http{
		Browser: browser,
	}

	if assert.NoError(t, h.ConvertHTML(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "attachment; filename=\"foo.pdf\"", rec.HeaderMap.Get("content-disposition"))
		assert.Equal(t, "application/pdf", rec.HeaderMap.Get("content-type"))
		assert.Equal(t, "bytes", rec.HeaderMap.Get("accept-ranges"))
		contentLength, _ := strconv.Atoi(rec.HeaderMap.Get("content-length"))
		assert.Greater(t, contentLength, 200)
	}
}

func TestRenderUrlToPDF(t *testing.T) {
	pw, err := playwright.Run()
	assert.NoError(t, err)

	browser, err := pw.Chromium.Launch()
	assert.NoError(t, err)
	defer pw.Stop()
	defer browser.Close()

	e := echo.New()

	params := map[string]string{
		"url": "https://google.com",
	}
	files := map[string]string{}
	req, err := newFileUploadRequest("/", params, files)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := Http{
		Browser: browser,
	}

	if assert.NoError(t, h.ConvertHTML(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "attachment; filename=\"result.pdf\"", rec.HeaderMap.Get("content-disposition"))
		assert.Equal(t, "application/pdf", rec.HeaderMap.Get("content-type"))
		assert.Equal(t, "bytes", rec.HeaderMap.Get("accept-ranges"))
		contentLength, _ := strconv.Atoi(rec.HeaderMap.Get("content-length"))
		assert.Greater(t, contentLength, 200)
	}
}

func TestRenderHTMLToPDF(t *testing.T) {
	pw, err := playwright.Run()
	assert.NoError(t, err)

	browser, err := pw.Chromium.Launch()
	assert.NoError(t, err)
	defer pw.Stop()
	defer browser.Close()

	e := echo.New()
	wd, _ := os.Getwd()

	params := map[string]string{}
	files := map[string]string{
		"html": path.Join(wd, "../../example/page.html"),
	}
	req, err := newFileUploadRequest("/", params, files)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := Http{
		Browser: browser,
	}

	if assert.NoError(t, h.ConvertHTML(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "attachment; filename=\"result.pdf\"", rec.HeaderMap.Get("content-disposition"))
		assert.Equal(t, "application/pdf", rec.HeaderMap.Get("content-type"))
		assert.Equal(t, "bytes", rec.HeaderMap.Get("accept-ranges"))
		contentLength, _ := strconv.Atoi(rec.HeaderMap.Get("content-length"))
		assert.Greater(t, contentLength, 200)
	}
}

func TestPingOK(t *testing.T) {
	pw, err := playwright.Run()
	assert.NoError(t, err)

	browser, err := pw.Chromium.Launch()
	assert.NoError(t, err)
	defer pw.Stop()
	defer browser.Close()

	e := echo.New()

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := Http{
		Browser: browser,
	}

	if assert.NoError(t, h.Ping(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestPingNotOK(t *testing.T) {
	pw, err := playwright.Run()
	assert.NoError(t, err)

	browser, err := pw.Chromium.Launch()
	assert.NoError(t, err)
	defer pw.Stop()

	onCloseWasCalled := make(chan bool, 1)
	onClose := func() {
		onCloseWasCalled <- true
	}

	e := echo.New()

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := Http{
		Browser: browser,
	}

	browser.On("close", onClose)
	assert.True(t, browser.IsConnected())
	assert.NoError(t, browser.Close())
	<-onCloseWasCalled

	if assert.NoError(t, h.Ping(c)) {
		assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
	}
}

// HTTP Client to use with Team Foundation Server
package tfs

import (
	"encoding/base64"
	"io"
	"net/http"
)

// Set headers on each request: https://developer20.com/add-header-to-every-request-in-go/
type TfsHttpClient struct {
	Token string
	c     http.Client
}

// Function send HTTP GET request and expect Json back
func (c *TfsHttpClient) GetJson(url string) (resp *http.Response, err error) {

	return c.Get(url, "application/json", "application/json; charset=utf-8")
}

// Function send HTTP GET request and expect binary data back
func (c *TfsHttpClient) DownloadFileFrom(url string) (resp *http.Response, err error) {

	return c.Get(url, "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7", "text/plain; api-version=3.2")
}

func (c *TfsHttpClient) Get(url, accept, content_type string) (resp *http.Response, err error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if accept != "" {
		req.Header.Set("Accept", accept)
	}

	if content_type != "" {
		req.Header.Set("Content-Type", content_type)
	}

	return c.Do(req)
}

// Function send HTTP POST request
func (c *TfsHttpClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err

	}

	req.Header.Set("Content-Type", contentType)

	return c.Do(req)
}

// Function actually execute HTTP request
func (c *TfsHttpClient) Do(req *http.Request) (*http.Response, error) {

	auth := "" + ":" + c.Token
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

	return c.c.Do(req)
}

// HTTP Client to use with Team Foundation Server
package tfs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Set headers on each request: https://developer20.com/add-header-to-every-request-in-go/
type TfsHttpClient struct {
	Token string
	c     http.Client
}

func (c *TfsHttpClient) GetWorkItem(url string) (*TfsWorkItem, error) {

	resp, err := c.getJson(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tfsWorkItem := &TfsWorkItem{}

	err = json.Unmarshal(body, tfsWorkItem)
	if err != nil {
		fmt.Printf("body is: %s", string(body))
		return nil, err
	}

	return tfsWorkItem, nil
}

func (c *TfsHttpClient) GetHistoryLinks(url string) (*WorkItemHistory, error) {

	resp, err := c.getJson(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	workItemHistory := &WorkItemHistory{}

	err = json.Unmarshal(body, workItemHistory)
	if err != nil {
		return nil, err
	}

	return workItemHistory, nil
}

// Function send HTTP GET request and expect Json back
func (c *TfsHttpClient) getJson(url string) (resp *http.Response, err error) {

	return c.get(url, "application/json", "application/json; charset=utf-8")
}

// Function send HTTP GET request and expect binary data back
func (c *TfsHttpClient) downloadFileFrom(url string) (resp *http.Response, err error) {

	return c.get(url, "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7", "text/plain; api-version=3.2")
}

func (c *TfsHttpClient) get(url, accept, content_type string) (resp *http.Response, err error) {

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

	return c.do(req)
}

// Function send HTTP POST request
func (c *TfsHttpClient) post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err

	}

	req.Header.Set("Content-Type", contentType)

	return c.do(req)
}

// Function actually execute HTTP request
func (c *TfsHttpClient) do(req *http.Request) (*http.Response, error) {

	auth := "" + ":" + c.Token
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

	return c.c.Do(req)
}

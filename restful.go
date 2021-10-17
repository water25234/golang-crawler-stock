package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	domainAPI = flag.String("API", "https://mops.twse.com.tw/", "api server domain")
)

// HTTPClient means
type HTTPClient interface {
	Get(path string) (resp interface{}, err error)

	Post(path string, requestBody interface{}) (resp interface{}, err error)

	Put(path string, requestBody interface{}) (resp interface{}, err error)

	Patch(path string, requestBody interface{}) (resp interface{}, err error)

	Delete(path string, requestBody interface{}) (resp interface{}, err error)
}

// RequestHTTP means
type RequestHTTP struct {
	Request       *http.Request
	HTTPClient    *http.Client
	RequestHeader *[]RequestHeader
}

// New means
func New() HTTPClient {
	return &RequestHTTP{
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// RequestHeader means
type RequestHeader struct {
	Key   string
	Value string
}

// NewHeader means
// func (req *RequestHTTP) NewHeader(key string, value string) {
func (req *RequestHTTP) NewHeader(headers *[]RequestHeader) {
	if headers == nil {
		return
	}

	for _, header := range *headers {
		req.Request.Header.Set(header.Key, header.Value)
	}
}

// NewRequest means
func (req *RequestHTTP) NewRequest(method string, path string, requestbody interface{}) (err error) {
	url := fmt.Sprintf("%s%s", *domainAPI, path)

	body, err := requestBody(requestbody)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	req.Request, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

// ExecuteAPI means
func (req *RequestHTTP) ExecuteAPI() (response interface{}, err error) {
	resp, err := req.HTTPClient.Do(req.Request)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer resp.Body.Close()
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	err = json.Unmarshal(respBodyBytes, &response)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return response, nil
}

// requestBody means
func requestBody(data interface{}) (body []byte, err error) {
	switch data := data.(type) {
	case nil:
		return nil, nil
	case string:
		return []byte(data), nil
	default:
		body, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
}

// Get means
func (req *RequestHTTP) Get(path string) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	headers := &[]RequestHeader{}
	req.NewHeader(headers)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}

// Post means
func (req *RequestHTTP) Post(path string, requestBody interface{}) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodPost, path, requestBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// headers := &[]RequestHeader{
	// 	{
	// 		Key:   "Content-Type",
	// 		Value: "application/json; charset=utf-8",
	// 	},
	// }
	req.NewHeader(req.RequestHeader)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}

// Put means
func (req *RequestHTTP) Put(path string, requestBody interface{}) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodPut, path, requestBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	headers := &[]RequestHeader{
		{
			Key:   "Content-Type",
			Value: "application/json; charset=utf-8",
		},
	}
	req.NewHeader(headers)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}

// Patch means
func (req *RequestHTTP) Patch(path string, requestBody interface{}) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodPatch, path, requestBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	headers := &[]RequestHeader{
		{
			Key:   "Content-Type",
			Value: "application/json; charset=utf-8",
		},
	}
	req.NewHeader(headers)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}

// Delete means
func (req *RequestHTTP) Delete(path string, requestBody interface{}) (resp interface{}, err error) {
	err = req.NewRequest(http.MethodDelete, path, requestBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	headers := &[]RequestHeader{
		{
			Key:   "Content-Type",
			Value: "application/json; charset=utf-8",
		},
	}
	req.NewHeader(headers)

	resp, err = req.ExecuteAPI()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return resp, nil
}

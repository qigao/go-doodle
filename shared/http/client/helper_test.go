package client

import (
	"bytes"
	"context"
	http_request "http/request"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestApiURL(t *testing.T) {
	t.Run("when baseURL is invalid", func(t *testing.T) {
		baseURL, err := apiURL("://", "hello", nil)

		assert.Nil(t, baseURL)
		assert.EqualError(t, err, `parse "://": missing protocol scheme`)
	})

	t.Run("when baseURL is valid", func(t *testing.T) {
		fullURL, err := apiURL("http://api.test.xx", "hello", nil)
		assert.Nil(t, err)
		assert.Equal(t, "http://api.test.xx/hello", fullURL.String())

		fullURL, err = apiURL("http://api.test.xx/", "hello", nil)
		assert.Nil(t, err)
		assert.Equal(t, "http://api.test.xx/hello", fullURL.String())

		fullURL, err = apiURL("http://api.test.xx", "/hello", nil)
		assert.Nil(t, err)
		assert.Equal(t, "http://api.test.xx/hello", fullURL.String())

		fullURL, err = apiURL("http://api.test.xx/", "/hello", nil)
		assert.Nil(t, err)
		assert.Equal(t, "http://api.test.xx/hello", fullURL.String())

		fullURL, err = apiURL("http://api.test.xx", "/hello/world", nil)
		assert.Nil(t, err)
		assert.Equal(t, "http://api.test.xx/hello/world", fullURL.String())

		fullURL, err = apiURL("http://test.xx/api", "hello/world", nil)
		assert.Nil(t, err)
		assert.Equal(t, "http://test.xx/api/hello/world", fullURL.String())

		fullURL, err = apiURL("http://test.xx/api", "/hello/world", nil)
		assert.Nil(t, err)
		assert.Equal(t, "http://test.xx/api/hello/world", fullURL.String())

		fullURL, err = apiURL("http://test.xx/api/", "/hello/world", nil)
		assert.Nil(t, err)
		assert.Equal(t, "http://test.xx/api/hello/world", fullURL.String())

		queryParams := url.Values{}
		queryParams.Add("foo", "bar")
		fullURL, err = apiURL("http://test.xx/api/", "/hello/world", queryParams)
		assert.Nil(t, err)
		assert.Equal(t, "http://test.xx/api/hello/world?foo=bar", fullURL.String())
	})
}

func TestSetHeaders(t *testing.T) {
	headers := []string{
		"Authorization",
		"X-Customer-Id",
		"X-Client-Bank-Id",
	}

	t.Run("when context includes `request_header.ForwardedHeadersKey{}`", func(t *testing.T) {
		ctx := context.WithValue(
			context.Background(),
			http_request.ForwardedHeadersKey{},
			http_request.ForwardedHeaders{
				"Authorization":    "token",
				"X-Customer-Id":    "1234",
				"X-Client-Bank-Id": "Bank01",
				"Unknown-Header":   "data",
			},
		)

		req, _ := http.NewRequest("GET", "/test", nil)
		setHeaders(ctx, req, headers)

		assert.Equal(t, "token", req.Header.Get("Authorization"))
		assert.Equal(t, "1234", req.Header.Get("X-Customer-Id"))
		assert.Equal(t, "Bank01", req.Header.Get("X-Client-Bank-Id"))
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

		assert.Equal(t, "", req.Header.Get("Unknown-Header"))
	})

	t.Run("when context does not include `request_header.ForwardedHeadersKey{}`", func(t *testing.T) {
		ctx := context.Background()

		req, _ := http.NewRequest("GET", "/test", nil)
		setHeaders(ctx, req, headers)

		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

		assert.Equal(t, "", req.Header.Get("Authorization"))
		assert.Equal(t, "", req.Header.Get("X-Customer-Id"))
		assert.Equal(t, "", req.Header.Get("X-Client-Bank-Id"))
	})
}

func TestSetInvalidHeaders(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	ctx := context.WithValue(
		context.Background(),
		http_request.ForwardedHeadersKey{},
		http_request.ForwardedHeaders{
			"Authorization":    "token",
			"X-Customer-Id":    "1234",
			"X-Client-Bank-Id": "Bank01",
			"Unknown-Header":   "data",
		},
	)

	headers := []string{
		"Header",
	}

	setHeaders(ctx, req, headers)

	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

	assert.Equal(t, "", req.Header.Get("Unknown-Header"))
}

func TestResponseBody(t *testing.T) {
	t.Run("when statusCode is 200", func(t *testing.T) {
		t.Run("when response body is erroneous", func(t *testing.T) {
			response := &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body:       &erroneousReadCloser{},
			}

			body, err := responseBody(response)

			assert.Nil(t, body)
			assert.EqualError(t, err, "error on ReadCloser.Read")
		})

		t.Run("when response body is valid", func(t *testing.T) {
			response := &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`{"hello": "world"}`)),
			}

			body, err := responseBody(response)

			assert.Nil(t, err)
			assert.Equal(t, `{"hello": "world"}`, string(body))
		})
	})

	t.Run("when statusCode is 201", func(t *testing.T) {
		t.Run("when response body is erroneous", func(t *testing.T) {
			response := &http.Response{
				Status:     "201 Created",
				StatusCode: 201,
				Body:       &erroneousReadCloser{},
			}

			body, err := responseBody(response)

			assert.Nil(t, body)
			assert.EqualError(t, err, "error on ReadCloser.Read")
		})

		t.Run("when response body is valid", func(t *testing.T) {
			response := &http.Response{
				Status:     "201 Created",
				StatusCode: 201,
				Body:       ioutil.NopCloser(strings.NewReader(`{"hello": "world"}`)),
			}

			body, err := responseBody(response)

			assert.Nil(t, err)
			assert.Equal(t, `{"hello": "world"}`, string(body))
		})
	})

	t.Run("when statusCode is 204", func(t *testing.T) {
		t.Run("when response body is erroneous", func(t *testing.T) {
			response := &http.Response{
				Status:     "204 Not Modified",
				StatusCode: 204,
				Body:       &erroneousReadCloser{},
			}

			body, err := responseBody(response)

			assert.Nil(t, body)
			assert.EqualError(t, err, "error on ReadCloser.Read")
		})

		t.Run("when response body is valid", func(t *testing.T) {
			response := &http.Response{
				Status:     "204 Not Modified",
				StatusCode: 204,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
			}

			body, err := responseBody(response)

			assert.Nil(t, err)
			assert.Equal(t, "", string(body))
		})
	})

	t.Run("when statusCode is 400", func(t *testing.T) {
		responseBodyString := `{"errorCode": "400", "message": "Validation error", "fieldErrors": [{"objectName": "user", "field": "first_name", "message": "is missing"}]}`
		response := &http.Response{
			Status:     "400 Bad Request",
			StatusCode: 400,
			Body:       ioutil.NopCloser(strings.NewReader(responseBodyString)),
		}

		body, err := responseBody(response)

		assert.Nil(t, body)
		assert.IsType(t, &Error{}, err)
		assert.EqualError(t, err, "Validation error")
	})

	t.Run("when statusCode is 401", func(t *testing.T) {
		response := &http.Response{
			Status:     "401 Unauthorized",
			StatusCode: 401,
			Body:       ioutil.NopCloser(strings.NewReader(``)),
		}

		body, err := responseBody(response)

		assert.Nil(t, body)
		assert.IsType(t, &Error{}, err)
		assert.EqualError(t, err, "Unauthorized")
	})

	t.Run("when statusCode is 404", func(t *testing.T) {
		responseBodyString := `{"errorCode": "404", "message": "Something not found", "status": "NOT_FOUND", "fieldErrors": []}`

		response := &http.Response{
			Status:     "404 Not Found",
			StatusCode: 404,
			Body:       ioutil.NopCloser(strings.NewReader(responseBodyString)),
		}

		body, err := responseBody(response)

		assert.Nil(t, body)
		assert.IsType(t, &NotFoundError{}, err)
		assert.EqualError(t, err, "Something not found")
	})

	t.Run("when statusCode is 422", func(t *testing.T) {
		response := &http.Response{
			Status:     "422 Unprocessable Entity",
			StatusCode: 422,
			Body:       ioutil.NopCloser(strings.NewReader(`{"errorCode": "422", "message": "Unprocessable Entity"}`)),
		}

		body, err := responseBody(response)

		assert.Nil(t, body)
		assert.IsType(t, &Error{}, err)
		assert.EqualError(t, err, "Unprocessable Entity")
	})

	t.Run("when statusCode is 500", func(t *testing.T) {
		response := &http.Response{
			Status:     "500 Internal Server Error",
			StatusCode: 500,
			Body:       ioutil.NopCloser(strings.NewReader(`Internal Server Error`)),
		}

		body, err := responseBody(response)

		assert.Nil(t, body)
		assert.EqualError(t, err, "client error: `500 Internal Server Error`")
	})
}

type erroneousReadCloser struct{}

func (d *erroneousReadCloser) Read(p []byte) (n int, err error) {
	return 0, errors.New("error on ReadCloser.Read")
}

func (d *erroneousReadCloser) Close() error {
	return nil
}

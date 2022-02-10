package http_request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var headersToForward = []string{
	"Authorization",
	"X-Customer-Id",
	"X-Client-Bank-Id",
	"X-Mock-Services",
}

func TestCreateHandler(t *testing.T) {
	req := &http.Request{
		Header: make(http.Header),
	}

	req.Header.Set("Authorization", "AUTH_HEADER")
	req.Header.Set("X-Customer-Id", "customer_id")
	req.Header.Set("X-Client-Bank-Id", "Bank01")
	req.Header.Set("X-Mock-Services", "true")
	req.Header.Set("Another-Header", "XYZ")

	// NOTE: This is a workaround for reading context values after making a request
	writeContextToBody := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.Context().Value(ForwardedHeadersKey{}).(ForwardedHeaders)
		body, _ := json.Marshal(values)

		w.Write(body)
	})

	handler := CreateHandler(writeContextToBody, headersToForward)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	forwardedHeaders := make(map[string]string)

	json.Unmarshal(recorder.Body.Bytes(), &forwardedHeaders)

	assert.Equal(t, "AUTH_HEADER", forwardedHeaders["Authorization"])
	assert.Equal(t, "Bank01", forwardedHeaders["X-Client-Bank-Id"])
	assert.Equal(t, "customer_id", forwardedHeaders["X-Customer-Id"])
	assert.Equal(t, "true", forwardedHeaders["X-Mock-Services"])
	assert.Equal(t, "", forwardedHeaders["Another-Header"])
}

func TestAddHeadersToContext(t *testing.T) {
	w := &httptest.ResponseRecorder{}
	req := &http.Request{
		Header: make(http.Header),
	}

	req.Header.Set("Authorization", "AUTH_HEADER")
	req.Header.Set("X-Customer-Id", "customer_id")
	req.Header.Set("X-Client-Bank-Id", "Bank01")
	req.Header.Set("X-Mock-Services", "true")
	req.Header.Set("Another-Header", "XYZ")

	initialForwardedHeaders := req.Context().Value(ForwardedHeadersKey{})

	assert.Equal(t, nil, initialForwardedHeaders)

	_, req = addHeadersToContext(w, req, headersToForward)

	forwardedHeaders := req.Context().Value(ForwardedHeadersKey{}).(ForwardedHeaders)

	assert.Equal(t, "AUTH_HEADER", forwardedHeaders["Authorization"])
	assert.Equal(t, "customer_id", forwardedHeaders["X-Customer-Id"])
	assert.Equal(t, "Bank01", forwardedHeaders["X-Client-Bank-Id"])
	assert.Equal(t, "true", forwardedHeaders["X-Mock-Services"])
	assert.Equal(t, "", forwardedHeaders["Another-Header"])
}

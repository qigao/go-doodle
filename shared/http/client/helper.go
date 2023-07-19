package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	http_request "http/request"

	"github.com/pkg/errors"
)

func apiURL(baseURL, path string, queryParams url.Values) (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = strings.TrimSuffix(u.Path, "/") + "/" + strings.TrimPrefix(path, "/")

	if queryParams != nil {
		u.RawQuery = queryParams.Encode()
	}

	return u, nil
}

func setHeaders(ctx context.Context, req *http.Request, headers []string) {
	req.Header.Set("Content-Type", "application/json")

	if ctx.Value(http_request.ForwardedHeadersKey{}) == nil {
		return
	}

	contextHeaders := ctx.Value(http_request.ForwardedHeadersKey{}).(http_request.ForwardedHeaders)

	for _, headers := range headers {
		val := contextHeaders[headers]
		if val == "" {
			continue
		}

		req.Header.Set(headers, val)
	}
}

func responseBody(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case 200, 201, 204:
		return body, nil
	case 401:
		return nil, &Error{
			ErrorCode: "401",
			Message:   "Unauthorized",
		}
	case 400, 422:
		apiError := &Error{}
		json.Unmarshal(body, apiError)

		return nil, apiError
	case 404:
		notFoundError := &NotFoundError{}
		json.Unmarshal(body, notFoundError)

		return nil, notFoundError
	default:
		return nil, errors.Errorf("client error: `%s`", response.Status)
	}
}

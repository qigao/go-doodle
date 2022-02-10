package http_client

import (
	"context"
	"github.com/go-resty/resty/v2"
	"http/http_request"
	"http/http_utils"
)

const applicationJSONMediaType = "application/json"

// forwardedHeaders contains common headers which are forwarded from a web request to http client requests
var forwardedHeaders = append(http_utils.GetInternalForwardedHeaders(), http_utils.GetTracingHeaders()...)

// NewClient returns a Client
func NewClient(baseURL string) *resty.Client {

	baseClient := resty.New()
	baseClient.SetBaseURL(baseURL)
	baseClient.OnBeforeRequest(setHeadersFromContext)

	return baseClient
}
func setHeadersFromContext(_ *resty.Client, req *resty.Request) error {
	ctx := req.Context()
	ctxForwardedHeaders := ctx.Value(http_request.ForwardedHeadersKey{})
	if ctxForwardedHeaders == nil {
		return nil
	}

	for _, header := range forwardedHeaders {
		headerFromContext := ctxForwardedHeaders.(http_request.ForwardedHeaders)[header]
		if headerFromContext == "" {
			continue
		}
		req.SetHeader(header, headerFromContext)
	}

	return nil
}

func NewRequest(ctx context.Context, client *resty.Client) *resty.Request {
	return client.R().
		SetContext(ctx).
		SetHeader("Content-Type", applicationJSONMediaType).
		SetHeader("Accept", applicationJSONMediaType).
		ForceContentType(applicationJSONMediaType).
		SetError(&Error{})
}

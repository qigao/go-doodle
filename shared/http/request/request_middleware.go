package request

import (
	"context"
	"net/http"
)

// ForwardedHeadersKey key in context
type ForwardedHeadersKey struct{}

// ForwardedHeaders structure of stored headers
type ForwardedHeaders map[string]string

// CreateHandler creates request middleware
func CreateHandler(next http.Handler, headersToForward []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w, r = addHeadersToContext(w, r, headersToForward)

		next.ServeHTTP(w, r)
	})
}

func addHeadersToContext(w http.ResponseWriter, r *http.Request, headersToForward []string) (http.ResponseWriter, *http.Request) {
	ctx := r.Context()
	values := make(ForwardedHeaders)

	for _, headerName := range headersToForward {
		values[headerName] = r.Header.Get(headerName)
	}

	updatedCtx := context.WithValue(ctx, ForwardedHeadersKey{}, values)

	r = r.WithContext(updatedCtx)

	return w, r
}

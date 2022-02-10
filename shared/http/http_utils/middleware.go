package http_utils

import "net/http"

// CombineMiddlewares wraps given slice of middlewares into the handler
func CombineMiddlewares(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

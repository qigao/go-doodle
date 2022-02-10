package http_utils

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockWriter struct {
	value string
}

func (w *mockWriter) Header() http.Header {
	return make(http.Header)
}

func (w *mockWriter) Write(body []byte) (int, error) {
	fmt.Println(string(body))
	w.value += "->"
	w.value += string(body)

	return 200, nil
}

func (w *mockWriter) WriteHeader(statusCode int) {
	return
}

func TestCombineMiddlewares(t *testing.T) {
	mainHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Main")
	})
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "First")
			next.ServeHTTP(w, r)
		})
	}

	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Second")

			next.ServeHTTP(w, r)
		})
	}

	t.Run("when combining without middlewares", func(t *testing.T) {
		combined := CombineMiddlewares(mainHandler)
		writer := &mockWriter{}
		req := &http.Request{}

		combined.ServeHTTP(writer, req)

		assert.Equal(t, "->Main", writer.value)
	})

	t.Run("when combining multiple middlewares", func(t *testing.T) {
		combined := CombineMiddlewares(mainHandler, middleware1, middleware2)
		writer := &mockWriter{}
		req := &http.Request{}

		combined.ServeHTTP(writer, req)

		assert.Equal(t, "->First->Second->Main", writer.value)
	})
}

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExternalForwardedHeaders(t *testing.T) {
	for _, h := range commonHeaders {
		assert.Contains(t, GetExternalForwardedHeaders(), h)
	}
	assert.Contains(t, GetExternalForwardedHeaders(), "X-Mock-Services")

	assert.Len(t, GetExternalForwardedHeaders(), 5)
}

func TestGetInternalForwardedHeaders(t *testing.T) {
	for _, h := range commonHeaders {
		assert.Contains(t, GetInternalForwardedHeaders(), h)
	}

	assert.Len(t, GetInternalForwardedHeaders(), 4)
}

func TestGetTracingHeaders(t *testing.T) {
	assert.Len(t, GetTracingHeaders(), 10)
}

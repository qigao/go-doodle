package utils

var commonHeaders = []string{
	"Authorization",
	"X-Customer-Id",
	"X-Client-Bank-Id",
	"X-Channel-Id",
}

// GetExternalForwardedHeaders returns list of headers
// need to be forwarded from frontend client to interaction framework
func GetExternalForwardedHeaders() []string {
	return append(commonHeaders, "X-Mock-Services")
}

// GetInternalForwardedHeaders returns list of headers
// need to be forwarded from interaction framework to back office
func GetInternalForwardedHeaders() []string {
	return commonHeaders
}

// GetTracingHeaders returns list of headers used for request tracing
func GetTracingHeaders() []string {
	return []string{
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"x-ot-span-context",
		"x-datadog-trace-id",
		"x-datadog-parent-id",
		"x-datadog-sampled",
	}
}

package logs

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestZeroLogWithConfig(t *testing.T) {
	e := echo.New()

	form := url.Values{}
	form.Add("username", "doejohn")

	req := httptest.NewRequest(echo.POST, "http://some?name=john", strings.NewReader(form.Encode()))

	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Header.Add("Referer", "http://foo.bar")
	req.Header.Add("User-Agent", "cli-agent")
	req.Header.Add(echo.HeaderXForwardedFor, "http://foo.bar")
	req.Header.Add("user", "admin")
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: "A1B2C3",
	})

	rec := httptest.NewRecorder()
	rec.Header().Add(echo.HeaderXRequestID, "123")

	c := e.NewContext(req, rec)
	b := new(bytes.Buffer)

	logger := log.Output(zerolog.ConsoleWriter{Out: b, NoColor: true})

	fields := DefaultZeroLogConfig.FieldMap
	fields["empty"] = ""
	fields["id"] = logID
	fields["path"] = logPath
	fields["protocol"] = logProtocol
	fields["referer"] = logReferer
	fields["user_agent"] = logUserAgent
	fields["repository"] = logHeaderPrefix + "repository"
	fields["filter_name"] = logQueryPrefix + "name"
	fields["username"] = logFormPrefix + "username"
	fields["session"] = logCookiePrefix + "session"
	fields["latency_human"] = logLatencyHuman
	fields["bytes_in"] = logBytesIn
	fields["bytes_out"] = logBytesOut
	fields["referer"] = logReferer
	fields["user"] = logHeaderPrefix + "user"

	config := ZeroLogConfig{
		Logger:   logger,
		FieldMap: fields,
	}

	_ = ZeroLogWithConfig(config)(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)

	res := b.String()

	tests := []struct {
		str string
		err string
	}{
		{"handle request", "invalid grpc: handle request info not found"},
		{"id=123", "invalid grpc: request id not found"},
		{`remote_ip=http://foo.bar`, "invalid grpc: remote ip not found"},
		{`uri=http://some?name=john`, "invalid grpc: uri not found"},
		{"host=some", "invalid grpc: host not found"},
		{"method=POST", "invalid grpc: method not found"},
		{"status=200", "invalid grpc: status not found"},
		{"latency=", "invalid grpc: latency not found"},
		{"latency_human=", "invalid grpc: latency_human not found"},
		{"bytes_in=0", "invalid grpc: bytes_in not found"},
		{"bytes_out=4", "invalid grpc: bytes_out not found"},
		{"path=/", "invalid grpc: path not found"},
		{"protocol=HTTP/1.1", "invalid grpc: protocol not found"},
		{`referer=http://foo.bar`, "invalid grpc: referer not found"},
		{"user_agent=cli-agent", "invalid grpc: user_agent not found"},
		{"user=admin", "invalid grpc: header user not found"},
		{"filter_name=john", "invalid grpc: query filter_name not found"},
		{"username=doejohn", "invalid grpc: form field username not found"},
		{"session=A1B2C3", "invalid grpc: cookie session not found"},
	}

	for _, test := range tests {
		if !strings.Contains(res, test.str) {
			t.Error(test.err)
		}
	}
}

func TestZeroLog(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/some", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_ = ZeroLog()(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)
}

func TestZeroLogWithEmptyConfig(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/some", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_ = ZeroLogWithConfig(ZeroLogConfig{})(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)
}

func TestZeroLogWithSkipper(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/some", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	config := DefaultZeroLogConfig
	config.Skipper = func(c echo.Context) bool {
		return true
	}

	_ = ZeroLogWithConfig(config)(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)
}

func TestZeroLogRetrievesAnError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/some", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	b := new(bytes.Buffer)

	logger := log.Output(zerolog.ConsoleWriter{Out: b, NoColor: true})

	config := ZeroLogConfig{
		Logger: logger,
	}

	_ = ZeroLogWithConfig(config)(func(c echo.Context) error {
		return errors.New("error")
	})(c)

	res := b.String()

	if !strings.Contains(res, "status=500") {
		t.Errorf("invalid grpc: wrong status code")
	}

	if !strings.Contains(res, `error=error`) {
		t.Errorf("invalid grpc: error not found")
	}
}

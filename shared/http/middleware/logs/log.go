package logs

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// Log middlewares constants.
const (
	logID           = "@id"
	logRemoteIP     = "@remote_ip"
	logURI          = "@uri"
	logHost         = "@host"
	logMethod       = "@method"
	logPath         = "@path"
	logProtocol     = "@protocol"
	logReferer      = "@referer"
	logUserAgent    = "@user_agent"
	logStatus       = "@status"
	logError        = "@error"
	logLatency      = "@latency"
	logLatencyHuman = "@latency_human"
	logBytesIn      = "@bytes_in"
	logBytesOut     = "@bytes_out"
	logHeaderPrefix = "@header:"
	logQueryPrefix  = "@query:"
	logFormPrefix   = "@form:"
	logCookiePrefix = "@cookie:"
)

type LogCtl struct {
	// Enable console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the grpc framework grpc JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework grpc to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to grpc to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

// string to int base conversion.
const base = 10

// mapFields maps fields based on tag name.
func mapFields(ec echo.Context, h echo.HandlerFunc, fm map[string]string) (map[string]interface{}, error) {
	logFields := map[string]interface{}{}
	start := time.Now()

	err := h(ec)
	if err != nil {
		ec.Error(err)
	}

	elapsed := time.Since(start)
	tags := mapTags(ec, elapsed)

	if err != nil {
		tags[logError] = err
	}

	for k, tag := range fm {
		if tag == "" {
			continue
		}

		if value, ok := tags[tag]; ok {
			logFields[k] = value
			continue
		}

		switch {
		case strings.HasPrefix(tag, logHeaderPrefix):
			key := tag[len(logHeaderPrefix):]
			logFields[k] = ec.Request().Header.Get(key)
		case strings.HasPrefix(tag, logQueryPrefix):
			key := tag[len(logQueryPrefix):]
			logFields[k] = ec.QueryParam(key)
		case strings.HasPrefix(tag, logFormPrefix):
			key := tag[len(logFormPrefix):]
			logFields[k] = ec.FormValue(key)
		case strings.HasPrefix(tag, logCookiePrefix):
			key := tag[len(logCookiePrefix):]
			cookie, err := ec.Cookie(key)
			if err == nil {
				logFields[k] = cookie.Value
			}
		}
	}

	return logFields, err
}

// mapTags maps the grpc tags with its related data. Populate previously the
// key/value avoids the cyclomatic complexity of the grpc middlewares to
// identify each tag and value.
func mapTags(ec echo.Context, latency time.Duration) map[string]interface{} {
	tags := map[string]interface{}{}

	req := ec.Request()
	res := ec.Response()

	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}

	tags[logID] = id
	tags[logRemoteIP] = ec.RealIP()
	tags[logURI] = req.RequestURI
	tags[logHost] = req.Host
	tags[logMethod] = req.Method

	path := req.URL.Path
	if path == "" {
		path = "/"
	}

	tags[logPath] = path
	tags[logProtocol] = req.Proto
	tags[logReferer] = req.Referer()
	tags[logUserAgent] = req.UserAgent()
	tags[logStatus] = res.Status
	tags[logLatency] = strconv.FormatInt(int64(latency), base)
	tags[logLatencyHuman] = latency.String()

	cl := req.Header.Get(echo.HeaderContentLength)
	if cl == "" {
		cl = "0"
	}

	tags[logBytesIn] = cl
	tags[logBytesOut] = strconv.FormatInt(res.Size, base)

	return tags
}

func newRollingFile(cfg LogCtl) io.Writer {
	return &lumberjack.Logger{
		Filename:   path.Join(cfg.Directory, cfg.Filename),
		MaxBackups: cfg.MaxBackups, // files
		MaxSize:    cfg.MaxSize,    // megabytes
		MaxAge:     cfg.MaxAge,     // days
	}
}

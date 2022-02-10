package grpc

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"path"
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

func newRollingFile(cfg LogCtl) io.Writer {
	return &lumberjack.Logger{
		Filename:   path.Join(cfg.Directory, cfg.Filename),
		MaxBackups: cfg.MaxBackups, // files
		MaxSize:    cfg.MaxSize,    // megabytes
		MaxAge:     cfg.MaxAge,     // days
	}
}

package grpc

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/grpclog"
)

// GrpcLogSetNewZeroLogger sets grpclog to a new GrpcZeroLogger.
func GrpcLogSetNewZeroLogger() {
	grpclog.SetLoggerV2(NewGrpcZeroLogger(zerolog.Logger{}))
}

// GrpcLogSetZeroLogger sets grpclog to a GrpcZeroLogger.
func GrpcLogSetZeroLogger(logger GrpcZeroLogger) {
	grpclog.SetLoggerV2(logger)
}

// NewZeroLogger returns a new logger logger with config
func NewZeroLogger(directory, filename string) zerolog.Logger {
	config := LogCtl{
		ConsoleLoggingEnabled: false,
		EncodeLogsAsJson:      false,
		FileLoggingEnabled:    true,
		Directory:             directory,
		Filename:              filename,
		MaxSize:               2,
		MaxBackups:            30,
		MaxAge:                30,
	}
	return setupLogFilePolicy(config)
}

func setupLogFilePolicy(cfg LogCtl) zerolog.Logger {
	var writers []io.Writer

	if cfg.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if cfg.FileLoggingEnabled {
		writers = append(writers, newRollingFile(cfg))
	}
	multiWriter := io.MultiWriter(writers...)

	// grpc_zerolog.SetGlobalLevel(grpc_zerolog.DebugLevel)
	logger := zerolog.New(multiWriter).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", cfg.FileLoggingEnabled).
		Bool("jsonLogOutput", cfg.EncodeLogsAsJson).
		Str("logDirectory", cfg.Directory).
		Str("fileName", cfg.Filename).
		Int("maxSizeMB", cfg.MaxSize).
		Int("maxBackups", cfg.MaxBackups).
		Int("maxAgeInDays", cfg.MaxAge).
		Msg("logging configured")
	return logger
}

// GrpcZeroLogger transforms grpc grpc calls to Zerolog logger.
type GrpcZeroLogger struct {
	log zerolog.Logger
}

// NewGrpcZeroLogger creates a new GrpcZeroLogger
func NewGrpcZeroLogger(logger zerolog.Logger) GrpcZeroLogger {
	return GrpcZeroLogger{log: logger}
}

// Fatal fatals arguments.
func (l GrpcZeroLogger) Fatal(args ...interface{}) {
	l.log.Fatal().Msg(fmt.Sprint(args...))
}

// Fatalf fatals formatted string with arguments.
func (l GrpcZeroLogger) Fatalf(format string, args ...interface{}) {
	l.log.Fatal().Msgf(format, args...)
}

// Fatalln fatals and new line.
func (l GrpcZeroLogger) Fatalln(args ...interface{}) {
	l.Fatal(args...)
}

// Error errors arguments.
func (l GrpcZeroLogger) Error(args ...interface{}) {
	l.log.Error().Msg(fmt.Sprint(args...))
}

// Errorf errors formatted string with arguments.
func (l GrpcZeroLogger) Errorf(format string, args ...interface{}) {
	l.log.Error().Msgf(format, args...)
}

// Errorln errors and new line.
func (l GrpcZeroLogger) Errorln(args ...interface{}) {
	l.Error(args...)
}

// Info infos arguments.
func (l GrpcZeroLogger) Info(args ...interface{}) {
	l.log.Info().Msg(fmt.Sprint(args...))
}

// Infof infos formatted string with arguments.
func (l GrpcZeroLogger) Infof(format string, args ...interface{}) {
	l.log.Info().Msgf(format, args...)
}

// Infoln infos and new line.
func (l GrpcZeroLogger) Infoln(args ...interface{}) {
	l.Info(args...)
}

// Warning warns arguments.
func (l GrpcZeroLogger) Warning(args ...interface{}) {
	l.log.Warn().Msg(fmt.Sprint(args...))
}

// Warningf warns formatted string with arguments.
func (l GrpcZeroLogger) Warningf(format string, args ...interface{}) {
	l.log.Warn().Msgf(format, args...)
}

// Warningln warns and new line.
func (l GrpcZeroLogger) Warningln(args ...interface{}) {
	l.Warning(args...)
}

// Print logs arguments.
func (l GrpcZeroLogger) Print(args ...interface{}) {
	l.Info(args...)
}

// Printf logs formatted string with arguments.
func (l GrpcZeroLogger) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}

// Println logs with new line.
func (l GrpcZeroLogger) Println(args ...interface{}) {
	l.Infoln(args...)
}

// V determines Verbosity Level.
func (l GrpcZeroLogger) V(level int) bool {
	switch level {
	case 0:
		return zerolog.InfoLevel <= zerolog.GlobalLevel()
	case 1:
		return zerolog.WarnLevel <= zerolog.GlobalLevel()
	case 2:
		return zerolog.ErrorLevel <= zerolog.GlobalLevel()
	case 3:
		return zerolog.FatalLevel <= zerolog.GlobalLevel()
	default:
		panic("unhandled gRPC logger level")
	}
}

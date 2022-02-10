package logs

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/stats/view"
)

// This example registers the ZeroLog middleware with default configuration.
func ExampleZeroLog() {
	e := echo.New()

	// Middleware
	e.Use(ZeroLog())
}

// This example registers the ZeroLog middleware with custom configuration.
func ExampleZeroLogWithConfig() {
	e := echo.New()

	// Custom logger logger instance
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Middleware
	logConfig := ZeroLogConfig{
		Logger: logger,
		FieldMap: map[string]string{
			"uri":    "@uri",
			"host":   "@host",
			"method": "@method",
			"status": "@status",
		},
	}

	e.Use(ZeroLogWithConfig(logConfig))
}

// This example registers the Logrus middleware with default configuration.
func ExampleLogrus() {
	e := echo.New()

	// Middleware
	e.Use(Logrus())
}

// This example registers the Logrus middleware with custom configuration.
func ExampleLogrusWithConfig() {
	e := echo.New()

	// Custom logrus logger instance
	logger := logrus.New()

	// Middleware
	logConfig := LogrusConfig{
		Logger: logger,
		FieldMap: map[string]string{
			"uri":    "@uri",
			"host":   "@host",
			"method": "@method",
			"status": "@status",
		},
	}

	e.Use(LogrusWithConfig(logConfig))
}

// This example registers the OpenCensus middleware with default configuration.
func ExampleOpenCensus() {
	e := echo.New()

	// Middleware
	e.Use(OpenCensus())
}

// This example registers the OpenCensus middleware with custom configuration.
func ExampleOpenCensusWithConfig() {
	e := echo.New()

	// Middleware
	cfg := OpenCensusConfig{
		Views: []*view.View{
			OpenCensusRequestCount,
		},
	}

	e.Use(OpenCensusWithConfig(cfg))
}

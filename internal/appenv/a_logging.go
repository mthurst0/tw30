package appenv

import (
	"go.uber.org/zap"
)

// WARNING: file is called "a_logging.go" to ensure its init runs
// before other files in the package. In env.go, there are calls
// to zap.L().Fatal that we want to fail if we are starting without
// a proper environment. In the absence of this init running before
// those calls Fatal calls in env.go will fail silently.

// ConfigureLogging configures Zap logging.
func ConfigureLogging() (func(), error) {
	factory := zap.NewProduction
	logger, err := factory()
	if err != nil {
		return nil, err
	}
	return zap.ReplaceGlobals(logger), nil
}

// MustConfigureLogging configures Zap logging. It will panic on error.
// The function to undo the logging configuration is returned.
func MustConfigureLogging() func() {
	undo, err := ConfigureLogging()
	if err != nil {
		panic(err)
	}
	return undo
}

func init() {
	MustConfigureLogging()
}

// MustConfigureDevelopmentLogging configures Zap logging for development.
func MustConfigureDevelopmentLogging() {
	factory := zap.NewDevelopment
	logger, err := factory()
	if err != nil {
		panic(err)
	}
	_ = zap.ReplaceGlobals(logger)
}

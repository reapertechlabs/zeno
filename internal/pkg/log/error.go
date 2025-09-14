package log

import "errors"

// ErrLoggerAlreadyInitialized is the error returned when the logger is already initialized
var ErrLoggerAlreadyInitialized = errors.New("logger already initialized")

package postprocessor

import "errors"

// ErrPostprocessorAlreadyInitialized is the error returned when the postprocessor is already initialized
var ErrPostprocessorAlreadyInitialized = errors.New("postprocessor already initialized")

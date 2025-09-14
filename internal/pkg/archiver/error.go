package archiver

import "errors"

// ErrArchiverAlreadyInitialized is the error returned when the preprocess is already initialized
var ErrArchiverAlreadyInitialized = errors.New("archiver already initialized")

package unity

import "errors"

var (
	ErrorGameWorldNotFound      = errors.New("GameWorld not found")
	ErrorEngineStringReadFailed = errors.New("could not read Unity Engine String")
)

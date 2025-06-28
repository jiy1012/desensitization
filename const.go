package desensitization

import "errors"

const (
	TagKey = "desensitization"
)

var (
	ErrTypeNotFound      = errors.New("desensitizer type not found")
	ErrParamsShouldBePtr = errors.New("params should be pointer")
)

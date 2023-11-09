package desensitization

import "errors"

const (
	DESENSITIZATION_TYPE = "desensitization"
)

var (
	ErrTypeNotFound      = errors.New("desensitizer type not found")
	ErrParamsShouldBePtr = errors.New("params should be pointer")
)

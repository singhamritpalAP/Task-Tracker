package constants

import "errors"

var (
	ErrUnableToBindJson    = errors.New("unable to bind json")
	ErrTaskIdRequired      = errors.New("task id is required")
	ErrTitleRequired       = errors.New("title is required and must be at least 2 characters long")
	ErrDescriptionRequired = errors.New("description is required")
	ErrInvalidStatus       = errors.New("invalid status")
)

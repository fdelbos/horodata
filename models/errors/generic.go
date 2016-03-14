package errors

// see http://www.postgresql.org/docs/current/static/errcodes-appendix.html

import (
	"errors"
)

var (
	NotFound        = errors.New("Not found")
	UniqueViolation = errors.New("Unique Violation")
	InvalidSchema   = errors.New("Invalid schema")
)

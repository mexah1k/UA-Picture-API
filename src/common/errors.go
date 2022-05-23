package common

import "errors"

var dbs int

var (
	ItemNotFoundError = errors.New("item not found")
	UnexpectedError   = errors.New("unexpected error")
)

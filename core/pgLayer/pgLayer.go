package pgLayer

import "errors"

type PGLayer interface {
	// 기능에 따라 분류
}

// customError
var RecordNotFoundError = errors.New("record not found")

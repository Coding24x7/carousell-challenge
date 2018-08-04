package lib

import (
	"fmt"
)

// ErrCode is type of common errors
type ErrCode uint32

// Errors common to all the core modules.
const (
	ErrUnknown ErrCode = 0x0

	ErrNotFound = 0x01

	// User
	ErrUserNotFound        = 0x10
	ErrUserWrongCredential = 0x11
	ErrUserWrongKey        = 0x12

	// Topic
	ErrTopicNotFound = 0x30
	ErrTopicWrongKey = 0x31
)

var errCodeName = map[ErrCode]string{
	ErrUnknown: "ERR_UNKNOWN",

	// Generic

	ErrNotFound: "NOT_FOUND",

	// User
	ErrUserNotFound:        "USER_NOT_FOUND",
	ErrUserWrongCredential: "USER_WRONG_CREDENTIAL",
	ErrUserWrongKey:        "USER_WRONG_INPUT",

	// Topic
	ErrTopicNotFound: "TOPIC_NOT_FOUND",
	ErrTopicWrongKey: "TOPIC_WRONG_INPUT",
}

// BadRequesErr is bad request error
type BadRequesErr struct {
	code    ErrCode
	message string
}

// NewBadReqErr creates new bad request error
func NewBadReqErr(code ErrCode, message string) *BadRequesErr {
	return &BadRequesErr{code, message}
}

func (e BadRequesErr) Error() string {
	if s, ok := errCodeName[e.code]; ok {
		return fmt.Sprintf("%d : %s : %s", e.code, s, e.message)
	}
	return fmt.Sprintf("0x%x:: Unknown error code", uint32(e.code))
}

// InternalErr is internal server error
type InternalErr struct {
	code    ErrCode
	message string
}

// NewInternalErr creates new internal server error
func NewInternalErr(code ErrCode, message string) *InternalErr {
	return &InternalErr{code, message}
}

func (e InternalErr) Error() string {
	if s, ok := errCodeName[e.code]; ok {
		return fmt.Sprintf("%d : %s : %s", e.code, s, e.message)
	}
	return fmt.Sprintf("0x%x:: Unknown error code", uint32(e.code))
}

// AuthErr is internal authentication error
type AuthErr struct {
	code    ErrCode
	message string
}

// NewAuthErr creates new authentication error
func NewAuthErr(code ErrCode, message string) *AuthErr {
	return &AuthErr{code, message}
}

func (e AuthErr) Error() string {
	if s, ok := errCodeName[e.code]; ok {
		return fmt.Sprintf("%d : %s : %s", e.code, s, e.message)
	}
	return fmt.Sprintf("0x%x:: Unknown error code", uint32(e.code))
}

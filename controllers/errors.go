package controllers

import (
	"github.com/Coding24x7/carousell-challenge/lib"
)

// errCtx is a general error context
type errCtx interface {
	BadRequest(string) error
	InternalServerError(string) error
	Unauthorized(string) error
}

// processErr processes general errors
func processErr(ctx errCtx, err error) error {

	_, ok := err.(*lib.BadRequesErr)
	if ok {
		return ctx.BadRequest(err.Error())
	}

	_, ok = err.(*lib.AuthErr)
	if ok {
		return ctx.Unauthorized(err.Error())
	}

	return ctx.InternalServerError(err.Error())
}

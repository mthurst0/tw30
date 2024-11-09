package apperr

import (
	"errors"
	"net/http"

	"go.uber.org/zap"
)

func ToHTTPStatus(err error) int {
	if IsNotFound(err) {
		return http.StatusNotFound
	}
	if IsAlreadyExists(err) {
		return http.StatusBadRequest
	}
	if IsBadRequest(err) {
		return http.StatusBadRequest
	}
	if IsInternal(err) {
		return http.StatusInternalServerError
	}
	if IsUnauthorized(err) {
		return http.StatusUnauthorized
	}
	zap.L().Error("could not translate error to HTTP status", zap.Error(err))
	return http.StatusInternalServerError
}

type NotFound struct {
	error
}

func NewNotFound(s string) error {
	if s == "" {
		return NotFound{error: errors.New("not found")}
	}
	return NotFound{error: errors.New(s)}
}

func NewNotFoundE(err error) error {
	return NotFound{error: err}
}

func IsNotFound(err error) bool {
	return errors.As(err, &NotFound{})
}

type Internal struct {
	error
}

func IsInternal(err error) bool {
	return errors.As(err, &Internal{})
}

func NewInternal(s string) error {
	if s == "" {
		return Internal{error: errors.New("internal error")}
	}
	return Internal{error: errors.New(s)}
}

func NewInternalE(err error) error {
	return Internal{error: err}
}

type BadRequest struct {
	error
}

func IsBadRequest(err error) bool {
	return errors.As(err, &BadRequest{})
}

func NewBadRequest(s string) error {
	if s == "" {
		return BadRequest{error: errors.New("bad request")}
	}
	return BadRequest{error: errors.New(s)}
}

func NewBadRequestE(err error) error {
	return BadRequest{error: err}
}

type AlreadyExists struct {
	error
}

func IsAlreadyExists(err error) bool {
	return errors.As(err, &AlreadyExists{})
}

func NewAlreadyExists(s string) error {
	if s == "" {
		return AlreadyExists{error: errors.New("already exists")}
	}
	return AlreadyExists{error: errors.New(s)}
}

func NewAlreadyExistsE(err error) error {
	return AlreadyExists{error: err}
}

type Unauthorized struct {
	error
}

func IsUnauthorized(err error) bool {
	return errors.As(err, &Unauthorized{})
}

func NewUnauthorized(s string) error {
	if s == "" {
		return Unauthorized{error: errors.New("unauthorized")}
	}
	return Unauthorized{error: errors.New(s)}
}

func NewUnauthorizedE(err error) error {
	return Unauthorized{error: err}
}

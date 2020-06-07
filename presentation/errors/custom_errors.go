package errors

import (
	"encoding/json"
	"errors"
)

var (
	InvalidParam = errors.New("InvalidParam")
)

type ApplicationError struct {
	Name string `json:"name"`
	Msg string `json:"msg"`
}

func (a ApplicationError) Error() string {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func InvalidParamError(msg string) error {
	return ApplicationError{
		Name: InvalidParam.Error(),
		Msg:  msg,
	}
}

func InternalServerError() error {
	return  ApplicationError{
		Name: "InternalServerError",
		Msg:  "internal server error",
	}
}

package controllers

import (
	"clean-arch/models"
	"clean-arch/presentation/errors"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestSignupController(t *testing.T) {

	controller := NewSignupController()

	t.Run("should return 400 if no name is provided", func(t *testing.T) {
		a := models.AccountModel{
			Email:                "john@doe.com",
			Password:             "123123",
			PasswordConfirmation: "123123",
		}

		r := CreateTestRequestJSON(a)
		response := controller.Handle(r)
		if response.Code != http.StatusBadRequest {
			t.Error("Expected Code 400")
		}

		err := UnmarshalApplicationError(response)
		if err.Name != errors.InvalidParam.Error() {
			t.Error("Expect InvalidParam")
		}
	})

	t.Run("should return 400 if no password is provided", func(t *testing.T) {
		a := models.AccountModel{
			Name:                 "John",
			Email:                "john@doe.com",
			PasswordConfirmation: "123123",
		}

		r := CreateTestRequestJSON(a)
		response := controller.Handle(r)
		if response.Code != http.StatusBadRequest {
			t.Error("Expected Code 400")
		}

		err := UnmarshalApplicationError(response)
		if err.Name != errors.InvalidParam.Error() {
			t.Error("Expect InvalidParam")
		}
	})

	t.Run("should return 400 if no email is provided", func(t *testing.T) {
		a := models.AccountModel{
			Name:                 "John",
			Password:             "123123",
			PasswordConfirmation: "123123",
		}

		r := CreateTestRequestJSON(a)
		response := controller.Handle(r)
		if response.Code != http.StatusBadRequest {
			t.Error("Expected Code 400")
		}

		err := UnmarshalApplicationError(response)
		if err.Name != errors.InvalidParam.Error() {
			t.Error("Expect InvalidParam")
		}
	})

	t.Run("should return 400 if no passwordConfirmation is provided", func(t *testing.T) {
		a := models.AccountModel{
			Name:     "John",
			Email:    "john@doe.com",
			Password: "123123",
		}

		r := CreateTestRequestJSON(a)
		response := controller.Handle(r)
		if response.Code != http.StatusBadRequest {
			t.Error("Expected Code 400")
		}

		err := UnmarshalApplicationError(response)
		if err.Name != errors.InvalidParam.Error() {
			t.Error("Expect InvalidParam")
		}
	})
}

func UnmarshalApplicationError(response models.HttpResponse) errors.ApplicationError {
	var err errors.ApplicationError
	s := fmt.Sprintf("%s", response.Body)
	e := json.Unmarshal([]byte(s), &err)
	if e != nil {
		panic(e.Error())
	}
	return err
}

func CreateTestRequestJSON(a interface{}) models.HttpRequest {
	b, err := json.Marshal(a)
	if err != nil {
		panic("Json encoding failed")
	}

	r := strings.NewReader(string(b))
	request := models.HttpRequest{Body: r}
	return request
}

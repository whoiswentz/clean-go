package controllers

import (
	errors "clean-arch/delivery/errors"
	"clean-arch/models"
	"encoding/json"
)

type SignupController struct{}

func NewSignupController() *SignupController {
	return &SignupController{}
}

func (s SignupController) Handle(r models.HttpRequest) models.HttpResponse {
	var a models.AccountModel
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return InternalServer()
	}

	if a.IsNameEmpty() {
		return BadRequest(errors.InvalidParamError("name"))
	}

	if a.IsEmailEmpty() {
		return BadRequest(errors.InvalidParamError("email"))
	}

	if a.IsPasswordEmpty() {
		return BadRequest(errors.InvalidParamError("password"))
	}

	if a.IsPasswordConfirmationEmpty() {
		return BadRequest(errors.InvalidParamError("passwordConfirmation"))
	}

	return OK(a)
}

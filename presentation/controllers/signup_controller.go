package controllers

import (
	"clean-arch/models"
	presentation "clean-arch/presentation/errors"
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
		return BadRequest(presentation.InvalidParamError("name"))
	}

	if a.IsEmailEmpty() {
		return BadRequest(presentation.InvalidParamError("email"))
	}

	if a.IsPasswordEmpty() {
		return BadRequest(presentation.InvalidParamError("password"))
	}

	if a.IsPasswordConfirmationEmpty() {
		return BadRequest(presentation.InvalidParamError("passwordConfirmation"))

	}

	return OK(a)
}

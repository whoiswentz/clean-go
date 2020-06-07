package controllers

import (
	"clean-arch/models"
	presentation "clean-arch/presentation/errors"
	"encoding/json"
)

type SignupController struct {}

func NewSignupController() *SignupController {
	return &SignupController{}
}

func (s SignupController) handle(r models.HttpRequest) models.HttpResponse {
	var a models.AccountModel
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return InternalServer()
	}

	if a.IsNameEmpty() {
		return BadRequest(presentation.InvalidParamError("invalid param"))
	}

	return OK(a)
}
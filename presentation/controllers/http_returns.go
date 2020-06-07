package controllers

import (
	"clean-arch/models"
	app "clean-arch/presentation/errors"
	"net/http"
)

func OK(data interface{}) models.HttpResponse {
	return models.HttpResponse{
		Code: http.StatusOK,
		Body: data,
	}
}

func BadRequest(err error) models.HttpResponse {
	return models.HttpResponse{
		Code: http.StatusBadRequest,
		Body: err.Error(),
	}
}

func InternalServer() models.HttpResponse {
	err := app.InternalServerError()

	return models.HttpResponse{
		Code: http.StatusInternalServerError,
		Body: err.Error(),
	}
}
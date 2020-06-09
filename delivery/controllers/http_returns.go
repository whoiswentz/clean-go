package controllers

import (
	app "clean-arch/delivery/errors"
	"clean-arch/models"
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

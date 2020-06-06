package controllers

import (
	"net/http"
	"testing"
)

func TestSignupController_ShouldReturn200(t *testing.T) {
	controller := SignupController{}
	request := HttpRequest{Body: nil}

	response := controller.handle(request)
	if response.Code != http.StatusOK {
		t.Error("Expected Code 200")
	}
}
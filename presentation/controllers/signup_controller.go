package controllers

type Controller interface {
	handle(request HttpRequest) HttpResponse
}

type HttpRequest struct {
	Body interface{}
}

type HttpResponse struct {
	Code int16
	Body interface{}
}

type SignupController struct {}

func (s SignupController) handle(r HttpRequest) HttpResponse {
	return HttpResponse{
		Code: 200,
		Body: nil,
	}
}
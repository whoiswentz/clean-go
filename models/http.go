package models

import "io"

type HttpRequest struct {
	Body io.Reader `json:"body"`
}

type HttpResponse struct {
	Code int16 `json:"code"`
	Body interface{} `json:"body"`
}


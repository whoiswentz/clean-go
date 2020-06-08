package protocols

import (
	"clean-arch/models"
)

type Controller interface {
	Handle(request models.HttpRequest) models.HttpResponse
}

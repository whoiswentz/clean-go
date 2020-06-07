package protocols

import (
	"clean-arch/models"
)

type Controller interface {
	handle(request models.HttpRequest) models.HttpResponse
}

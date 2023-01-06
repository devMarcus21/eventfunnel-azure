package eventingress

import (
	"github.com/devMarcus21/eventfunnel-azure/pkg/utils/events"
	response "github.com/devMarcus21/eventfunnel-azure/pkg/utils/responses"
)

func CreateEventIngressHandler() func(events.Event) response.ServiceResponse {
	return func(event events.Event) response.ServiceResponse {
		return response.ServiceResponse{
			"status":  "sucess",
			"payload": event,
		}
	}
}

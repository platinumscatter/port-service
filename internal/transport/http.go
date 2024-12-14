package transport

import (
	"context"
	"net/http"

	"github.com/platinumscatter/port-service/internal/domain"
	"github.com/platinumscatter/port-service/internal/common/server"
)

type PortService interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
}

type HttpServer struct {
	service PortService
}

func NewHttpService(service PortService) HttpServer {
	return HttpServer{
		service: service,
	}
}

func (h HttpServer) GetPort(w http.ResponseWriter, r *http.Request) {
	port, err := h.service.GetPort(r.Context(), r.URL.Query().Get("id"))
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}
	response := Port{
		ID:          port.ID(),
		Name:        port.Name(),
		City:        port.City(),
		Country:     port.Country(),
		Alias:       port.Alias(),
		Regions:     port.Regions(),
		Coordinates: port.Coordinates(),
		Province:    port.Province(),
		Timezone:    port.Timezone(),
		Unlocs:      port.Unlocs(),
	}
	server.RespondOK(response, w, r)
}

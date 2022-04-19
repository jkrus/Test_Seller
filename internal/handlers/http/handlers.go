package http

import (
	"net/http"

	"github.com/jkrus/Test_Seller/internal/announcement"
	"github.com/jkrus/Test_Seller/internal/services"
	"github.com/jkrus/Test_Seller/pkg/server"
)

type (
	handlers struct {
		r *http.ServeMux

		announcement *announcement.Handler
	}
)

func NewHandlers(service *services.Services) server.Handlers {
	return &handlers{
		r:            http.NewServeMux(),
		announcement: announcement.NewHandler(service.Announcement),
	}
}

func (h *handlers) Get() http.Handler {
	return h.r
}

func (h *handlers) Register() {
	r := http.NewServeMux()

	r.HandleFunc("/api/v1", hello)

	h.announcement.Register(r)

	h.r.Handle("/", r)

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from HTTP"))
}

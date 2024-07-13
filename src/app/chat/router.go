package chat

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.Chat)

	return r
}

package http

import (
	"context"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Route(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	charhandler, err := InitializeChatHandler(ctx)
	if err != nil {
		panic(err)
	}

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/chat", charhandler.Router())
	})

	return r
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	httpkg "ai-chat/src/http"
)

type ops func(context.Context) error

func gracefulShutdown(ctx context.Context, operations map[string]ops) <-chan struct{} {
	done := make(chan struct{})
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for name, op := range operations {
		wg.Add(1)
		go func(name string, op ops) {
			defer wg.Done()
			if err := op(ctx); err != nil {
				panic(err)
			}
		}(name, op)
	}

	wg.Wait()
	close(done)

	return done
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := httpkg.Route(ctx)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-sig
	wait := gracefulShutdown(ctx, map[string]ops{
		"server": func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	<-wait
	log.Println("server stopped")
}

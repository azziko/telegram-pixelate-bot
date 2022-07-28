package app

import (
	"context"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server        *http.Server
	updateHandler http.Handler
}

func NewApp(s *http.Server, h http.Handler) *App {
	return &App{
		server:        s,
		updateHandler: h,
	}
}

func (a *App) Run() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	http.Handle("/", a.updateHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))

	//gracefully shutting down
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM)
	signal.Notify(exit, syscall.SIGINT)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		<-exit
		log.Println("interruption triggered")
		if err := a.server.Shutdown(ctx); err != nil {
			log.Printf("shutdown error: %v\n", err)
		}

		log.Println("shut down complete")
	}()

	log.Println("Listenning on port", a.server.Addr)
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start a server: %v\n", err)
	}
}

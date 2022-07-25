package app

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
)

type App struct {
	PORT    string
	Handler http.Handler
}

func NewApp(p string, h http.Handler) *App {
	return &App{
		PORT:    p,
		Handler: h,
	}
}

func (a *App) Start() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	http.Handle("/", a.Handler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))

	fmt.Println("Listenning on port", a.PORT)
	if err := http.ListenAndServe(":"+a.PORT, nil); err != nil {
		log.Fatal(err)
	}
}

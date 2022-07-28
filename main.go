package main

import (
	"fmt"
	"net/http"
	"os"
	"pixelate/app"
	"pixelate/app/client/telegram"
	"pixelate/app/handler"
	"pixelate/app/storage/img"
)

const (
	host           = "api.telegram.org"
	storageDirPath = "./img/"
)

func main() {

	var token = os.Getenv("TOKEN")
	var port = os.Getenv("PORT")

	tgClient := telegram.NewClient(host, token)
	storage := img.NewStorage(storageDirPath)

	if err := storage.Init(); err != nil {
		panic("Could not initiate storage: " + err.Error())
	}
	fmt.Println("Storage initiated successfuly!")

	handler := handler.NewUpdateHandler(tgClient, storage)

	server := &http.Server{
		Addr: ":" + port,
	}

	app := app.NewApp(server, handler)
	app.Run()
}

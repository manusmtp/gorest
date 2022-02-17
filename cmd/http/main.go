package main

import (
	"github.com/manusmtp/cmd/http/handler"
	"github.com/manusmtp/cmd/internal/container"
)

func main() {

	di := container.Inject()
	//delegate to start our Http Server
	handler.StartServer(di)
}

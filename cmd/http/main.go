package main

import (
	"github.com/manusmtp/gorest/cmd/http/handler"
	"github.com/manusmtp/gorest/internal/container"
)

func main() {

	di := container.Inject()
	//delegate to start our Http Server
	handler.StartServer(di)
}

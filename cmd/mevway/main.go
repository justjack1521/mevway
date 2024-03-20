package main

import (
	"context"
	"fmt"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevium/pkg/server"
	"mevway/internal/app"
	"mevway/internal/service"
	"net/http"
)

func main() {

	ctx := context.Background()

	fmt.Println("teehee")

	application := service.NewApplication(ctx)

	go application.WebServer.Server.Run()

	server.RunHTTPServer("8080", func(server *http.Server) {
		server.Handler = application.Engine
	})

	defer func(application app.Application) {
		application.EventPublisher.Notify(mevent.ApplicationShutdownEvent{})
	}(application)

}

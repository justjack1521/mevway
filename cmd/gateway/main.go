package main

import (
	"context"
	"github.com/jinzhu/configor"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevium/pkg/server"
	"mevway/internal/app"
	"mevway/internal/config"
	"mevway/internal/service"
	"net/http"
)

func main() {

	ctx := context.Background()

	var conf config.Application

	if err := configor.Load(&conf, "/home/xdiradmin/go/src/mevway/internal/config/config.dev.json"); err != nil {
		panic(err)
	}

	application := service.NewApplication(ctx, conf)

	go application.WebServer.Server.Run()

	server.RunHTTPServer("8080", func(server *http.Server) {
		server.Handler = application.Engine
	})

	defer func(application app.Application) {
		application.EventPublisher.Notify(mevent.ApplicationShutdownEvent{})
	}(application)

}

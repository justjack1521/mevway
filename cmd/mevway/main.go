package main

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevium/pkg/server"
	"log"
	"mevway/internal/app"
	"mevway/internal/service"
	"net/http"
	"os"
	"runtime/pprof"
)

func main() {

	ctx := context.Background()

	application := service.NewApplication(ctx)

	var cpuprofile = "mevway.prof"

	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	go application.WebServer.Server.Run()

	server.RunHTTPServer("8080", func(server *http.Server) {
		server.Handler = application.Engine
	})

	defer func(application app.Application) {
		application.EventPublisher.Notify(mevent.ApplicationShutdownEvent{})
	}(application)

}

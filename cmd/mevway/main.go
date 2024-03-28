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
	"os/signal"
	"runtime/pprof"
	"syscall"
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

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		pprof.StopCPUProfile()
		os.Exit(1)
	}()

	go application.WebServer.Server.Run()

	server.RunHTTPServer("8080", func(server *http.Server) {
		server.Handler = application.Engine
	})

	defer func(application app.Application) {
		application.EventPublisher.Notify(mevent.ApplicationShutdownEvent{})
	}(application)

}

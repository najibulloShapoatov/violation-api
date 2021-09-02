package main

import (
	// "encoding/json"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "mobile-rest-api/app"
	"mobile-rest-api/config"
	"mobile-rest-api/libs/initLogger"
	//"mobile-rest-api/models"
	"mobile-rest-api/routes"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	//load Config
	config.Init()
	var cfg = config.Get()

	initLogger.InitLog(cfg.Logfile)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	var router = routes.Init()

	headers := handlers.AllowedHeaders([]string{"*"})
	methods := handlers.AllowedMethods([]string{"*"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//migrate
	//models.Migrate()

	//fmt.Println("Listining on host: ", cfg.Server.Host, "\tport:", cfg.Server.Port, "\t")

	log.Println("Listining on host: ", cfg.Server.Host, "\tport:", cfg.Server.Port, "\t")

	srv := &http.Server{
		Addr: cfg.Server.Host + ":" + cfg.Server.Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handlers.CORS(headers, methods, origins)(router),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

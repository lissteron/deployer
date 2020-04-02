package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lissteron/deployer/app/git"
	"github.com/spf13/viper"
)

func main() {
	// Init config, logger, exit chan
	viper.AutomaticEnv()
	logger := log.New(os.Stdout, "", log.Lshortfile)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr: viper.GetString("HTTP_HOST"),
	}

	ctx, cancel := context.WithCancel(context.Background())

	makeRoute(logger)

	logger.Println("server start")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatalln("[error]", err)
		}
	}()

	<-exit

	go func() {
		time.Sleep(viper.GetDuration("EXIT_TIMEOUT"))
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logger.Println("[error]", err)
	}

	logger.Println("[info]", "server stopped")
}

func makeRoute(logger *log.Logger) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		git.WebhookRequest(logger, w, r)
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		// TODO
	})
}

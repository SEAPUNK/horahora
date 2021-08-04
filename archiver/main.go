package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/SEAPUNK/horahora/archiver/internal/config"
	"github.com/SEAPUNK/horahora/archiver/internal/grpcserver"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Could not get config. Err: %s", err)
	}

	// graceful signal handling
	ctx, close := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sigChan
		log.Errorf("Signal %s received. Canceling context", s)
		close()
	}()

	wg := sync.WaitGroup{}

	// grpc server
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := grpcserver.NewGRPCServer(ctx, cfg)
		if err != nil {
			log.Error(err)
		}
		log.Info("GRPC server exited")
	}()

	log.Info("Goroutines started, waiting")
	wg.Wait()
	log.Info("All goroutines have returned. Exiting...")
}

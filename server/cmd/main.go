// Package main is an entry point to the server app.
package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"gophkeeper/server/internal/logger"
	"gophkeeper/server/internal/runner"
)

func main() {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	server := grpc.NewServer()
	go runner.Run(server)

	<-termChan
	logger.Logger.Info("Server shutdown gracefully")
	time.Sleep(time.Second)
	server.Stop()
}

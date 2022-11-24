package main

import (
	"binginx.com/brush/cmd/api/routers"
	"binginx.com/brush/config"
	"binginx.com/brush/internal/clients"
	"binginx.com/brush/internal/logs"
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	config.MustInit()
	logs.Init()
	clients.MustInit()
	router := routers.Init()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GlobalConfig.ServerPort),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	logs.Logger.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logs.Logger.Fatal("Server forced to shutdown: ", err)
	}

	logs.Logger.Info("Server exiting")
}

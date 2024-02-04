// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Naoji1984/go-rana8/config"
	"github.com/Naoji1984/go-rana8/listener"
	"github.com/labstack/echo"
)

func main() {
	config.LoadConfig()
	e := echo.New()
	listener.Init(e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", config.Config.AppPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

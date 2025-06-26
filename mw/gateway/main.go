package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prabhakk-mw/learngo/mw/common/defs"
	_ "github.com/prabhakk-mw/learngo/mw/gateway/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	httpPort = ":8081"
)

//	@title			MathWorks API Gateway
//	@version		1.0
//	@description	Control and communicate with MATLAB through API Endpoints
//	@termsOfService	https://www.mathworks.com/license/mll/license.txt

//	@contact.name	The MathWorks Inc.
//	@contact.url	https://www.mathworks.com/support
//	@contact.email	support@mathworks.com

//	@license.name	MathWorks Cloud Reference Architecture License
//	@license.url	https://github.com/mathworks/jupyter-matlab-proxy/blob/main/LICENSE.md

// @host		localhost:8081
// @BasePath	/
func main() {
	log.Printf("Use : http://localhost%s/capitalize?payload=yourtext to capitalize text\n", httpPort)

	mainCtx, mainCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer mainCancel()

	handler := &defs.Handlers{RootCtx: mainCtx}

	// Create a new server
	mux := http.NewServeMux()

	mux.HandleFunc("/capitalize", handler.CapitalizeHandler)
	mux.HandleFunc("/static-capitalize", handler.StaticCapitalizeHandler)
	mux.HandleFunc("/exit", func(_ http.ResponseWriter, _ *http.Request) { mainCancel() })
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	srv := &http.Server{
		Addr:    httpPort,
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for a shutdown signal
	<-mainCtx.Done()
	log.Printf("[[%v]] signal received... shutting down (5 seconds)", mainCtx.Err().Error())

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	log.Println("gateway shutdown complete.")
}

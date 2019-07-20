package proxy

import (
	"context"
	"fmt"
	"github.com/forsam-education/kerberos/utils"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

// ConnectionCount represents the count of active connections on the proxy server.
var ConnectionCount = 0

// StartServer launches the reverse proxy main http server.
func StartServer(ctx context.Context, group *sync.WaitGroup) {
	host := fmt.Sprintf("%s:%d", viper.GetString(utils.ProxyServerHost), viper.GetInt(utils.ProxyServerPort))

	// Catch interrupt signal in channel.
	interruptSignalChannel := make(chan os.Signal, 1)
	signal.Notify(interruptSignalChannel, os.Interrupt)

	// Initiate routes.
	router := mux.NewRouter()
	for _, middleware := range middlewares {
		router.Use(middleware)
	}
	for _, service := range services {
		router.HandleFunc(service.Path, func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, service.TargetURL, http.StatusMovedPermanently)
		}).Methods(strings.Split(service.Methods, ",")...)
	}

	server := &http.Server{
		Handler:      router,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		utils.Logger.Info(fmt.Sprintf("Proxy server listening on http://%s...", host), nil)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.StdError(err, nil)
			os.Exit(1)
		}
	}()

	// Wait for interruption signal.
	<-interruptSignalChannel

	// Shutdown server.
	if err := server.Shutdown(ctx); err != nil {
		utils.LogAndForceExit(err)
	}

	utils.Logger.Info("Proxy server stopped.", nil)
	group.Done()
}

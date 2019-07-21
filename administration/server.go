package administration

import (
	"context"
	"fmt"
	"github.com/forsam-education/kerberos/utils"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

// StartServer starts the administration HTTP server.
func StartServer(ctx context.Context, group *sync.WaitGroup) {
	host := fmt.Sprintf("%s:%d", viper.GetString(utils.AdministrationServerHost), viper.GetInt(utils.AdministrationServerPort))

	// Catch interrupt signal in channel.
	interruptSignalChannel := make(chan os.Signal, 1)
	signal.Notify(interruptSignalChannel, os.Interrupt)

	// Initiate routes.
	router := mux.NewRouter()
	for _, middleware := range globalMiddlewares {
		router.Use(middleware)
	}
	for _, route := range routes {
		router.Handle(route.Path, route.Handler).Methods(route.Methods...)
	}

	server := &http.Server{
		Handler:      router,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		utils.Logger.Info(fmt.Sprintf("Administration server listening on http://%s...", host), nil)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.LogAndForceExit(err)
		}
	}()

	// Wait for interruption signal.
	<-interruptSignalChannel

	// Shutdown server.
	if err := server.Shutdown(ctx); err != nil {
		utils.LogAndForceExit(err)
	}

	utils.Logger.Info("Administration server stopped.", nil)
	group.Done()
}

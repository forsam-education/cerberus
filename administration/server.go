package administration

import (
	"context"
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/gobuffalo/packr/v2"
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

	// Generate administration frontend fileserver.
	frontApplicationBox := packr.New("frontApplication", "../web/dist")
	frontAppServer := spaHandler{indexPath: "index.html", staticPath: frontApplicationBox.Path}

	// Catch interrupt signal in channel.
	interruptSignalChannel := make(chan os.Signal, 1)
	signal.Notify(interruptSignalChannel, os.Interrupt)

	// Initiate routes.
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	router.PathPrefix("/").Handler(frontAppServer)
	for _, middleware := range globalMiddlewares {
		apiRouter.Use(middleware)
	}
	for _, route := range apiRoutes {
		apiRouter.Handle(route.Path, route.Handler).Methods(route.Methods...)
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

package api

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

// StartServer starts the administration API HTTP server.
func StartServer(group *sync.WaitGroup) {
	host := fmt.Sprintf("%s:%d", viper.GetString(utils.KerberosHost), viper.GetInt(utils.KerberosAPIPort))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	router := mux.NewRouter()

	for _, route := range getRoutes() {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Methods...)
	}

	server := &http.Server{
		Handler:      router,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		utils.Logger.Info(fmt.Sprintf("API server listening on http://%s...", host), nil)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.StdError(err, nil)
			os.Exit(1)
		}
	}()

	<-stop

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if err := server.Shutdown(ctx); err != nil {
		utils.Logger.StdError(err, nil)
		os.Exit(1)
	}

	utils.Logger.Info("API server stopped.", nil)
	group.Done()
}

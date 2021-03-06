package administration

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// StartServer starts the administration HTTP server.
func StartServer(group *sync.WaitGroup) {
	host := fmt.Sprintf("%s:%d", viper.GetString(utils.AdministrationServerHost), viper.GetInt(utils.AdministrationServerPort))

	// Catch interrupt signal in channel.
	signalCatcher := make(chan os.Signal, 1)
	signal.Notify(signalCatcher, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	// Initiate routes.
	router := initRouter()

	ln, err := reuseport.Listen("tcp4", host)
	if err != nil {
		utils.Logger.StdError(err, nil)
		os.Exit(1)
	}

	server := fasthttp.Server{
		Name:         "Admin",
		Handler:      router.Handler,
		LogAllErrors: true,
	}

	go func() {
		utils.Logger.Info(fmt.Sprintf("Administration server listening on http://%s...", host), nil)

		if err := server.Serve(ln); err != nil {
			utils.Logger.StdErrorCritical(err, nil)
			os.Exit(1)
		}
	}()

	// Wait for interruption signal.
	<-signalCatcher

	// Shutdown server.
	if err := server.Shutdown(); err != nil {
		utils.Logger.Critical(err.Error(), nil)
	}

	utils.Logger.Info("Administration server stopped.", nil)
	group.Done()
}

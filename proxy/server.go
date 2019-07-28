package proxy

import (
	"context"
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// ConnectionCount represents the count of active connections on the proxy server.
var ConnectionCount = 0

// StartServer launches the reverse proxy main http server.
func StartServer(_ context.Context, group *sync.WaitGroup) {
	host := fmt.Sprintf("%s:%d", viper.GetString(utils.ProxyServerHost), viper.GetInt(utils.ProxyServerPort))

	// Catch interrupt signal in channel.
	signalCatcher := make(chan os.Signal, 1)
	signal.Notify(signalCatcher, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	ln, err := reuseport.Listen("tcp4", host)
	if err != nil {
		utils.Logger.StdError(err, nil)
		os.Exit(1)
	}

	server := fasthttp.Server{
		Name:               "Cerberus",
		Handler:            proxify,
		LogAllErrors:       true,
		ReadTimeout:        viper.GetDuration(utils.ProxyServerReadTimeout) * time.Second,
		TCPKeepalive:       viper.GetBool(utils.ProxyServerKeepAliveState),
		TCPKeepalivePeriod: viper.GetDuration(utils.ProxyServerKeepAlivePeriod) * time.Second,
	}

	go func() {
		utils.Logger.Info(fmt.Sprintf("Proxy server listening on http://%s...", host), nil)

		if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
			utils.Logger.StdError(err, nil)
			os.Exit(1)
		}
	}()

	// Wait for interruption signal.
	<-signalCatcher

	// Shutdown server.
	if err := server.Shutdown(); err != nil {
		utils.Logger.Critical(err.Error(), nil)
	}

	utils.Logger.Info("Proxy server stopped.", nil)
	group.Done()
}

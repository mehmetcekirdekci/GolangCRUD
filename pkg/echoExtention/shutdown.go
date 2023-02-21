package echoextention

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

func Shutdown(instance *echo.Echo, timeout time.Duration) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := instance.Shutdown(ctx)
	if err != nil {
		println(err)
	} else {
		println("server was shut down succesfuly.")
	}
}

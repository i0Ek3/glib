package graceful

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func Shutdown(signals ...os.Signal) {
	quit := make(chan os.Signal)
	signal.Notify(quit, signals...)
	<-quit
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Println("timeout")
	}
	log.Println("server exiting...")
}

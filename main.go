package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sig := UseSignalHook()

	ctx, cancel := context.WithCancel(context.Background())

	period := time.Second
	go TaskRunner(ctx, period)

	DoShutDownProcess(sig, cancel)

}

func UseSignalHook() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	return sig
}

func DoShutDownProcess(sig <-chan os.Signal, cancel context.CancelFunc) {

	<-sig
	cancel()
	fmt.Println("main func was stopped by Signal")
}

func TaskRunner(ctx context.Context, period time.Duration) {
	timer := time.NewTicker(period)
	func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done")
				return
			case <-timer.C:
				MainProcess()
			}
		}
	}()
}

func MainProcess() {
	fmt.Println("MainProcessess")
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/toretto460/pooly"
)

func waitSignal(done func()) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done()
	}()
}

func main() {
	ctx, done := context.WithCancel(context.Background())
	waitSignal(done)

	p := pooly.New(ctx, 4)

	p.RunFunc(func() {
		time.Sleep(time.Second * 2)
		fmt.Print("Job 1 Done\n")
	})

	p.RunFunc(func() {
		time.Sleep(time.Second * 10)
		fmt.Print("Job 2 Done\n")
	})

	p.Wait()
}

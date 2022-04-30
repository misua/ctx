package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background() //parent

	//withtimeout example = WithTimeout
	//ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	//cancel() // if you call this directly you will get context cancelled
	//defer cancel()
	ctx, cancel := context.WithCancel(ctx)

	go func() { // you can define any event u want,sample stdin event only
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	//time.AfterFunc(time.Second, cancel)
	SleepAndTalk(ctx, 5*time.Second, "hello")
}

func SleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Print(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())

	}

}

package main

import (
	"context"
	"fmt"
	"time"
)

func eterno() {
	for {
		time.Sleep(time.Second)
		fmt.Println(".")
	}
}

func main() {

	ctx := context.Background()

	ctx, _ = context.WithTimeout(ctx, time.Second*5)
	// go eterno()
	<-ctx.Done()
	fmt.Println(ctx.Err().Error())
}

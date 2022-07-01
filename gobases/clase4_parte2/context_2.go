package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "saludo", "hola digital house!!")
	saludoWrapper(ctx)
}

func saludoWrapper(ctx context.Context) {
	fmt.Println(ctx.Value("saludo"))
	saludo(ctx)
}

func saludo(ctx context.Context) {
	fmt.Println(ctx.Value("saludo"))
}

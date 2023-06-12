package main

import (
	"context"
	"fmt"
	"time"
)

// 1. context.Background() - на самом высоком уровне. В функции main
// 2. context.TODO - когда не уверены, какой контекст использовать
// 3. context.WithValue - стоит использовать как можно реже и передавать только
// необязательные параметры
// 4. ctx всегда передается первым аргументом в функцию

func main() {
	ctx := context.Background()
	// ctx, _ = context.WithTimeout(ctx, time.Second*3)
	// ctx = context.WithValue(ctx, "id", 1)

	// go func() {
	// 	time.Sleep(time.Millisecond * 100)
	// 	cancel()
	// }()

	parse(ctx)
}

func parse(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	// id := ctx.Value("id")
	// fmt.Println(id.(int))
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceded")
			return
		}
	}
}

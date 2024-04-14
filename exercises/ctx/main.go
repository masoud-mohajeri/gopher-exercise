package ctx

import (
	"context"
	"fmt"
	"time"
)

func Main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// defer func() {
	// 	fmt.Println("cancel in defer")
	// 	// cancel()
	// }()

	go performTask(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("WithTimeout")
	case <-time.After(1 * time.Second):
		cancel()
		fmt.Println("canceled")

	}
}

func performTask(ctx context.Context) {
	select {
	case <-time.After(4 * time.Second):
		fmt.Println("Task completed successfully")
	}
}

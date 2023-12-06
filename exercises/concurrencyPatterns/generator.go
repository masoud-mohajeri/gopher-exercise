package concurrencyPatterns

import (
	"fmt"
	"math/rand"
	"time"
)

// put data in single chanel from multiple goroutines
// and use in multiple goroutines
func fanOut(message string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			out <- message
		}
	}()

	return out
}

func Generator() {
	c := fanOut("Hi")

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	time.Sleep((time.Duration(rand.Intn(1e8))))

	fmt.Println("enough talking")
}

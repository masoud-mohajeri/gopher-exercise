package concurrencyPatterns

import (
	"fmt"
)

// From multiple goroutines get data
// and put them in a single chanel and use it in one goroutine
func fanIn(input1, input2 string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- input1
		}
	}()

	go func() {
		for {
			c <- input2
		}

	}()

	return c
}

func Multiplexed() {
	c := fanIn("joe", "clair")

	for i := 0; i < 22; i++ {
		fmt.Printf("%s: %d \n", <-c, i)
	}

	fmt.Println("enough talking")
}

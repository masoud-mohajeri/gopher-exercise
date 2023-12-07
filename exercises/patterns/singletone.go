package patterns

import (
	"fmt"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type Single struct {
}

var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			singleInstance = &Single{}
			fmt.Println("instance created successfully")

		} else {
			fmt.Println("instance already exists")
		}
	} else {
		fmt.Println("instance already exists")

	}
	return singleInstance
}

func SingletonPattern() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

	time.Sleep(time.Second)
}

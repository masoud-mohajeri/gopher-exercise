package schaduler

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// simple outbox scenario
type outbox struct {
	messagesCount int
}

func (o *outbox) getOutbox() {
	o.messagesCount = int(math.Floor(rand.Float64() * 10))
	fmt.Println("number of new messages in outbox table", o.messagesCount) // SELECT * FROM outbox WHERE NOT is_sent

}

func (o outbox) setOutBox() {
	fmt.Println("get message from db and write it in kafka") //set message in kafka then: UPDATE outbox SET is_sent=true WHERE id === ?
	<-time.After(500 * time.Millisecond)
}

func (o outbox) execute() {
	o.getOutbox()

	var wg sync.WaitGroup
	wg.Add(o.messagesCount)

	for i := 0; i < o.messagesCount; i++ {
		go func() {
			o.setOutBox()
			wg.Done()
		}()
	}

	wg.Wait()
}

// The planner code
type plannerAction interface{ execute() }

func planner(action plannerAction, d time.Duration) {
	for {
		action.execute()
		<-time.After(d)
	}
}

func Main() {
	out := outbox{}
	planner(out, time.Second)
}

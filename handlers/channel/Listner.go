package channel

import (
	"fmt"
	"sync"
	"time"
)

func Listener(wg *sync.WaitGroup) {
	for msg := range LogChannel {
		message := fmt.Sprintf("[%s] %s /books -%s", msg.Time.Format(time.RFC3339), msg.Method, msg.Data)
		fmt.Println(message)

	}
	wg.Done()
}
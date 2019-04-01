package cancel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	cancelChan := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "cancelled")
		}(i, cancelChan)
	}
	time.Sleep(time.Second * 1)
	closeChannel(cancelChan)
}

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel(cancelChan chan struct{}) {
	cancelChan<- struct{}{}
}

func closeChannel(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

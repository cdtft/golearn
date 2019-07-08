package goroutine

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"testing"
)

var (
	counter   int
	waitGroup sync.WaitGroup
)

func TestIncCounter(t *testing.T) {
	waitGroup.Add(2)

	go incCounter(2)
	go incCounter(1)

	waitGroup.Wait()
	log.Println("final count:", counter)
}

func incCounter(id int) {
	defer waitGroup.Done()
	for count := 0; count < 2; count ++ {
		runtime.Gosched()
		counter ++
	}
}

func TestGomaxprocs(t *testing.T) {
	waitGroup.Add(2)
	runtime.GOMAXPROCS(1)
	go func() {
		defer waitGroup.Done()
		for i := 'a'; i < 'a'+26; i++ {
			fmt.Printf(" %c", i)
			runtime.Gosched()
		}
	}()

	go func() {
		defer waitGroup.Done()
		for j := 'A'; j < 'A'+26; j++ {
			fmt.Printf(" %c", j)
			runtime.Gosched()
		}
	}()
	waitGroup.Wait()
}

func TestDo(t *testing.T) {
	go do("hello")
	do("world")
}

func do(s string) {
	log.Printf("do:%s", s)
}

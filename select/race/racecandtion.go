package race

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg sync.WaitGroup
)

func MainRaceCondition() {
	wg.Add(2)
	go incCounter()
	go incCounter()
	wg.Wait()
	fmt.Println(counter)
}

func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}


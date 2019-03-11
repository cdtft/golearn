package race

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg sync.WaitGroup
	mutex sync.Mutex
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

func MainSyncIncCounter() {
	wg.Add(2)
	go syncIncCounter()
	go syncIncCounter()
	wg.Wait()
	fmt.Println(counter)
}

func syncIncCounter() {
	defer wg.Done()
	for count := 0; count < 3; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}


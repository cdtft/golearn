package goroutine

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
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

var shutdown int64

func TestLoadAndStoreInt(t *testing.T) {
	waitGroup.Add(2)

	go doWork("张三")
	go doWork("李四")

	time.Sleep(600 * time.Microsecond)
	atomic.StoreInt64(&shutdown, 1)
	waitGroup.Wait()

}

func doWork(name string) {
	defer waitGroup.Done()
	for {
		fmt.Printf("Doing %s work \n", name)
		time.Sleep(100 * time.Microsecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			break
		}
	}
}

func TestUseChain(t *testing.T) {
	var nums []int
	nums = append(nums, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	infoChan := make(chan int, 3)
	numsLen := len(nums)
	boundary := numsLen % 5
	for index := 0; index <= numsLen; index += 5 {

		if (numsLen > boundary && numsLen-index < 5) || numsLen <= boundary {
			go calculateNum(nums[index:], infoChan)
			break
		}
		//每5个就开一个协程去处理
		//if (index+1)%5 == 0 {
		go calculateNum(nums[index:index+5], infoChan)
		//}
	}

	var resultNums []int
	for {
		image := <-infoChan
		resultNums = append(resultNums, image)

		if len(resultNums) == len(nums) {
			close(infoChan)
			fmt.Print(resultNums)
			return
		}
	}
}

func calculateNum(myNums []int, c chan int) {
	for _, i := range myNums {
		//randomNum := rand.Int31n(5)
		//time.Sleep(time.Duration(randomNum) * time.Second)
		fmt.Println(i)
		c <- i * 3
	}
}

func TestDivision(t *testing.T) {

	fmt.Println(5 % 2)
}

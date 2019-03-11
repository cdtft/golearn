package player

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//无缓冲通道

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court

		//通道关闭
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}
		//生成随机数判断是否丢球
		randInt := rand.Intn(100)
		if randInt%16 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s hite ball %d times\n", name, ball)
		ball++
		court <- ball
	}
}

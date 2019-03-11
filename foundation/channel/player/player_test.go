package player

import "testing"

func TestPlayer(t *testing.T) {
	court := make(chan int)

	wg.Add(2)

	go Player("wang", court)

	go Player("cheng", court)

	court <- -1
	wg.Wait()
}

package runner

import (
	"fmt"
	"testing"
)

func TestRunner_Start(t *testing.T) {
	runner := New(100)

	runner.Add(func(i int) {
		fmt.Println(i)
	}, func(i int) {
		fmt.Println(i)
	})

	err := runner.Start()
	t.Log(err.Error())
}
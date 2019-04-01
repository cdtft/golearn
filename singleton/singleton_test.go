package singleton

import (
	"fmt"
	"testing"
)

func TestGetSingleton(t *testing.T) {

	for i := 0; i < 10; i++ {
		singleton := GetSingleton()

		fmt.Println(&singleton)
	}
}

package sqrt

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {

	var x float32 = 2
	var y float32 = 1
	for i := 0; i < 20; i++ {
		y = sqrtIter(x, y)
	}
	fmt.Println(y)
}

func sqrtIter(x float32, y float32) float32 {
	z := x / y
	return (z + y) / 2
}

package sqrt

import (
	"fmt"
	"testing"
)

//逼近法计算开平方
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

func TestFactorial(t *testing.T) {
	fmt.Println(factorial(5))
	fmt.Println(factorialIter1(5))
}

func factorial(n int) int {
	return factorialIter(1, 1, n)
}

func factorialIter(product int, counter int, maxCount int) int {
	if counter > maxCount {
		return product
	}
	return factorialIter(product * counter, counter + 1, maxCount)
}

func factorialIter1(n int) int {
	if n == 1 {
		return 1
	}
	return factorialIter1(n-1) * n
}


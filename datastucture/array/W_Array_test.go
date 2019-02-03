package array

import "testing"

func TestWArray_Get(t *testing.T) {
	myArray := NewWArray(10)
	_ = myArray.Set(0, 1)

	value, _ := myArray.Get(0)

	t.Log(value)
}
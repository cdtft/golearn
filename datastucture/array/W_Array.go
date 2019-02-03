package array

import (
	"errors"
	"fmt"
)

type WArray struct {
	data     []int
	length   uint
	capacity uint
}

//创建WArrayList
func NewWArray(capacity uint) *WArray {
	if capacity == 0 {
		return nil
	}
	return &WArray{
		data:     make([]int, capacity, capacity),
		length:   0,
		capacity: capacity,
	}
}

//长度
func (wArray *WArray) Len() uint {
	return wArray.length
}

//对应下标设置值
func (wArray *WArray) Set(index int, value int) error {
	wArray.data[index] = value
	wArray.length++
	return nil
}

func (wArray *WArray) Get(index int) (int, error) {
	if wArray.rangeCheck(index) {
		return 0, errors.New("array out of bound")
	}
	return wArray.data[index], nil
}

func (wArray *WArray) rangeCheck(index int) bool {
	if uint(index) >= wArray.capacity {
		return true
	}
	return false
}

//打印
func (wArray *WArray) PrintToString() {
	for i := uint(0); i < wArray.capacity; i++ {
		fmt.Print(wArray.data[i])
		if i < wArray.capacity-1 {
			fmt.Print(",")
		} else {
			fmt.Println()
		}
	}
}


package main

import (
	"fmt"
	"container/heap"
)


// we need to implement heap.Interface for our data structure

type MyMinHeap []int


// Len
func (mh MyMinHeap) Len() int {
	return len(mh)
}

// Less
func (mh MyMinHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}


// Swap
func (mh MyMinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}


// Pop - remove item at len-1
func (mh* MyMinHeap) Pop() interface{} {
	old := *mh				
	l := len(old)
	val := old[l-1]
	*mh = old[0:l-1]
	
	return val
}


// Push - add item to back
func (mh* MyMinHeap) Push(item interface{}) {
	*mh = append(*mh, item.(int))
}


// test some functionality of go heap

func main() {
	mh := &MyMinHeap{2, 15, 6}
	heap.Init(mh)
	
	heap.Push(mh, 1)
	heap.Push(mh, 18)
	
	val := heap.Pop(mh)
	fmt.Println(val)
	
	for mh.Len() > 0 {
		val := heap.Pop(mh)
		fmt.Println(val)
	}
}



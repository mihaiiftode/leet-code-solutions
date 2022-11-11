package main

import (
	"container/heap"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.AddNum(12)
	obj.AddNum(10)
	obj.AddNum(13)
	obj.AddNum(11)
	obj.AddNum(5)
	obj.AddNum(15)
	obj.AddNum(1)
	obj.AddNum(11)
	obj.AddNum(6)
	obj.AddNum(17)
	obj.AddNum(14)
	obj.AddNum(8)
	obj.AddNum(17)
	obj.AddNum(6)
	obj.AddNum(4)
	obj.AddNum(16)
	fmt.Println(obj.FindMedian())
	obj.AddNum(8)
	obj.AddNum(10)
	obj.AddNum(2)
	obj.AddNum(12)
	fmt.Println(obj.FindMedian())

	obj.AddNum(0)

}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Zero() int          { return h[0] }
func (h *MinHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Zero() int          { return h[0] }

func (h *MaxHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {

	minH := &MinHeap{}
	maxH := &MaxHeap{}
	heap.Init(minH)
	heap.Init(maxH)

	md := MedianFinder{
		minHeap: minH,
		maxHeap: maxH,
	}

	return md
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxHeap, num)
	heap.Push(this.minHeap, heap.Pop(this.maxHeap))

	if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap.Zero())
	} else {
		return float64(this.maxHeap.Zero()+this.minHeap.Zero()) / 2.0
	}
}

// func Constructor() MedianFinder {
// 	return MedianFinder{}
// }

// func (this *MedianFinder) AddNum(num int) {
// 	count := len(this.array)
// 	if count == 0 {
// 		this.array = append(this.array, num)
// 		return
// 	}
// 	if count != 0 && num >= this.array[len(this.array)-1] {
// 		this.array = append(this.array, num)
// 		return
// 	} else if count != 0 && num <= this.array[0] {
// 		this.array = append([]int{num}, this.array...)
// 		return
// 	}

// 	insert := sort.SearchInts(this.array, num)

// 	this.array = append(this.array[:insert], append([]int{num}, this.array[insert:]...)...)

// }

// func (this *MedianFinder) FindMedian() float64 {
// 	middle := len(this.array) / 2

// 	if len(this.array)%2 == 1 {
// 		return float64(this.array[middle])
// 	} else {
// 		return float64(this.array[middle]+this.array[middle-1]) / 2.0
// 	}
// }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

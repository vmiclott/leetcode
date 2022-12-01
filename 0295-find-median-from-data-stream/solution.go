package findmedianfromdatastream

import "container/heap"

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	ret := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return ret
}

type MinHeap struct {
	PriorityQueue
}

func (heap *MinHeap) Less(i, j int) bool {
	return heap.PriorityQueue[i] < heap.PriorityQueue[j]
}

type MaxHeap struct {
	PriorityQueue
}

func (heap *MaxHeap) Less(i, j int) bool {
	return heap.PriorityQueue[i] > heap.PriorityQueue[j]
}

type MedianFinder struct {
	smallNumbers *MaxHeap
	largeNumbers *MinHeap
	length       int
}

func Constructor() MedianFinder {
	return MedianFinder{&MaxHeap{}, &MinHeap{}, 0}
}

func (this *MedianFinder) AddNum(num int) {
	if this.length%2 == 0 {
		if this.length == 0 || this.largeNumbers.PriorityQueue[0] >= num {
			heap.Push(this.smallNumbers, num)
		} else {
			heap.Push(this.smallNumbers, heap.Pop(this.largeNumbers))
			heap.Push(this.largeNumbers, num)
		}
	} else {
		if this.smallNumbers.PriorityQueue[0] > num {
			heap.Push(this.largeNumbers, heap.Pop(this.smallNumbers))
			heap.Push(this.smallNumbers, num)
		} else {
			heap.Push(this.largeNumbers, num)
		}
	}
	this.length = this.length + 1
}

func (this *MedianFinder) FindMedian() float64 {
	if this.length%2 == 0 {
		return float64(this.smallNumbers.PriorityQueue[0]+this.largeNumbers.PriorityQueue[0]) / 2.0
	}
	return float64(this.smallNumbers.PriorityQueue[0])
}

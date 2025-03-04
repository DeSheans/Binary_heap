package heap

import (
	"cmp"
)

// Heap is a data structure for working with a min binary heap.
type Heap[T any] struct {
	heap    []T
	compare func(x, y T) int
}

// NewHeap creates and returns a new Heap instance initialized with the given slice.
func NewHeap[T cmp.Ordered](s []T) *Heap[T] {
	h := &Heap[T]{
		heap:    s,
		compare: cmp.Compare[T],
	}
	buildHeap(h)
	return h
}

// NewHeapWithComparer ...
func NewHeapWithComparer[T any](s []T, comparer func(x, y T) int) *Heap[T] {
	h := &Heap[T]{
		heap:    s,
		compare: comparer,
	}
	buildHeap(h)
	return h
}

// Insert adds a new element to the heap.
func (h *Heap[T]) Insert(val T) {
	h.heap = append(h.heap, val)
	h.siftUp(len(h.heap) - 1)
}

// Remove deletes the specified value from the heap if it exists. Returns true if the value was found and removed, otherwise false.
func (h *Heap[T]) Remove(val T) bool {
	for i, v := range h.heap {
		if h.compare(v, val) == 0 {
			size := len(h.heap) - 1
			h.heap[i], h.heap[size] = h.heap[size], h.heap[i]
			size--
			h.heap = h.heap[:size]
			h.heapify(i, size)
			return true
		}
	}
	return false
}

// Pop removes and returns the minimum element (root) from the heap. It panics if the heap is empty.
func (h *Heap[T]) Pop() (ok bool, result T) {
	if len(h.heap) == 0 {
		return false, result
	}
	size := len(h.heap) - 1
	result = h.heap[0]
	h.heap[0], h.heap[size] = h.heap[size], result
	h.heap = h.heap[:size]
	size--
	h.heapify(0, size)
	return true, result
}

// Peek returns the minimum element (root) from the heap without removing it. It panics if the heap is empty.
func (h *Heap[T]) Peek() (ok bool, result T) {
	if len(h.heap) == 0 {
		return false, result
	}
	result = h.heap[0]
	return true, result
}

// Sort returns a sorted slice of elements in ascending order using the heap. After sorting, the original heap is restored to maintain the heap property for future operations.
func (h *Heap[T]) Sort() []T {
	result := make([]T, len(h.heap))
	size := len(h.heap) - 1
	for i := range h.heap {
		result[i] = h.heap[0]
		h.heap[0], h.heap[size] = h.heap[size], h.heap[0]
		size--
		h.heapify(0, size)
	}
	buildHeap(h)
	return result
}

func buildHeap[T any](heap *Heap[T]) {
	size := len(heap.heap) - 1

	for i := (len(heap.heap) >> 1) - 1; i >= 0; i-- {
		heap.heapify(i, size)
	}
}

func (h *Heap[T]) heapify(index int, size int) {
	leftInd := (index << 1) + 1
	rightInd := leftInd + 1
	minInd := index
	if leftInd <= size && h.compare(h.heap[leftInd], h.heap[minInd]) == -1 {
		minInd = leftInd
	}
	if rightInd <= size && h.compare(h.heap[rightInd], h.heap[minInd]) == -1 {
		minInd = rightInd
	}
	if minInd != index {
		h.heap[index], h.heap[minInd] = h.heap[minInd], h.heap[index]
		h.heapify(minInd, size)
	}
}

func (h *Heap[T]) siftUp(ind int) {
	for ind != 0 && h.compare(h.heap[ind], h.heap[(ind-1)>>1]) == -1 {
		h.heap[ind], h.heap[(ind-1)>>1] = h.heap[(ind-1)>>1], h.heap[ind]
		ind = (ind - 1) >> 1
	}
}

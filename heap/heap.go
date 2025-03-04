package heap

import "cmp"

// Heap is a data structure for working with a min binary heap.
type Heap[T cmp.Ordered] struct {
	heap []T
}

// NewHeap creates and returns a new Heap instance initialized with the given slice.
func NewHeap[T cmp.Ordered](s []T) *Heap[T] {
	buildHeap(s)
	return &Heap[T]{
		heap: s,
	}
}

// Insert adds a new element to the heap.
func (h *Heap[T]) Insert(val T) {
	h.heap = append(h.heap, val)
	siftUp(h.heap, len(h.heap)-1)
}

// Remove deletes the specified value from the heap if it exists. Returns true if the value was found and removed, otherwise false.
func (h *Heap[T]) Remove(val T) bool {
	for i, v := range h.heap {
		if v == val {
			size := len(h.heap) - 1
			h.heap[i], h.heap[size] = h.heap[size], h.heap[i]
			size--
			h.heap = h.heap[:size]
			heapify(h.heap, i, size)
			return true
		}
	}
	return false
}

// Pop removes and returns the minimum element (root) from the heap. It panics if the heap is empty.
func (h *Heap[T]) Pop() T {
	if len(h.heap) == 0 {
		panic("Heap is empty.")
	}
	size := len(h.heap) - 1
	result := h.heap[0]
	h.heap[0], h.heap[size] = h.heap[size], result
	h.heap = h.heap[:size]
	size--
	heapify(h.heap, 0, size)
	return result
}

// Peek returns the minimum element (root) from the heap without removing it. It panics if the heap is empty.
func (h *Heap[T]) Peek() T {
	if len(h.heap) == 0 {
		panic("Heap is empty.")
	}
	return h.heap[0]
}

// Sort returns a sorted slice of elements in ascending order using the heap. After sorting, the original heap is restored to maintain the heap property for future operations.
func (h *Heap[T]) Sort() []T {
	result := make([]T, len(h.heap))
	size := len(h.heap) - 1
	for i := range h.heap {
		result[i] = h.heap[0]
		h.heap[0], h.heap[size] = h.heap[size], h.heap[0]
		size--
		heapify(h.heap, 0, size)
	}
	buildHeap(h.heap)
	return result
}

func buildHeap[T cmp.Ordered](s []T) {
	size := len(s) - 1
	for i := (len(s) >> 1) - 1; i >= 0; i-- {
		heapify(s, i, size)
	}
}

func heapify[T cmp.Ordered](heap []T, index int, size int) {
	leftInd := (index << 1) + 1
	rightInd := leftInd + 1
	minInd := index
	if leftInd <= size && heap[leftInd] < heap[minInd] {
		minInd = leftInd
	}
	if rightInd <= size && heap[rightInd] < heap[minInd] {
		minInd = rightInd
	}
	if minInd != index {
		heap[index], heap[minInd] = heap[minInd], heap[index]
		heapify(heap, minInd, size)
	}
}

func siftUp[T cmp.Ordered](heap []T, ind int) {
	for ind != 0 && heap[ind] < heap[(ind-1)>>1] {
		heap[ind], heap[(ind-1)>>1] = heap[(ind-1)>>1], heap[ind]
		ind = (ind - 1) >> 1
	}
}

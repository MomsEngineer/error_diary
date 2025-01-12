package main

import (
	"container/heap"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := &IntHeap{}
	heap.Init(h)

	t.Run("EmptyHeap", func(t *testing.T) {
		if h.Len() != 0 {
			t.Errorf("Expected heap length to be 0, got %d", h.Len())
		}
	})

	t.Run("PushAndPeek", func(t *testing.T) {
		heap.Push(h, 10)
		if (*h)[0] != 10 {
			t.Errorf("Expected minimum element to be 10, got %d", (*h)[0])
		}

		heap.Push(h, 5)
		if (*h)[0] != 5 {
			t.Errorf("Expected minimum element to be 5, got %d", (*h)[0])
		}

		heap.Push(h, 20)
		if (*h)[0] != 5 {
			t.Errorf("Expected minimum element to be 5, got %d", (*h)[0])
		}
	})

	t.Run("ExtractMin", func(t *testing.T) {
		if heap.Pop(h).(int) != 5 {
			t.Errorf("Expected extracted element to be 5")
		}

		if heap.Pop(h).(int) != 10 {
			t.Errorf("Expected extracted element to be 10")
		}

		if heap.Pop(h).(int) != 20 {
			t.Errorf("Expected extracted element to be 20")
		}

		if h.Len() != 0 {
			t.Errorf("Expected heap length to be 0 after extractions, got %d", h.Len())
		}
	})

	t.Run("HeapProperty", func(t *testing.T) {
		values := []int{3, 1, 6, 5, 2, 4}
		for _, v := range values {
			heap.Push(h, v)
		}

		expected := []int{1, 2, 3, 4, 5, 6}
		for _, exp := range expected {
			if heap.Pop(h).(int) != exp {
				t.Errorf("Heap property violated, expected %d", exp)
			}
		}
	})
}

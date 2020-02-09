package infra

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHeap(t *testing.T) {

	Convey("TestNewHeap1", t, func() {
		Convey("构建小顶堆", func() {

			got := NewHeap1([]int{9, 3, 7, 6, 5, 1, 10, 2}, true)
			want := []int{1, 2, 7, 3, 5, 9, 10, 6}

			if !IntArrSeqCmp(got.tree, want) {
				t.Errorf("failed")
			}

		})

		Convey("构建大顶堆", func() {

			got := NewHeap1([]int{9, 4, 2, 3, 5, 1, 8, 6, 7, 0}, false)
			want := []int{9, 7, 8, 6, 5, 1, 2, 4, 3, 0}

			if !IntArrSeqCmp(got.tree, want) {
				t.Errorf("failed")
			}

		})

	})

	Convey("TestNewHeap2", t, func() {
		Convey("构建小顶堆", func() {

			h := NewHeap2(true)
			arr := []int{9, 3, 7, 6, 5, 1, 10, 2}
			for i := 0; i < len(arr); i++ {
				h.TopInsert(arr[i])
			}
			/*
				got := h.tree
				want := []int{1, 2, 7, 3, 5, 9, 10, 6}

				if !IntArrSeqCmp(got, want) {
					t.Errorf("failed")
				}*/

		})

		Convey("构建大顶堆", func() {

			h := NewHeap2(false)
			arr := []int{9, 4, 2, 3, 5, 1, 8, 6, 7, 0}
			for i := 0; i < len(arr); i++ {
				h.BtmInsert(arr[i])
			}
			/*
				got := h.tree
				want := []int{9, 7, 8, 6, 5, 1, 2, 4, 3, 0}

				if !IntArrSeqCmp(got, want) {
					t.Errorf("failed")
				}*/

		})

	})

	Convey("TestHeapSort", t, func() {
		Convey("小顶堆排序", func() {

			h := NewHeap2(true)
			arr := []int{9, 3, 7, 6, 5, 1, 10, 2}
			for i := 0; i < len(arr); i++ {
				h.TopInsert(arr[i])
			}
			got := h.HeapSort(true)
			want := []int{1, 2, 3, 5, 6, 7, 9, 10}

			if !IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("大顶堆排序", func() {

			h := NewHeap2(false)
			arr := []int{9, 4, 2, 3, 5, 1, 8, 6, 7, 0}
			for i := 0; i < len(arr); i++ {
				h.BtmInsert(arr[i])
			}
			got := h.HeapSort(false)
			want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

			if !IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})
}

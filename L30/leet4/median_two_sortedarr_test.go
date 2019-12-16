package leet4

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMedianTwoSortedArr(t *testing.T) {

	Convey("TestMedianTwoSortedArr", t, func() {
		Convey("总元素为基数情况", func() {
			arr1 := []int{3, 5, 10, 11, 17}
			arr2 := []int{9, 13, 20, 21, 23, 27}
			got := MedianTwoSortedArr(arr1, arr2)
			want := 13.0

			if got != want {
				t.Errorf("got %.1f, want %.1f", got, want)
			}
			//ShouldNotBeNil(err)
		})

		Convey("总元素为偶数情况", func() {
			arr1 := []int{2, 3, 5, 8}
			arr2 := []int{10, 12, 14, 16, 18, 20}
			got := MedianTwoSortedArr(arr1, arr2)
			want := 11.0

			if got != want {
				t.Errorf("got %.1f, want %.1f", got, want)
			}
			//ShouldNotBeNil(err)
		})

		Convey("偶数且原生就左右分开,且上长下短", func() {
			arr1 := []int{1, 2, 3}
			arr2 := []int{4, 5}
			got := MedianTwoSortedArr(arr1, arr2)
			want := 3.0

			if got != want {
				t.Errorf("got %.1f, want %.1f", got, want)
			}
			//ShouldNotBeNil(err)
		})

	})

}

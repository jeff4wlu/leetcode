package leet31

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNextPermutation(t *testing.T) {

	Convey("TestNextPermutation", t, func() {
		Convey("最小元素的有意义数组", func() {
			arr := []int{1, 2}
			NextPermutation(arr)
			want := []int{2, 1}

			for i := 0; i < len(want); i++ {
				if arr[i] != want[i] {
					t.Errorf("got is not equal with want")
				}
			}

		})

		Convey("最大值的下一个排列应该是最小值", func() {
			arr := []int{3, 2, 1}
			NextPermutation(arr)
			want := []int{1, 2, 3}

			for i := 0; i < len(want); i++ {
				if arr[i] != want[i] {
					t.Errorf("got is not equal with want")
				}
			}

		})

		Convey("常规用例", func() {
			arr := []int{1, 2, 7, 4, 3, 1}
			NextPermutation(arr)
			want := []int{1, 3, 1, 2, 4, 7}

			for i := 0; i < len(want); i++ {
				if arr[i] != want[i] {
					t.Errorf("got is not equal with want")
				}
			}

		})

	})

}

package leet1

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum(t *testing.T) {

	Convey("TestTwoSum", t, func() {
		Convey("数组中相加等于目标值的两个索引", func() {
			arr := []int{3, 6, 7, 9}
			got := TwoSum(arr, 10)
			want := []int{0, 2}

			if got[0] != want[0] && got[1] != want[1] {
				t.Errorf("got []int{%d, %d}, want []int{%d, %d}", got[0], got[1], want[0], want[1])
			}
			//ShouldNotBeNil(err)
		})

		Convey("数组中有两个相同值的情况", func() {
			arr := []int{3, 6, 6, 7, 5}
			got := TwoSum(arr, 12)
			want := []int{3, 4}

			if got[0] != want[0] && got[1] != want[1] {
				t.Errorf("got []int{%d, %d}, want []int{%d, %d}", got[0], got[1], want[0], want[1])
			}
		})

	})

}

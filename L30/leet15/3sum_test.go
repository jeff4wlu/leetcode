package leet15

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThreeSum(t *testing.T) {

	Convey("TestThreeSum", t, func() {
		Convey("求三和等于0的solution", func() {
			in := []int{-1, 0, 1, 2, -1, -4}
			got := ThreeSum(in)
			/*
				want := [][]int{
					{-1, 0, 1},
					{-1, -1, 2},
				}*/
			if got == nil {
				t.Errorf("nil")
			}
			/*
				if got[0] != want[0] && got[1] != want[1] {
					t.Errorf("got []int{%d, %d}, want []int{%d, %d}", got[0], got[1], want[0], want[1])
				}*/

		})

	})

}

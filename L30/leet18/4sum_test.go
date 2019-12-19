package leet18

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFourSum(t *testing.T) {

	Convey("TestFourSum", t, func() {
		Convey("求四和等于target的solution", func() {
			in := []int{1, 0, -1, 0, -2, 2}
			got := FourSum(in, 0)
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

package leet167

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum(t *testing.T) {

	Convey("TestTwoSum", t, func() {
		Convey("用例1", func() {

			got := TwoSum([]int{2, 7, 11, 15}, 9)
			want := []int{1, 2}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed.")
			}

		})

	})

	Convey("TestTwoSum2", t, func() {
		Convey("用例1", func() {

			got := TwoSum2([]int{2, 7, 11, 15}, 9)
			want := []int{1, 2}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed.")
			}

		})

	})

}

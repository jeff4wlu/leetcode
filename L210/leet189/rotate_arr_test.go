package leet189

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRotateArr(t *testing.T) {

	Convey("TestRotateArr", t, func() {
		Convey("用例1", func() {

			got := RotateArr([]int{1, 2, 3, 4, 5, 6, 7}, 3)
			want := []int{5, 6, 7, 1, 2, 3, 4}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := RotateArr([]int{-1, -100, 3, 99}, 2)
			want := []int{3, 99, -1, -100}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

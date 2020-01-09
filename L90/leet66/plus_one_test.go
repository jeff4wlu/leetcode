package leet66

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlusOne(t *testing.T) {

	Convey("TestPlusOne", t, func() {
		Convey("用例1", func() {

			got := PlusOne([]int{1, 2, 3})
			want := []int{1, 2, 4}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("各位无进位", func() {

			got := PlusOne([]int{1})
			want := []int{2}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("各位有进位", func() {

			got := PlusOne([]int{9})
			want := []int{1, 0}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("多位进位", func() {

			got := PlusOne([]int{9, 9})
			want := []int{1, 0, 0}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

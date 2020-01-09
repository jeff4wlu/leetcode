package leet89

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGrayCode(t *testing.T) {

	Convey("TestGrayCode", t, func() {
		Convey("用例1", func() {

			got := GrayCode(2)
			want := []int{0, 1, 3, 2}

			if !infra.IntArrNonSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := GrayCode(3)
			want := []int{0, 1, 2, 3, 4, 5, 6, 7}

			if !infra.IntArrNonSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

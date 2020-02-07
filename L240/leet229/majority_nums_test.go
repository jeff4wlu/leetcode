package leet229

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMajorityNums(t *testing.T) {

	Convey("TestMajorityNums", t, func() {
		Convey("基数", func() {

			got := MajorityNums([]int{3, 2, 3})
			want := []int{3}

			if !infra.IntArrNonSeqCmp(got, want) {
				t.Errorf("failed.")
			}

		})

		Convey("偶数", func() {

			got := MajorityNums([]int{1, 1, 1, 3, 3, 2, 2, 2})
			want := []int{1, 2}

			if !infra.IntArrNonSeqCmp(got, want) {
				t.Errorf("failed.")
			}

		})

	})

}

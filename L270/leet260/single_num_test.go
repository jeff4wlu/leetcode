package leet260

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleNum(t *testing.T) {

	Convey("TestSingleNum", t, func() {
		Convey("用例1", func() {

			got := SingleNum([]int{1, 2, 1, 3, 2, 5})
			want := []int{3, 5}

			if !infra.IntArrNonSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

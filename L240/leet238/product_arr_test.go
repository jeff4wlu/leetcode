package leet238

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProductArr(t *testing.T) {

	Convey("TestProductArr", t, func() {
		Convey("用例1", func() {

			got := ProductArr([]int{1, 2, 3, 4})
			want := []int{24, 12, 8, 6}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

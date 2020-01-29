package leet163

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMissingRange(t *testing.T) {

	Convey("TestMissingRange", t, func() {
		Convey("用例1", func() {

			got := MissingRange([]int{0, 1, 3, 50, 75}, 0, 99)
			want := []string{
				"2", "4->49", "51->74", "76->99",
			}

			if !infra.StringArrSeqComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

package leet59

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpiralMatrix2(t *testing.T) {

	Convey("TestSpiralMatrix2", t, func() {
		Convey("用例1", func() {
			got := SpiralMatrix2(3)
			want := [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			}

			if !infra.IntArr2DComp(want, got) {
				t.Errorf("failed")
			}

		})

	})

}

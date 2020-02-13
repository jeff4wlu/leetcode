package leet259

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSumSmaller(t *testing.T) {

	Convey("TestSumSmaller", t, func() {

		Convey("2", func() {

			got := SumSmaller([]int{-2, 0, 1, 3}, 2)
			want := [][]int{
				{-2, 0, 1},
				{-2, 0, 3},
			}

			if !infra.IntarrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

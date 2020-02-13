package leet254

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFactorCombination(t *testing.T) {

	Convey("TestFactorCombination", t, func() {

		Convey("2", func() {

			got := FactorCombination(32)
			want := [][]int{
				{2, 16},
				{2, 2, 8},
				{2, 2, 2, 4},
				{2, 2, 2, 2, 2},
				{2, 4, 4},
				{4, 8},
			}

			if !infra.IntArr2DComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("1", func() {

			got := FactorCombination(12)
			want := [][]int{
				{2, 6},
				{2, 2, 3},
				{3, 4},
			}

			if !infra.IntArr2DComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			got := FactorCombination(37)
			want := [][]int{}

			if !infra.IntArr2DComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

package leet241

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddParenthese(t *testing.T) {

	Convey("TestAddParenthese", t, func() {
		Convey("1", func() {

			got := AddParenthese("2-1-1")
			want := []int{0, 2}

			if !infra.IntArrNonSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := AddParenthese("2*3-4*5")
			want := []int{-34, -14, -10, -10, 10}

			if !infra.IntArrNonSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

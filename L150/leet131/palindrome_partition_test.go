package leet131

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPalindromePartiion(t *testing.T) {

	Convey("TestPalindromePartiion", t, func() {
		Convey("用例1", func() {

			got := PalindromePartiion("aab")
			want := [][]string{
				{"aa", "b"},
				{"a", "a", "b"},
			}

			if !infra.StringArrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

package leet126

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWordLadder2(t *testing.T) {

	Convey("TestWordLadder2", t, func() {
		Convey("用例1", func() {

			dict := map[string]int{}
			dict["hot"] = 0
			dict["dot"] = 0
			dict["dog"] = 0
			dict["lot"] = 0
			dict["log"] = 0
			dict["cog"] = 0

			got := WordLadder2("hit", "cog", dict)
			want := [][]string{
				{"hit", "hot", "dot", "dog", "cog"},
				{"hit", "hot", "lot", "log", "cog"},
			}

			if !infra.StringArrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			dict := map[string]int{}
			dict["hot"] = 0
			dict["dot"] = 0
			dict["dog"] = 0
			dict["lot"] = 0
			dict["log"] = 0

			got := WordLadder2("hit", "cog", dict)
			want := [][]string{}

			if !infra.StringArrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

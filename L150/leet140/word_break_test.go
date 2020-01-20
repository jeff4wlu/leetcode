package leet140

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWordBreak(t *testing.T) {

	Convey("TestWordBreak", t, func() {
		Convey("用例1", func() {

			got := WordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
			want := []string{
				"cats and dog",
				"cat sand dog",
			}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := WordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
			want := []string{
				"pine apple pen apple",
				"pineapple pen apple",
				"pine applepen apple",
			}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {

			got := WordBreak("catsandog", []string{"cat", "cats", "and", "sand", "dog"})
			want := []string{}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

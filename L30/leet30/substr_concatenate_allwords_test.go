package leet30

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubstrConcatAllWords(t *testing.T) {

	Convey("TestSubstrConcatAllWords", t, func() {
		Convey("常规1", func() {

			s := "barfoothefoobarman"
			words := []string{"foo", "bar"}

			got := SubstrConcatAllWords(s, words)
			want := []int{0, 9}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("常规2", func() {

			s := "wordgoodgoodgoodbestword"
			words := []string{"word", "good", "best", "word"}

			got := SubstrConcatAllWords(s, words)
			var want []int

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

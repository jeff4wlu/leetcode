package leet212

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWordSearch2(t *testing.T) {

	Convey("TestWordSearch2", t, func() {
		Convey("常规", func() {

			board := [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}
			got := WordSearch2(board, []string{"oath", "pea", "eat", "rain"})
			want := []string{"eat", "oath"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

	})

}

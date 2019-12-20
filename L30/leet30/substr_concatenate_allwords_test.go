package leet30

import (
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

			if len(got) != len(want) {
				t.Errorf("got is not equal with want")
			} else {
				for i := 0; i < len(got); i++ {
					if got[i] != want[i] {
						t.Errorf("got is not equal with want")
						break
					}
				}
			}

		})

		Convey("常规2", func() {

			s := "wordgoodgoodgoodbestword"
			words := []string{"word", "good", "best", "word"}

			got := SubstrConcatAllWords(s, words)
			var want []int

			if len(got) != len(want) {
				t.Errorf("got is not equal with want")
			} else {
				for i := 0; i < len(got); i++ {
					if got[i] != want[i] {
						t.Errorf("got is not equal with want")
						break
					}
				}
			}

		})

	})

}

package leet139

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleNum(t *testing.T) {

	Convey("TestSingleNum", t, func() {
		Convey("用例1", func() {

			got := WordBreak("leetcode", []string{"leet", "code"})
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := WordBreak("applepenapple", []string{"apple", "pen"})
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {

			got := WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

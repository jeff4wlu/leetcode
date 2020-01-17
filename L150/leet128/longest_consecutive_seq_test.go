package leet128

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestConsecutiveSeq(t *testing.T) {

	Convey("TestLongestConsecutiveSeq", t, func() {
		Convey("用例1", func() {

			got := LongestConsecutiveSeq([]int{100, 4, 200, 1, 3, 2})
			want := 4

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

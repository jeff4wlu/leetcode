package leet234

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPalindromeLinked(t *testing.T) {

	Convey("TestPalindromeLinked", t, func() {
		Convey("1", func() {

			one := infra.MakeListNode([]int{1, 2})
			got := PalindromeLinked(one)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			one := infra.MakeListNode([]int{1, 2, 2, 1})
			got := PalindromeLinked(one)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

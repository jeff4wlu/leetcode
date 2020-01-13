package leet97

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInterLeavingStr(t *testing.T) {

	Convey("TestInterLeavingStr", t, func() {
		Convey("用例1", func() {

			got := InterLeavingStr("aabcc", "dbbca", "aadbbcbcac")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := InterLeavingStr("aabcc", "dbbca", "aadbbbaccc")
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

	Convey("TestInterLeavingStrDP", t, func() {
		Convey("用例1", func() {

			got := InterLeavingStrDP("aabcc", "dbbca", "aadbbcbcac")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := InterLeavingStrDP("aabcc", "dbbca", "aadbbbaccc")
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

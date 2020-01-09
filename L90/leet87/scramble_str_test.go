package leet87

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum(t *testing.T) {

	Convey("TestScrambleStr", t, func() {
		Convey("1", func() {
			got := ScrambleStr("great", "rgeat")
			want := true

			if got != want {
				t.Errorf("got %v, want  %v", got, want)
			}

		})

		Convey("2", func() {
			got := ScrambleStr("abcde", "caebd")
			want := false

			if got != want {
				t.Errorf("got %v, want  %v", got, want)
			}

		})

	})

}

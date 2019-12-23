package leet38

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountAndSay(t *testing.T) {

	Convey("TestCountAndSay", t, func() {
		Convey("最小边界", func() {
			got := CountAndSay(1)
			want := "1"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("常规1", func() {
			got := CountAndSay(4)
			want := "1211"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("常规2", func() {
			got := CountAndSay(12)
			want := "3113112221232112111312211312113211"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}

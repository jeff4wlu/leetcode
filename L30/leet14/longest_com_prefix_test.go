package leet14

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestComPrefix(t *testing.T) {

	Convey("TestLongestComPrefix", t, func() {
		Convey("有前缀", func() {
			ss := []string{"flower", "flow", "flight"}
			got := LongestComPrefix(ss)
			want := "fl"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("没前缀", func() {
			ss := []string{"dog", "racecar", "car"}
			got := LongestComPrefix(ss)
			want := ""

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("整个词是前缀", func() {
			ss := []string{"dog", "dogaa", "dogbb"}
			got := LongestComPrefix(ss)
			want := "dog"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("个别空串", func() {
			ss := []string{"dge", "dg", ""}
			got := LongestComPrefix(ss)
			want := ""

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("全空串", func() {
			ss := []string{"", "", ""}
			got := LongestComPrefix(ss)
			want := ""

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}

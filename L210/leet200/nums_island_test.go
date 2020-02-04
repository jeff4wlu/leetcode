package leet200

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumsIsland(t *testing.T) {

	Convey("TestNumsIsland", t, func() {
		Convey("用例1", func() {

			m := []string{
				"11110",
				"11010",
				"11000",
				"00000",
			}

			got := NumsIsland(m)
			want := 1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			m := []string{
				"11000",
				"11000",
				"00100",
				"00011",
			}

			got := NumsIsland(m)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}

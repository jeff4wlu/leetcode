package leet13

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRomanToInt(t *testing.T) {

	Convey("TestRomanToInt", t, func() {
		Convey("1", func() {
			s := "III"
			got := RomanToInt(s)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {
			s := "IV"
			got := RomanToInt(s)
			want := 4

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("3", func() {
			s := "IX"
			got := RomanToInt(s)
			want := 9

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("4", func() {
			s := "LVIII"
			got := RomanToInt(s)
			want := 58

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("5", func() {
			s := "MCMXCIV"
			got := RomanToInt(s)
			want := 1994

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}

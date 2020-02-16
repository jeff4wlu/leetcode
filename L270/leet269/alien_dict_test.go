package leet269

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAlienDict(t *testing.T) {

	Convey("TestAlienDict", t, func() {

		Convey("1", func() {

			words := []string{
				"wrt",
				"wrf",
				"er",
				"ett",
				"rftt",
			}

			got := AlienDict(words)
			want := "wertf"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("2", func() {

			words := []string{
				"z",
				"x",
			}

			got := AlienDict(words)
			want := "zx"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("3", func() {

			words := []string{
				"z",
				"x",
				"z",
			}

			got := AlienDict(words)
			want := ""

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("4", func() {

			words := []string{
				"ac",
				"aa",
				"b",
				"c",
			}

			got := AlienDict(words)
			want := ""

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}

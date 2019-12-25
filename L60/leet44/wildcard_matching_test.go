package leet44

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWildcardMatching(t *testing.T) {

	Convey("TestWildcardMatching", t, func() {
		Convey("1", func() {
			s := "acdcb"
			p := "a*c?b"
			got := WildcardMatching(s, p)
			want := false

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

		Convey("2", func() {
			s := "aa"
			p := "a"
			got := WildcardMatching(s, p)
			want := false

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

		Convey("3", func() {
			s := "aa"
			p := "*"
			got := WildcardMatching(s, p)
			want := true

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

		Convey("4", func() {
			s := "cb"
			p := "?a"
			got := WildcardMatching(s, p)
			want := false

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

		Convey("5", func() {
			s := "adceb"
			p := "*a*b"
			got := WildcardMatching(s, p)
			want := true

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

	})

}

package leet10

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegularMatching(t *testing.T) {

	Convey("TestRegularMatching", t, func() {
		Convey("纯字符串", func() {
			s := "aa"
			p := "a"
			got := RegularMatching(s, p)
			want := false

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}

		})

		Convey("*通配符", func() {
			s := "aa"
			p := "a*"
			got := RegularMatching(s, p)
			want := true

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}

		})

		Convey(".*通配符", func() {
			s := "ab"
			p := ".*"
			got := RegularMatching(s, p)
			want := true

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}

		})

		Convey("用例4", func() {
			s := "aab"
			p := "c*a*b"
			got := RegularMatching(s, p)
			want := true

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}

		})

		Convey("用例5", func() {
			s := "mississippi"
			p := "mis*is*p*."
			got := RegularMatching(s, p)
			want := false

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}

		})

	})

}

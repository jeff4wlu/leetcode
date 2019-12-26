package leet50

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPow(t *testing.T) {

	Convey("TestPow", t, func() {
		Convey("1", func() {
			got := Pow(-2, 31)
			want := -2147483648.00000000000000000

			if got != want {
				t.Errorf("got %27.17f, want  %27.17f", got, want)
			}

		})

		Convey("2", func() {
			got := Pow(2.1, 3)
			want := 9.26100000000000101

			if got != want {
				t.Errorf("got %27.17f, want  %27.17f", got, want)
			}

		})

	})

	Convey("TestPow2", t, func() {
		Convey("1", func() {
			got := Pow2(-2, 31)
			want := -2147483648.00000000000000000

			if got != want {
				t.Errorf("got %27.17f, want  %27.17f", got, want)
			}

		})

		Convey("2", func() {
			got := Pow2(2.1, 3)
			want := 9.26100000000000101

			if got != want {
				t.Errorf("got %27.17f, want  %27.17f", got, want)
			}

		})

		Convey("3", func() {
			got := Pow2(2.0, 4)
			want := 9.26100000000000101

			if got != want {
				t.Errorf("got %27.17f, want  %27.17f", got, want)
			}

		})

	})

}

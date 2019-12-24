package leet43

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMultiplyString(t *testing.T) {

	Convey("TestMultiplyString", t, func() {
		Convey("1", func() {
			a := "123"
			b := "456"
			got := MultiplyString(a, b)
			want := "56088"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("2", func() {
			a := "2"
			b := "3"
			got := MultiplyString(a, b)
			want := "6"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("3", func() {
			a := "99"
			b := "99"
			got := MultiplyString(a, b)
			want := "9801"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}

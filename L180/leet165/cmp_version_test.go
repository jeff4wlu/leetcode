package leet165

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCmpVersion(t *testing.T) {

	Convey("TestCmpVersion", t, func() {
		Convey("用例1", func() {

			got := CmpVersion("0.1", "1.1")
			want := -1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := CmpVersion("1.0.1", "1")
			want := 1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例3", func() {

			got := CmpVersion("7.5.2.4", "7.5.3")
			want := -1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}

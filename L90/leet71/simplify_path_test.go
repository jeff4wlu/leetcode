package leet71

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSimplifyPath(t *testing.T) {

	Convey("TestSimplifyPath", t, func() {
		Convey("用例1", func() {

			got := SimplifyPath("/home/")
			want := "/home"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("用例2", func() {

			got := SimplifyPath("/a/./b/../../c/")
			want := "/c"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}

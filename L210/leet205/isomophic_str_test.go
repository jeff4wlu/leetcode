package leet205

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsomophicStr(t *testing.T) {

	Convey("TestIsomophicStr", t, func() {
		Convey("用例1", func() {

			got := IsomophicStr("egg", "eee")
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := IsomophicStr("foo", "bar")
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {

			got := IsomophicStr("paper", "title")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

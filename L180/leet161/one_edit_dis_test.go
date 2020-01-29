package leet161

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOneEditDis(t *testing.T) {

	Convey("TestOneEditDis", t, func() {
		Convey("字数不相等", func() {

			got := OneEditDis("ab", "acb")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("false", func() {

			got := OneEditDis("cab", "ad")
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("相同字数", func() {

			got := OneEditDis("1203", "1213")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

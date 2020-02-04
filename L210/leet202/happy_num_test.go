package leet202

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHappyNum(t *testing.T) {

	Convey("TestHappyNum", t, func() {
		Convey("用例1", func() {

			got := HappyNum(19)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("死循环", func() {

			got := HappyNum(11)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

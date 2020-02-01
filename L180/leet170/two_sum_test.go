package leet170

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum(t *testing.T) {

	Convey("TestTwoSum", t, func() {
		Convey("基数", func() {

			twoSum := New()
			twoSum.add(1)
			twoSum.add(3)
			twoSum.add(5)

			got := twoSum.find(4)
			want := true

			if got != want {
				t.Errorf("failed.")
			}

			got = twoSum.find(7)
			want = false

			if got != want {
				t.Errorf("failed.")
			}

		})

	})

}

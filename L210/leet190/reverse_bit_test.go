package leet190

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseBit(t *testing.T) {

	Convey("TestReverseBit", t, func() {
		Convey("用例1", func() {

			got := ReverseBit(43261596)
			want := uint32(964176192)

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := ReverseBit(4294967293)
			want := uint32(3221225471)

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}

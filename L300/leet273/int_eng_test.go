package leet273

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt2Eng(t *testing.T) {

	Convey("TestInt2Eng", t, func() {

		Convey("1", func() {

			got := Int2Eng(123)
			want := "One Hundred Twenty Three"

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := Int2Eng(12345)
			want := "Twelve Thousand Three Hundred Forty Five"

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			got := Int2Eng(1234567)
			want := "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("4", func() {

			got := Int2Eng(1000010)
			want := "One Million Ten"

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("5", func() {

			got := Int2Eng(0)
			want := "Zero"

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

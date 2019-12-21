package leet6

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestZigzagConversion(t *testing.T) {

	Convey("TestZigzagConversion", t, func() {
		Convey("测试1", func() {
			in := "PAYPALISHIRING"
			got := ZigzagConversion(in, 3)
			want := "PAHNAPLSIIGYIR"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("测试2", func() {
			in := "PAYPALISHIRING"
			got := ZigzagConversion(in, 4)
			want := "PINALSIGYAHRPI"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}

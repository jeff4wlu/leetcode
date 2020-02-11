package leet247

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrobogrammaticNum(t *testing.T) {

	Convey("TestStrobogrammaticNum", t, func() {
		Convey("偶数", func() {

			got := StrobogrammaticNum(2)
			want := []string{"11", "69", "88", "96"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("基数", func() {

			got := StrobogrammaticNum(3)
			want := []string{"111", "181", "818", "888", "689", "986", "619", "916"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("基数1", func() {

			got := StrobogrammaticNum(1)
			want := []string{"1", "8"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

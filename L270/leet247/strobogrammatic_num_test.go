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
			want := []string{"111", "181", "818", "888", "689", "986", "619", "916", "101", "609", "808", "906"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("基数1", func() {

			got := StrobogrammaticNum(4)
			want := []string{"1001", "6009", "8008", "9006", "1111", "6119", "8118", "9116", "1691", "6699", "8698", "9696", "1881", "6889", "8888", "9886", "1961", "6969", "8968", "9966"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

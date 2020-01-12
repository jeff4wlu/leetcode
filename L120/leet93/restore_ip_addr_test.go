package leet93

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRestoreIPAddr(t *testing.T) {

	Convey("TestRestoreIPAddr", t, func() {
		Convey("用例1", func() {

			got := RestoreIPAddr("25525511135")
			want := []string{
				"255.255.11.135", "255.255.111.35",
			}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := RestoreIPAddr("166611")
			want := []string{
				"1.6.66.11",
				"1.66.6.11",
				"1.66.61.1",
				"16.6.6.11",
				"16.6.61.1",
				"16.66.1.1",
				"166.6.1.1",
			}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {

			got := RestoreIPAddr("1200121")
			want := []string{
				"1.200.1.21", "1.200.12.1",
			}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

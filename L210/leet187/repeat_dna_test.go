package leet187

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRepeatDNA(t *testing.T) {

	Convey("TestRepeatDNA", t, func() {
		Convey("用例1", func() {

			got := RepeatDNA("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
			want := []string{"AAAAACCCCC", "CCCCCAAAAA"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

	Convey("TestRepeatDNA1", t, func() {
		Convey("用例1", func() {

			got := RepeatDNA1("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
			want := []string{"AAAAACCCCC", "CCCCCAAAAA"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

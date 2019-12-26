package leet17

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLetterComPhonenum(t *testing.T) {

	Convey("TestLetterComPhonenum", t, func() {
		Convey("常规", func() {
			num := "23"
			got := LetterComPhonenum(num)
			want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

			if got == nil || !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

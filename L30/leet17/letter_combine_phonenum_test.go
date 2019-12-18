package leet17

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLetterComPhonenum(t *testing.T) {

	Convey("TestLetterComPhonenum", t, func() {
		Convey("常规", func() {
			num := "23"
			got := LetterComPhonenum(num)
			want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

			if got == nil || !comp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

func comp(a, b []string) bool {
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}

	}
	return true
}

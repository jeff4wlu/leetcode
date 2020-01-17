package leet125

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidPalindrome(t *testing.T) {

	Convey("TestValidPalindrome", t, func() {
		Convey("用例1", func() {

			got := ValidPalindrome("A man, a plan, a canal: Panama")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := ValidPalindrome("race a car")
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

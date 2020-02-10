package leet242

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidAnagram(t *testing.T) {

	Convey("TestValidAnagram", t, func() {
		Convey("1", func() {

			got := ValidAnagram("anagram", "nagaram")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := ValidAnagram("rat", "car")
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

package leet8

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrToInt(t *testing.T) {

	Convey("TestStrToInt", t, func() {
		Convey("正数", func() {
			in := "42"
			got := StrToInt(in)
			want := 42

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("负数", func() {
			in := "-42"
			got := StrToInt(in)
			want := -42

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("数字在前", func() {
			in := "4193 with words"
			got := StrToInt(in)
			want := 4193

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("字符在前", func() {
			in := "words and 987"
			got := StrToInt(in)
			want := 0

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("溢出", func() {
			in := "-91283472332"
			got := StrToInt(in)
			want := -2147483648

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("前置了空格", func() {
			in := "   -345"
			got := StrToInt(in)
			want := -345

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}

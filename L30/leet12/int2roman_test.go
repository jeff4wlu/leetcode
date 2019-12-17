package leet12

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntToRoman(t *testing.T) {

	Convey("TestIntToRoman", t, func() {
		Convey("常规测试1", func() {
			in := 3
			got := IntToRoman(in)
			want := "III"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
			//ShouldNotBeNil(err)
		})

		Convey("常规测试2", func() {
			in := 4
			got := IntToRoman(in)
			want := "IV"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
			//ShouldNotBeNil(err)
		})

		Convey("常规测试3", func() {
			in := 9
			got := IntToRoman(in)
			want := "IX"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
			//ShouldNotBeNil(err)
		})

		Convey("常规测试4", func() {
			in := 58
			got := IntToRoman(in)
			want := "LVIII"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
			//ShouldNotBeNil(err)
		})

		Convey("常规测试5", func() {
			in := 1994
			got := IntToRoman(in)
			want := "MCMXCIV"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
			//ShouldNotBeNil(err)
		})

	})

}

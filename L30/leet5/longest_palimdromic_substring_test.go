package leet5

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestPalimdromicSubstr(t *testing.T) {

	Convey("TestLongestPalimdromicSubstr", t, func() {
		Convey("总元素为基数情况", func() {
			in := "bscecee"
			got := LongestPalimdromicSubstr(in)
			want := "cec"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("总元素为偶数情况", func() {
			in := "bscenooncee"
			got := LongestPalimdromicSubstr(in)
			want := "noon"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("重复多位", func() {
			in := "bsceooooooonee"
			got := LongestPalimdromicSubstr(in)
			want := "ooooooo"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}

func TestLongestPalimdromicSubstr1(t *testing.T) {

	Convey("TestLongestPalimdromicSubstr1", t, func() {
		Convey("总元素为基数情况", func() {
			in := "aabcbd"
			got := LongestPalimdromicSubstr1(in)
			want := "bcb"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("总元素为偶数情况", func() {
			in := "bscenooncee"
			got := LongestPalimdromicSubstr1(in)
			want := "noon"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("重复多位", func() {
			in := "bscenooooooncee"
			got := LongestPalimdromicSubstr1(in)
			want := "noooooon"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}

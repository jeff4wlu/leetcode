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
			//ShouldNotBeNil(err)
		})

		Convey("总元素为偶数情况", func() {
			in := "bscenooncee"
			got := LongestPalimdromicSubstr(in)
			want := "noon"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
			//ShouldNotBeNil(err)
		})

		Convey("重复多位", func() {
			in := "bscenoooooooncee"
			got := LongestPalimdromicSubstr(in)
			want := "nooooooon"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
			//ShouldNotBeNil(err)
		})

	})

}

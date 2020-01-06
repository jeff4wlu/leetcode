package leet61

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRotateList(t *testing.T) {

	Convey("TestRotateList", t, func() {
		Convey("1", func() {
			one := infra.MakeListNilEnd([]int{1, 2, 3, 4, 5})
			got := RotateList(one, 2)
			want := infra.MakeListNilEnd([]int{4, 5, 1, 2, 3})

			for got != nil && want != nil {
				if got.Value != want.Value {
					t.Errorf("got is not equal with want")
					break
				}
				got = got.Next
				want = want.Next
			}

		})

		Convey("超过一个周期", func() {
			one := infra.MakeListNilEnd([]int{1, 2, 3, 4, 5})
			got := RotateList(one, 7)
			want := infra.MakeListNilEnd([]int{4, 5, 1, 2, 3})

			for got != nil && want != nil {
				if got.Value != want.Value {
					t.Errorf("got is not equal with want")
					break
				}
				got = got.Next
				want = want.Next
			}

		})

		Convey("边界1", func() {
			one := infra.MakeListNilEnd([]int{1, 2, 3, 4, 5})
			got := RotateList(one, 4)
			want := infra.MakeListNilEnd([]int{2, 3, 4, 5, 1})

			for got != nil && want != nil {
				if got.Value != want.Value {
					t.Errorf("got is not equal with want")
					break
				}
				got = got.Next
				want = want.Next
			}

		})
		Convey("边界2", func() {
			one := infra.MakeListNilEnd([]int{1, 2, 3, 4, 5})
			got := RotateList(one, 5)
			want := infra.MakeListNilEnd([]int{1, 2, 3, 4, 5})

			for got != nil && want != nil {
				if got.Value != want.Value {
					t.Errorf("got is not equal with want")
					break
				}
				got = got.Next
				want = want.Next
			}

		})

	})

}
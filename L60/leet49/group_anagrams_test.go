package leet49

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGroupAnagrams(t *testing.T) {

	Convey("TestGroupAnagrams", t, func() {
		Convey("成功", func() {
			arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
			got := GroupAnagrams(arr)

			want := [][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			}

			if !infra.StringArrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

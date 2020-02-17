package leet271

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncoder(t *testing.T) {

	Convey("TestEncoder", t, func() {

		Convey("1", func() {

			words := []string{
				"Jeff", "is", "an", "engineer",
			}

			got := Decode(Encode(words))
			want := words

			if !infra.StringArrSeqComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("有特殊asc字符", func() {

			words := []string{
				"Je$ff", "is*", "/a\n", "/eng/ineer",
			}

			got := Decode(Encode(words))
			want := words

			if !infra.StringArrSeqComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

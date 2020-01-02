package leet68

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTextJustification(t *testing.T) {

	Convey("TestTextJustification", t, func() {
		Convey("用例1", func() {

			got := TextJustification([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
			want := []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			}

			if !infra.StringArrSeqComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("最后一行的特殊处理的测试", func() {

			got := TextJustification([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)
			want := []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			}

			if !infra.StringArrSeqComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {

			got := TextJustification([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
			want := []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			}

			if !infra.StringArrSeqComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("只有一个单词", func() {

			got := TextJustification([]string{"This"}, 16)
			want := []string{
				"This            ",
			}

			if !infra.StringArrSeqComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

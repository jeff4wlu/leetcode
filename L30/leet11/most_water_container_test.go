package leet11

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMostWaterContainer(t *testing.T) {

	Convey("TestMostWaterContainer", t, func() {
		Convey("常规测试", func() {
			arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
			got := MostWaterContainer(arr)
			want := 49

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			//ShouldNotBeNil(err)
		})

	})

}

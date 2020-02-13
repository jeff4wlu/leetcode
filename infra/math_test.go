package infra

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPrime(t *testing.T) {

	Convey("TestIsPrime", t, func() {
		Convey("1", func() {

			got := IsPrime(11)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := IsPrime(12)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

func TestGetPrimes(t *testing.T) {

	Convey("TestGetPrimes", t, func() {
		Convey("1", func() {

			got := GetPrimes(12)
			want := []int{2, 3, 5, 7, 11}

			if !IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

package leet146

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLRUCache(t *testing.T) {

	Convey("TestLRUCache", t, func() {
		Convey("用例1", func() {

			cache := InitLRUCache(2)
			cache.Put(1, 1)
			cache.Put(2, 2)
			got := cache.Get(1)
			want := 1
			if got != want {
				t.Errorf("failed")
			}

			cache.Put(3, 3)
			got = cache.Get(2)
			want = -1
			if got != want {
				t.Errorf("failed")
			}

			cache.Put(4, 4)    // evicts key 1
			got = cache.Get(1) // returns -1 (not found)
			want = -1
			if got != want {
				t.Errorf("failed")
			}

			got = cache.Get(3) // returns 3
			want = 3
			if got != want {
				t.Errorf("failed")
			}

			got = cache.Get(4) // returns 4
			want = 4
			if got != want {
				t.Errorf("failed")
			}

		})

	})

}

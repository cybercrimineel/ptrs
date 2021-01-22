package ptrs

import (
	"sync"
	"testing"
)

const count = 99999

func Benchmark(b *testing.B) {
	for idx := 0; idx < b.N; idx++ {
		r := Alloc(idx)
		Deref(r)
		Free(r)
	}
}

func Test(t *testing.T) {
	t.Run("synchronous", func(t *testing.T) {
		r1 := Alloc('a')
		r2 := Alloc('b')
		Free(r1)
		Free(r2)
		r3 := Alloc('c')
		r4 := Alloc('d')

		if Deref(r1) != 'c' || Deref(r2) != 'd' || r1 != r3 || r2 != r4 {
			t.Fail()
		}
	})

	t.Run("asynchronous", func(t *testing.T) {
		wg := sync.WaitGroup{}

		for idx := 0; idx < count; idx++ {
			wg.Add(1)

			go func(v int) {
				r := Alloc(v)

				if Deref(r) != v {
					t.Fail()
				}

				Free(r)
				wg.Done()
			}(idx)
		}

		wg.Wait()
	})
}

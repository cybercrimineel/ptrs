package ptrs

import (
	"sync"
)

var (
	idx = uintptr(0)
	mem = []interface{}{nil}
	mtx = sync.Mutex{}
	old = []uintptr{}
)

// Alloc provides a reference to the given value
func Alloc(val interface{}) uintptr {
	mtx.Lock()
	defer mtx.Unlock()

	if len(old) != 0 {
		defer func() {
			old = old[1:]
		}()

		mem[old[0]] = val
		return old[0]
	}

	defer func() {
		idx++
	}()

	mem[idx] = val
	mem = append(mem, nil)
	return idx
}

// Deref provides the referenced value of a given reference
func Deref(ref uintptr) interface{} {
	mtx.Lock()
	defer mtx.Unlock()
	return mem[ref]
}

// Free marks a given reference as overwritable
func Free(ref uintptr) {
	mtx.Lock()
	mem[ref] = nil
	old = append(old, ref)
	mtx.Unlock()
}

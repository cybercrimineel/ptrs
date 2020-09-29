package ptrs

import "sync"

var (
	mem = []interface{}{}
	mtx = sync.RWMutex{}
)

// Alloc allocates a reference
func Alloc(v interface{}) int {
	i := 0
	mtx.Lock()

	for ; i < len(mem); i++ {
		if mem[i] == nil {
			mem[i] = v
			mtx.Unlock()
			return i
		}
	}

	mem = append(mem, v)
	mtx.Unlock()
	return i
}

// Deref dereferences a reference
func Deref(i int) interface{} {
	mtx.RLock()
	defer mtx.RUnlock()
	return mem[i]
}

// DerefAndFree dereferences and frees a reference
func DerefAndFree(i int) interface{} {
	defer Free(i)
	return Deref(i)
}

// Free frees a reference
func Free(i int) {
	mtx.Lock()
	mem[i] = nil
	mtx.Unlock()
}

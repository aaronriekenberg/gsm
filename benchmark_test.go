package gsm_test

import (
	"sync"
	"testing"

	"github.com/aaronriekenberg/gsm"
)

func BenchmarkSyncMapParallelLoad(b *testing.B) {
	var m sync.Map

	m.Store(1, "one")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Load(1)
		}
	})
}

func BenchmarkGenericSyncMapParallelLoad(b *testing.B) {
	var m gsm.GenericSyncMap[int, string]

	m.Store(1, "one")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Load(1)
		}
	})
}

func BenchmarkRWMutexMapParallelLoad(b *testing.B) {
	m := make(map[int]string)
	rwm := sync.RWMutex{}

	m[1] = "one"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rwm.RLock()
			defer rwm.RUnlock()

			_ = m[1]
		}
	})
}

package genericsyncmap

import (
	"iter"
	"sync"
)

type GenericSyncMap[K comparable, V any] struct {
	syncMap sync.Map
}

func (gsm *GenericSyncMap[K, V]) Store(
	key K,
	value V,
) {

	gsm.syncMap.Store(key, value)
}

func (gsm *GenericSyncMap[K, V]) Load(
	key K,
) (value V, ok bool) {

	anyValue, ok := gsm.syncMap.Load(key)
	if !ok {
		return
	}

	value = anyValue.(V)

	return
}

func (gsm *GenericSyncMap[K, V]) LoadAndDelete(
	key K,
) (value V, loaded bool) {

	anyValue, loaded := gsm.syncMap.LoadAndDelete(key)

	if !loaded {
		return
	}

	value = anyValue.(V)
	return
}

func (gsm *GenericSyncMap[K, V]) Range() iter.Seq2[K, V] {

	return func(yield func(K, V) bool) {

		for anyKey, anyValue := range gsm.syncMap.Range {
			key := anyKey.(K)
			value := anyValue.(V)

			if !yield(key, value) {
				break
			}
		}
	}
}

func (gsm *GenericSyncMap[K, V]) Values() iter.Seq[V] {

	return func(yield func(V) bool) {

		for _, anyValue := range gsm.syncMap.Range {
			value := anyValue.(V)

			if !yield(value) {
				break
			}
		}
	}
}

package extended_map

import (
	"fmt"
	"sort"
)

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

type Keys[K comparable] []K

func (k Keys[K]) Len() int           { return len(k) }
func (k Keys[K]) Less(i, j int) bool { return fmt.Sprintf("%v", k[i]) < fmt.Sprintf("%v", k[j]) }
func (k Keys[K]) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }

type ExtendedMap[K comparable, V any] struct {
	data map[K]V
	keys []K
}

func NewExtendedMap[K comparable, V any]() *ExtendedMap[K, V] {
	return &ExtendedMap[K, V]{
		data: make(map[K]V),
		keys: []K{},
	}
}

func (m *ExtendedMap[K, V]) Get(key K) (V, bool) {
	value, exists := m.data[key]
	return value, exists
}

func (m *ExtendedMap[K, V]) Push(key K, value V) {
	m.data[key] = value
	m.keys = append(m.keys, key)
}

func (m *ExtendedMap[K, V]) Set(key K, value V) {
	m.data[key] = value
}

func (m *ExtendedMap[K, V]) Pop(key K) (V, bool) {
	value, exists := m.data[key]
	if exists {
		delete(m.data, key)
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
	return value, exists
}

func (m *ExtendedMap[K, V]) Sort() []Pair[K, V] {
	keys := Keys[K](m.keys)
	sort.Sort(keys)
	var sortedPairs []Pair[K, V]
	for _, key := range m.keys {
		sortedPairs = append(sortedPairs, Pair[K, V]{Key: key, Value: m.data[key]})
	}
	return sortedPairs
}

package extended_map

import (
	"reflect"
	"testing"
)

func TestNewExtendedMap(t *testing.T) {
	m := NewExtendedMap[string, int]()
	if m == nil || len(m.data) != 0 || len(m.keys) != 0 {
		t.Errorf("NewExtendedMap() failed, expected an empty map, got %v", m)
	}
}

func TestPush(t *testing.T) {
	m := NewExtendedMap[string, int]()
	m.Push("a", 10)
	m.Push("b", 20)
	if len(m.data) != 2 || len(m.keys) != 2 || m.data["a"] != 10 || m.data["b"] != 20 {
		t.Errorf("Push() failed, got %v", m)
	}
}

func TestSet(t *testing.T) {
	m := NewExtendedMap[string, int]()
	m.Push("a", 10)
	m.Set("a", 20)
	if m.data["a"] != 20 {
		t.Errorf("Set() failed, got %v", m)
	}
}

func TestGet(t *testing.T) {
	m := NewExtendedMap[string, int]()
	m.Push("a", 10)
	value, exists := m.Get("a")
	if !exists || value != 10 {
		t.Errorf("Get() failed, got %v", value)
	}
	_, exists = m.Get("b")
	if exists {
		t.Errorf("Get() failed, key should not exist")
	}
}

func TestPop(t *testing.T) {
	m := NewExtendedMap[string, int]()
	m.Push("a", 10)
	value, exists := m.Pop("a")
	if !exists || value != 10 || len(m.keys) != 0 {
		t.Errorf("Pop() failed, got %v", value)
	}
	_, exists = m.Pop("b")
	if exists {
		t.Errorf("Pop() failed, key should not exist")
	}
}

func TestSort(t *testing.T) {
	m := NewExtendedMap[string, int]()
	m.Push("b", 20)
	m.Push("a", 10)
	sortedPairs := m.Sort()
	expectedPairs := []Pair[string, int]{
		{Key: "a", Value: 10},
		{Key: "b", Value: 20},
	}
	if !reflect.DeepEqual(sortedPairs, expectedPairs) {
		t.Errorf("Sort() failed, got %v", sortedPairs)
	}
}

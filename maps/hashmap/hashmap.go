// Package hashmap implements a map backed by a hash table.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package hashmap

import (
	"fmt"
	"github.com/itchenyi/common/maps"
	"github.com/itchenyi/common/utils"
)

func assertMapImplementation() {
	var _ maps.Map = (*Map)(nil)
}

type HashMap map[interface{}]interface{}

// Map holds the elements in go's native map
type Map struct{ m HashMap }

// New instantiates a hash map.
func New() *Map {
	return &Map{m: make(map[interface{}]interface{})}
}

// Put inserts element into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	m.m[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	value, found = m.m[key]
	return
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key interface{}) {
	delete(m.m, key)
}

// Empty returns true if map does not contain any elements
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return len(m.m)
}

// Keys returns all keys (random order).
func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

// Values returns all values (random order).
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

// StringVal returns string value
func (m *Map) StringVal(key interface{}) (val string) {
	if value, ok := m.Get(key); ok {
		return value.(string)
	}

	return val
}

// IntVal returns int value
func (m *Map) IntVal(key interface{}) (val int) {
	if value, ok := m.Get(key); ok {
		return value.(int)
	}

	return val
}

// StringItems returns Json Format Item
func (m *Map) StringItems() map[string]interface{} {
	elements := make(map[string]interface{})
	for key, value := range m.m {
		elements[utils.ToString(key)] = value
	}

	return elements
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.m = make(map[interface{}]interface{})
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}

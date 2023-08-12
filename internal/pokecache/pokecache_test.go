package pokecache

import (
	"testing"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "foo1",
			inputVal: []byte("bar1"),
		},
		{
			inputKey: "foo2",
			inputVal: []byte("bar2"),
		},
		{
			inputKey: "foo3",
			inputVal: []byte("bar3"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)
		actual, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Errorf("%s not found", cs.inputKey)
			continue
		}

		if string(actual) != string(cs.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(cs.inputVal))
		}
	}
}

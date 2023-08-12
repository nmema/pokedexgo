package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)

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

func TestReapOk(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "foo"
	cache.Add(keyOne, []byte("bar"))

	time.Sleep(interval + time.Millisecond*10)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}

func TestReapFailed(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "foo"
	cache.Add(keyOne, []byte("bar"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s shouldn't have been reaped", keyOne)
	}
}

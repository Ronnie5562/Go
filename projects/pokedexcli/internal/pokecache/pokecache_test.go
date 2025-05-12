package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	if cache.cache == nil {
		t.Errorf("Expected cache to be initialized, got nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
		{
			inputKey: "",
			inputVal: []byte("value3"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)
		actual, ok := cache.Get(cs.inputKey)

		if !ok {
			t.Errorf(
				"Expected %s to be present in cache, got false",
				cs.inputKey,
			)
			continue
		}

		if string(actual) != string(cs.inputVal) {
			t.Errorf(
				"Expected %s, got %s",
				string(cs.inputVal),
				string(actual),
			)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))

	time.Sleep(interval + time.Second)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("Expected %s to be removed from cache, got true", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyTwo := "key2"
	cache.Add(keyTwo, []byte("value2"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyTwo)
	if !ok {
		t.Errorf("Expected %s to still be in cache, got false", keyTwo)
	}
}

package pokecache

import (
	"reflect"
	"testing"
	"time"
)

func TestCacheMethods(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	key := "test"
	val := []byte("pikaaaaa")
	cache.Add(key, val)
	data, is_cached := cache.Get(key)
	if !is_cached || !reflect.DeepEqual(data, val) {
		t.Errorf("Wrong cached value: %v; want %v", data, val)

	}
	time.Sleep(20 * time.Millisecond)
	data, is_cached = cache.Get(key)
	if is_cached || data != nil {
		t.Errorf("Cache should have been emptied but is: %v", data)

	}
}

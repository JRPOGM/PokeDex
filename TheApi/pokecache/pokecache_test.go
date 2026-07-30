package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key		string
		value	[]byte
	}{
		{
			key:		"https://allpoetry.com",
			value:		[]byte("testdata"),
		},
		{
			key:		"https://allpoetry.com/Jrpogm?tab=links",
			value:		[]byte("moretestdata"),
		},
		{
			key:		"https://allpoetry.com/poem/19227703-God-is-All-God-is-Not-by-Jrpogm",
			value:		[]byte("extratestdata"),
		},
		{
			key:		"https://allpoetry.com/groups",
			value:		[]byte("finaltestdata"),
		},
	}
	for a, b := range cases {
		t.Run(fmt.Sprintf("Test case %v", a), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(b.key, b.value)
			value, ok := cache.Get(b.key)
			if !ok {
				t.Errorf("unable to find key")
				return
			}
			if string(value) != string(b.value) {
				t.Errorf("unable to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://allpoetry.com", []byte("testdata"))
	_, ok := cache.Get("https://allpoetry.com")
	if !ok {
		t.Errorf("unable to find key")
		return
	}
	time.Sleep(waitTime)
	_, ok = cache.Get("https://allpoetry.com")
	if ok {
		t.Errorf("unable to not find key")
		return
	}
}
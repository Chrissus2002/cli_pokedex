package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet (t *testing.T){
	const interval = time.Minute * 5
	cases := []struct{
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://betterexample.com",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases{
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T){
			cache := NewCache(interval)
			cache.Add(c.key,c.val)
			val, ok := cache.Get(c.key)
			if !ok{
				t.Errorf("Expected to find a key")
				return
			}
			if string(val) != string(c.val){
				t.Errorf("The value data is incorrect")
				return
			}
		})
	}
}

func TestReapLoop (t *testing.T){
	const baseTime = time.Millisecond * 5
	const waitTime = baseTime + time.Millisecond * 5
	const url = "https://Pokemon.com"

	cache := NewCache(baseTime)
	cache.Add(url, []byte("testdata"))

	_, ok := cache.Get(url)
	if !ok{
		t.Errorf("Expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(url)
	if ok{
		t.Errorf("Expected not to find key")
		return
	}

}

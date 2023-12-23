package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require" //nolint:depguard
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("clear", func(t *testing.T) {
		// Write me
		c := NewCache(3)
		for i := 0; i < 2; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
		c.Clear()
		val, ok := c.Get("0")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		// a lot of set
		c := NewCache(3)
		for i := 0; i < 4; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
		val, ok := c.Get("0") // [3,2,1]
		require.False(t, ok)
		require.Nil(t, val)

		// set exist item and order
		wasInCache := c.Set("1", 1) // [1,3,2]
		require.True(t, wasInCache)
		val, ok = c.Get("2") // [2,1,3]
		require.True(t, ok)
		require.Equal(t, 2, val)

		// set new item and order
		wasInCache = c.Set("4", 4) // [4,2,1]
		require.False(t, wasInCache)

		val, ok = c.Get("3") // [4,2,1]
		require.False(t, ok)
		require.Nil(t, val)

		// get exist item and order
		val, ok = c.Get("1") // [1,4,2]
		require.True(t, ok)
		require.Equal(t, 1, val)
		wasInCache = c.Set("5", 5) // [5,1,4]
		require.False(t, wasInCache)
		val, ok = c.Get("1") // [1,5,4]
		require.True(t, ok)
		require.Equal(t, 1, val)

		// get not exist item and order
		val, ok = c.Get("0") // [1,5,4]
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(_ *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

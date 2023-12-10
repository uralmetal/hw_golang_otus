package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPushFront(t *testing.T) {
	l := NewList()
	t.Run("test with len = 1", func(t *testing.T) {
		l.PushFront(10)
		require.Equal(t, 1, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Prev)
	})

	t.Run("test with len = 2", func(t *testing.T) {
		l.PushFront(20)
		require.Equal(t, 2, l.Len())
		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, l.Front(), l.Back().Prev)
		require.Equal(t, l.Back(), l.Front().Next)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Front().Prev)
	})

	t.Run("test with len = 3", func(t *testing.T) {
		l.PushFront(30)
		require.Equal(t, 3, l.Len())
		require.Equal(t, 30, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, l.Front().Next, l.Back().Prev)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Front().Prev)
	})
}

func TestPushBack(t *testing.T) {
	l := NewList()
	t.Run("test with len = 1", func(t *testing.T) {
		l.PushBack(10)
		require.Equal(t, 1, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Prev)
	})

	t.Run("test with len = 2", func(t *testing.T) {
		l.PushBack(20)
		require.Equal(t, 2, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 20, l.Back().Value)
		require.Equal(t, l.Front(), l.Back().Prev)
		require.Equal(t, l.Back(), l.Front().Next)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Front().Prev)
	})

	t.Run("test with len = 3", func(t *testing.T) {
		l.PushBack(30)
		require.Equal(t, 3, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)
		require.Equal(t, l.Front().Next, l.Back().Prev)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Front().Prev)
	})
}

func TestListRemove(t *testing.T) {
	// ToDo: testing memory leaks, but it seems not required now
	t.Run("remove with len = 1", func(t *testing.T) {
		// current ----> *empty list*
		l := NewList()
		item := l.PushFront(10)

		l.Remove(item)
		require.Nil(t, item.Value)
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("remove front with len = 2", func(t *testing.T) {
		// current <-> next ----> next
		l := NewList()

		l.PushFront(10)
		l.Remove(l.PushFront(20))

		require.Equal(t, 1, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, l.Front(), l.Back())
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Prev)
	})

	t.Run("remove back with len = 2", func(t *testing.T) {
		// prev <-> current ----> prev
		l := NewList()

		l.PushBack(10)
		l.Remove(l.PushBack(20))

		require.Equal(t, 1, l.Len())
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, l.Front(), l.Back())
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Prev)
	})

	t.Run("remove middle with len = 3", func(t *testing.T) {
		// prev <-> current <-> next ----> prev <=> next
		l := NewList()
		front := l.PushBack(10)
		middle := l.PushBack(20)
		back := l.PushBack(30)
		l.Remove(middle)

		require.Equal(t, 2, l.Len())
		require.Equal(t, front.Value, l.Front().Value)
		require.Equal(t, back.Value, l.Back().Value)
		require.Equal(t, l.Front().Next, l.Back())
		require.Equal(t, l.Front(), l.Back().Prev)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Next)
	})

	t.Run("remove front with len = n", func(t *testing.T) {
		// current(first) <-> next <-> next.next <-> ... ----> next <=> next.next <-> ...
		l := NewList()
		front := l.PushBack(10)
		middle := l.PushBack(20)
		back := l.PushBack(30)
		l.Remove(front)

		require.Equal(t, 2, l.Len())
		require.Equal(t, middle.Value, l.Front().Value)
		require.Equal(t, back.Value, l.Back().Value)
		require.Equal(t, l.Front().Next, l.Back())
		require.Equal(t, l.Front(), l.Back().Prev)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Next)
	})

	t.Run("remove middle with len = n", func(t *testing.T) {
		// ... <-> prev <-> current <-> next <-> ... ----> ... <-> prev <=> next <-> ...
		l := NewList()
		l.PushBack(10)
		l.PushBack(20)
		middle := l.PushBack(30)
		l.PushBack(40)
		l.PushBack(50)
		l.Remove(middle)

		require.Equal(t, 4, l.Len())
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Next)
		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{10, 20, 40, 50}, elems)
	})

	t.Run("remove back with len = n", func(t *testing.T) {
		// ... <-> prev.prev <-> prev <-> current(first) ----> ... <-> prev.prev <=> prev
		l := NewList()
		front := l.PushBack(10)
		middle := l.PushBack(20)
		back := l.PushBack(30)
		l.Remove(back)

		require.Equal(t, 2, l.Len())
		require.Equal(t, front.Value, l.Front().Value)
		require.Equal(t, middle.Value, l.Back().Value)
		require.Equal(t, l.Front().Next, l.Back())
		require.Equal(t, l.Front(), l.Back().Prev)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Next)
	})
}

func TestMoveToFront(t *testing.T) {
	t.Run("move to front back", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.PushBack(20)
		back := l.PushBack(30)
		l.MoveToFront(back)

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{30, 10, 20}, elems)
	})

	t.Run("move to front middle", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		middle := l.PushBack(20)
		l.PushBack(30)
		l.MoveToFront(middle)

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{20, 10, 30}, elems)
	})

	t.Run("move to front front", func(t *testing.T) {
		l := NewList()
		front := l.PushBack(10)
		l.PushBack(20)
		l.PushBack(30)
		l.MoveToFront(front)

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{10, 20, 30}, elems)
	})
}

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

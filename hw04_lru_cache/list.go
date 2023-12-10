package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List
	length    int
	frontItem *ListItem
	backItem  *ListItem
}

func NewList() List {
	return new(list)
}

func (linkedList list) Len() int {
	return linkedList.length
}

func (linkedList list) Front() *ListItem {
	return linkedList.frontItem
}

func (linkedList list) Back() *ListItem {
	return linkedList.backItem
}

func (linkedList *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{
		Value: v,
		Next:  linkedList.frontItem,
		Prev:  nil,
	}
	if linkedList.frontItem != nil {
		linkedList.frontItem.Prev = &newItem
	}
	if linkedList.backItem == nil {
		linkedList.backItem = &newItem
	}
	linkedList.length++
	linkedList.frontItem = &newItem
	return &newItem
}

func (linkedList *list) PushBack(v interface{}) *ListItem {
	newItem := ListItem{
		Value: v,
		Next:  nil,
		Prev:  linkedList.backItem,
	}
	if linkedList.backItem != nil {
		linkedList.backItem.Next = &newItem
	}
	if linkedList.frontItem == nil {
		linkedList.frontItem = &newItem
	}
	linkedList.length++
	linkedList.backItem = &newItem
	return &newItem
}

func (linkedList *list) Remove(i *ListItem) {
	/*
		1. current ----> *empty list*
		backItem = nil;  frontItem = nil;

		2. current <-> next ----> next
		backItem = next;  frontItem = next; next.previous = current.previous

		3. prev <-> current ----> prev
		backItem = prev;  frontItem = prev; prev.next = current.next;

		4. prev <-> current <-> next ----> prev <=> next
		prev.next = current.next; next.previous = current.previous

		5. current(first) <-> next <-> next.next <-> ... ----> next <=> next.next <-> ...
		next.previous = current.previous; frontItem = next;

		6. ... <-> prev.prev <-> prev <-> current(first) ----> ... <-> prev.prev <=> prev
		prev.next = current.next; backItem = next;

		7. ... <-> prev <-> current <-> next <-> ... ----> ... <-> prev <=> next <-> ...
		prev.next = current.next; next.previous = current.previous
	*/
	previousItem := i.Prev
	nextItem := i.Next

	if nextItem != nil {
		// not back
		nextItem.Prev = previousItem
	} else {
		// back
		linkedList.backItem = previousItem
	}
	if previousItem != nil {
		// not front
		previousItem.Next = nextItem
	} else {
		// front
		linkedList.frontItem = nextItem
	}
	linkedList.length--
	// clear allocate memory
	*i = ListItem{}
}

func (linkedList *list) MoveToFront(i *ListItem) {
	// Lazy solution
	// ToDo: change prev and next pointers without reallocation for improve performance
	value := i.Value
	linkedList.Remove(i)
	linkedList.PushFront(value)
}

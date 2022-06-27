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

func (lst *list) Len() int {
	return lst.len
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (lst *list) Front() *ListItem {
	return lst.front
}

func (lst *list) Back() *ListItem {
	return lst.back
}

func (lst *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v, Next: lst.front}

	if lst.front == nil {
		lst.back = item
	} else {
		lst.front.Prev = item
	}
	lst.front = item
	lst.len++
	return item
}

func (lst *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v, Prev: lst.back}

	if lst.back == nil {
		lst.back = item
	} else {
		lst.back.Next = item
	}
	lst.back = item
	lst.len++
	return item
}

func (lst *list) Remove(i *ListItem) {
	lst.gap(i)
	lst.len--
}

func (lst *list) MoveToFront(i *ListItem) {
	if lst.front == i {
		return
	}
	if lst.back == i {
		lst.back = i.Prev
		lst.back.Next = nil
	} else {
		lst.gap(i)
	}
	currFront := lst.front

	lst.front = i
	lst.front.Prev = nil
	lst.front.Next = currFront
	lst.front.Next.Prev = i
}

// Функция для разрыва и сшивания новых связей

func (lst *list) gap(item *ListItem) {
	if item.Next != nil {
		item.Next.Prev = item.Prev
	} else {
		lst.back = item.Prev
	}

	if item.Prev != nil {
		item.Prev.Next = item.Next
	} else {
		lst.front = item.Next
	}
}

func NewList() List {
	return new(list)
}

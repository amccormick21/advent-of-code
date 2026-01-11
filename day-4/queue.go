// Need to implement a double-ended queue
// This allows us to track the number of accessible rolls on a rolling basis
package main

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Queue[T Numeric] interface {
	PushBack(item T)
	PopFront() T
	Sum() T
}

type Deque[T Numeric] struct {
	items []T
}

func NewDeque[T Numeric]() *Deque[T] {
	return &Deque[T]{items: []T{}}
}

func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

func (d *Deque[T]) PopFront() T {
	if len(d.items) == 0 {
		panic("PopFront from empty deque")
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item
}

func (d *Deque[T]) Sum() T {
	var sum T
	for _, item := range d.items {
		sum += item
	}
	return sum
}

type FixedCapacityQueue[T Numeric] struct {
	*Deque[T]
	capacity int
}

func NewFixedCapacityQueue[T Numeric](capacity int) *FixedCapacityQueue[T] {
	return &FixedCapacityQueue[T]{
		Deque:    NewDeque[T](),
		capacity: capacity,
	}
}

func (fcq *FixedCapacityQueue[T]) PushBack(item T) T {
	var delta T
	if len(fcq.items) == fcq.capacity {
		delta -= fcq.PopFront()
	}
	fcq.Deque.PushBack(item)
	return delta + item
}

package generic

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Push(newItem T) {
	q.items = append(q.items, newItem)
}

func (q *Queue[T]) Pop() (T, error) {
	if q.Empty() {
		var zero T
		return zero, nil
	}
	lastIndex := len(q.items) - 1
	lastItem := q.items[lastIndex]
	q.items = q.items[:lastIndex]
	return lastItem, nil
}

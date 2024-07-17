package generic

import "testing"

func TestQueue(t *testing.T) {
	t.Run("queue functions", func(t *testing.T) {
		q := new(Queue[int])

		AssertTrue(t, q.Empty())
		AssertEqual(t, q.Size(), 0)

		q.Push(5)
		AssertFalse(t, q.Empty())
		AssertEqual(t, q.Size(), 1)

		val, err := q.Pop()
		AssertEqual(t, err, nil)
		AssertEqual(t, val, 5)
	})
}
func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "世界", "世界")
		AssertNotEqual(t, "世界", "World")
	})
}

func AssertTrue(t *testing.T, value bool) {
	t.Helper()
	if value != true {
		t.Error("expected true")
	}
}
func AssertFalse(t *testing.T, value bool) {
	t.Helper()
	if value != false {
		t.Error("expected false")
	}
}

func AssertNotEqual[T comparable](t *testing.T, i1, i2 T) {
	t.Helper()
	if i1 == i2 {
		t.Errorf("%+v and %+v equal", i1, i2)
	}
}

func AssertEqual[T comparable](t *testing.T, i1, i2 T) {
	t.Helper()
	if i1 != i2 {
		t.Errorf("%+v and %+v not equal", i1, i2)
	}
}

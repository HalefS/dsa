package queue

import "testing"

func TestPeek(t *testing.T) {
	values := []int{10, 20, 30}
	q := prepareQueue(t, values)
	if *q.Peek() != values[0] {
		t.Errorf("expected peek value to be %d, got %d", values[0], *q.Peek())
		return
	}

	q = prepareQueue(t, []int{})
	if q.Peek() != nil {
		t.Errorf("expected peek on empty queue to be nil")
	}
}

func TestEnqueue(t *testing.T) {
	q := New[int]()

	want := 20
	q.Enqueue(want)
	if *q.Peek() != want {
		t.Errorf("expected q.Peek to be %d, got %d", want, *q.Peek())
		return
	}

	q.Enqueue(50)
	if *q.Peek() != want {
		t.Errorf("expected q.Peek to be %d, got %d", want, *q.Peek())
	}
}

func TestDequeue(t *testing.T) {
	values := []int{30, 60, 50, 30}
	q := prepareQueue(t, values)

	got := *q.Dequeue()
	if got != values[0] {
		t.Errorf("expected q.dequeue to be %d, got %d", values[0], got)
		return
	}

	q = prepareQueue(t, []int{})
	if q.Dequeue() != nil {
		t.Errorf("expected dequeue on empty queue to be nil")
	}
}

func prepareQueue(t *testing.T, values []int) *Queue[int] {
	t.Helper()

	q := New[int]()
	for _, v := range values {
		q.Enqueue(v)
	}

	return q
}

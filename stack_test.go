package stack

import "testing"

func TestPush(t *testing.T) {
	var stack Stack

	l := len(stack)
	if l != 0 {
		t.Error("Empty stack has non-zero length: ", l)
	}

	for i := 1; i <= 100; i++ {
		stack.Push(i)
		l = len(stack)

		if l != i {
			t.Errorf("Expected stack of length %d, got length %d\n", i, l)
		}
	}
}

func TestLength(t *testing.T) {
	var stack Stack

	l := stack.Length()
	if l != 0 {
		t.Error("Expected length of 0, got ", l)
	}

	for i := 1; i <= 100; i++ {
		stack.Push(i)
		l = stack.Length()
		if l != i {
			t.Errorf("Expected length %d, got length %d\n", i, l)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var stack Stack

	if !stack.IsEmpty() {
		t.Error("Empty stack but IsEmpty returned true")
	}

	stack.Push("a")
	if stack.IsEmpty() {
		t.Error("Non-empty stack but IsEmpty returned false")
	}
}

func TestPop(t *testing.T) {
	var stack Stack

	val, err := stack.Pop()
	if val != nil {
		t.Error("Non-nil value popped from an empty stack")
	}
	if err == nil {
		t.Error("Error returned from popping empty stack should not be nil")
	}

	stack.Push(1)
	val, err = stack.Pop()
	if val.(int) != 1 {
		t.Errorf("Wrong value returned by stack.Pop: expected 1, got %d\n",
			val)
	}
	if err != nil {
		t.Error("Non-nil error returned by stack.Pop: ", err)
	}
	if stack.Length() != 0 {
		t.Error("Value not removed from stack when popped")
	}
}

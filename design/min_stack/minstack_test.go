package minstack

import "testing"

func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	if min := stack.GetMin(); -3 != min {
		t.Fatalf("expected minimum to be -3, got %d", min)
	}
	stack.Pop()
	if top := stack.Top(); 0 != top {
		t.Fatalf("expected top to be 0, got %d", top)
	}
	if min := stack.GetMin(); -2 != min {
		t.Fatalf("expected minimum to be -2, got %d", min)
	}
}

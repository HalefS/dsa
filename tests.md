> can you write unit tests for this?
> 
> ```go
> func (s *Stack[T]) Push(value T) {
> 	node := &Node[T]{Value: value}
> 	s.size++
> 
> 	if s.head == nil {
> 		s.head = node
> 		return
> 	}
> 
> 	node.Prev = s.head
> 	s.head = node
> }
> ```

{{{ 1 context files
- /Users/halefspencer/code/dsa/minheap.go
}}}

Here is how I would write unit tests for the Push method on a Stack:

```go
func TestStack_Push(t *testing.T) {

  t.Run("it increases the size", func(t *testing.T) {
    s := NewStack[int]()

    if s.Size() != 0 {
      t.Errorf("new stack should have size 0")
    }

    s.Push(1)

    if s.Size() != 1 {
      t.Errorf("after push, size should be 1")
    }
  })

  t.Run("it sets head to pushed node", func(t *testing.T) {
    s := NewStack[int]()
    n1 := &Node[int]{Value: 1}

    s.Push(1)

    if s.head != n1 {
      t.Errorf("head should be pushed node")
    }
  })

  t.Run("it links previous head", func(t *testing.T) {
    s := NewStack[int]()
    n1 := &Node[int]{Value: 1}
    n2 := &Node[int]{Value: 2}

    s.Push(1)
    s.Push(2)

    if n2.Prev != n1 {
      t.Errorf("new head should link to previous")
    }
  })

}
```

This covers the key behaviors:

- Size increments
- Head set to new node
- New node links previous head

Let me know if you need any clarification or have additional functions to test!

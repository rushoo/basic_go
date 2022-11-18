package nodestack

type Node[T any] struct {
	value T
	next  *Node[T]
}
type Stack[T any] struct {
	first *Node[T]
}

func (stack *Stack[T]) Push(item T) {
	// 数据入栈，先将数据实例化为node，将原来的数据链放在此node之后，再将此node命名为stack.first
	// 这样新的数据就插入到（成为了）stack的第一个节点
	newNode := Node[T]{item, nil}
	newNode.next = stack.first //node.next指向原来的数据链
	stack.first = &newNode     //将node为首的新链挂到stack.first
}
func (stack *Stack[T]) Top() T {
	return stack.first.value
}
func (stack *Stack[T]) Pop() T {
	result := stack.first.value
	stack.first = stack.first.next //直接跳过第一个节点，相当于删除，即为出栈
	return result
}
func (stack Stack[T]) IsEmpty() bool {
	return stack.first == nil
}

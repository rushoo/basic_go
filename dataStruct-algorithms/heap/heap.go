package main

type Ordered interface {
	~float64 | ~int | ~string
}
type Heap[T Ordered] struct {
	Items []T
}

func (heap *Heap[T]) swap(index1, index2 int) {
	heap.Items[index1], heap.Items[index2] = heap.Items[index2], heap.Items[index1]
}
func (heap *Heap[T]) buildHeap(index int) {
	var parent int
	if index > 0 {
		//当新插入的元素大于其父节点元素时，堆结构被破坏，须重新调整结构
		parent = (index - 1) / 2
		if heap.Items[index] > heap.Items[parent] {
			heap.swap(index, parent)
		}
		heap.buildHeap(parent)
	}
}
func (heap *Heap[T]) rebuildHeap(index int) {
	//	将index位置元素与后面可能的两个元素交换最大最小
	length := len(heap.Items)
	// 当前非最后一层
	if 2*index+1 < length {
		left := 2*index + 1
		right := 2*index + 2
		if 2*index+2 == length { //左子节点最后元素了
			if heap.Items[left] > heap.Items[index] {
				heap.swap(index, left)
			}
		} else {
			max := heap.maxItem(index, left, right)
			if max > index {
				heap.swap(index, max)
				heap.rebuildHeap(max)
			}
		}
	}
}
func (heap *Heap[T]) maxItem(a, b, c int) int {
	//获取三个元素中最大值值元素的位置，逻辑是：
	//如果a>=b&&a>=c就认为a最大，b>=a&&b>=c就认为b最大，否则认为c最大
	switch {
	case heap.Items[a] >= heap.Items[b] && heap.Items[a] >= heap.Items[c]:
		return a
	case heap.Items[b] >= heap.Items[a] && heap.Items[b] >= heap.Items[c]:
		return b
	default:
		return c
	}
}

func NewHeap[T Ordered](input []T) *Heap[T] {
	heap := &Heap[T]{}
	for i := 0; i < len(input); i++ {
		heap.Insert(input[i])
	}
	return heap
}
func (heap *Heap[T]) Insert(value T) {
	//先将元素插入，再调整堆结构
	heap.Items = append(heap.Items, value)
	heap.buildHeap(len(heap.Items) - 1)
}
func (heap *Heap[T]) Remove() {
	//删除最大值,将最后一个元素的值给到第一个元素，删除掉最后一个元素
	//再根据堆的性质调整各位置元素的值
	length := len(heap.Items)
	heap.Items[0] = heap.Items[length-1]
	heap.Items = heap.Items[:length-1]
	heap.rebuildHeap(0)
}
func (heap *Heap[T]) Largest() T {
	return heap.Items[0]
}

package main

type DataType interface {
	~string | ~int | ~float64
}

// 一堆键值对的合集
type Set[T DataType] struct {
	items map[T]bool
}

func (set *Set[T]) Insert(item T) {
	if set.items == nil {
		set.items = make(map[T]bool)
	}
	_, ok := set.items[item]
	// 空时插入
	if !ok {
		set.items[item] = true
	}
}
func (set *Set[T]) Delete(item T) {
	_, ok := set.items[item]
	if ok {
		//内置函数，从map中删除对应的key值
		delete(set.items, item)
	}
}
func (set *Set[T]) In(item T) bool {
	_, ok := set.items[item]
	return ok
}
func (set *Set[T]) Items() []T {
	items := []T{}
	for item := range set.items {
		items = append(items, item)
	}
	return items
}
func (set *Set[T]) Size() int {
	return len(set.items)
}

// 合并两个set
func (set *Set[T]) Union(set2 Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for index := range set.items {
		result.items[index] = true
	}
	for j := range set2.items {
		_, present := result.items[j]
		if !present {
			result.items[j] = true
		}
	}
	return &result
}

// 找到两个set共同的子集
func (set *Set[T]) Intersection(set2 Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	//遍历获取set2的key值，如果这个值在set中也存在，就将它加入到result
	for i := range set2.items {
		_, present := set.items[i]
		if present {
			result.items[i] = true
		}
	}
	return &result
}

// 找到两个set的差集set-set2
func (set *Set[T]) Difference(set2 Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for i := range set.items {
		_, present := set2.items[i]
		if !present {
			result.items[i] = true
		}
	}
	return &result
}

// 判断是否子集
func (set *Set[T]) Subset(set2 Set[T]) bool {
	for i := range set.items {
		_, present := set2.items[i]
		if !present {
			return false
		}
	}
	return true
}

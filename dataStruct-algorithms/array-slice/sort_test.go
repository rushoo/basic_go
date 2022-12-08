package main

import (
	"math/rand"
	"testing"
	"time"
)

func genArray[T dataType](len int) []T {
	data := make([]T, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		item := rand.Intn(len) - len/2
		data[i] = T(item)
	}
	return data
}
func isSorted[T dataType](data []T) bool {
	end := len(data)
	for i := 0; i < end-1; i++ {
		if data[i] > data[i+1] {
			return false
		}
	}
	return true
}

func TestQuickSort(t *testing.T) {
	data := genArray[int32](10)
	QuickSort[int32](data)
	if res := isSorted[int32](data); res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}
func TestMergeSort(t *testing.T) {
	data := genArray[int64](10)
	data = MergeSort[int64](data)
	if res := isSorted[int64](data); res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}
func TestBubbleSort(t *testing.T) {
	data := genArray[float32](200)
	BubbleSort[float32](data)
	if res := isSorted[float32](data); res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}
func TestInsertSort(t *testing.T) {
	data := genArray[float64](200)
	InsertSort[float64](data)
	if res := isSorted[float64](data); res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

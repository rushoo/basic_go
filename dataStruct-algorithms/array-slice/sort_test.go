package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func genArray(len int) []int {
	items := make([]int, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		item := rand.Intn(len) - len/2
		items[i] = item
	}
	return items
}
func isSorted(items []int) bool {
	end := len(items)
	for i := 0; i < end-1; i++ {
		if items[i] > items[i+1] {
			return false
		}
	}
	return true
}

func TestMergeSort(t *testing.T) {
	items := genArray(10)
	items = MergeSort(items)
	fmt.Println(items)
	if res := isSorted(items); res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}
func TestBubbleSort(t *testing.T) {
	items := genArray(200)
	BubbleSort(items)
	if res := isSorted(items); res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}
func TestInsertSort(t *testing.T) {
	items := genArray(200)
	InsertSort(items)
	if res := isSorted(items); res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

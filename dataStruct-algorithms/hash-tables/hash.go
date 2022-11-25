package main

import (
	"hash/fnv"
	"strings"
)

const (
	tableSize = 100_000
	Radix     = uint64(10)
	Q         = uint64(10 ^ 9 + 9)
)

var length int

// word相当于逻辑上的缓存，插入查询时候先看这里，否则再处理内层[]list
type WordType struct {
	word string
	list []string
}

// 定义一个[]WordType数组,长度为100000的类型别名HashTable
type HashTable [tableSize]WordType

func hash(s string) uint32 {
	h := fnv.New32a() // Fowler-Noll-Vo algorithm,New32a returns a new 32-bit FNV-1a hash.Hash.
	h.Write([]byte(s))
	return h.Sum32()
}

func NewTable() HashTable {
	var table HashTable
	for i := 0; i < tableSize; i++ {
		table[i] = WordType{"", []string{}}
	}
	return table
}
func (table *HashTable) Insert(word string) {
	index := hash(word) % tableSize
	if table[index].word == word {
		//数据已存在，无操作返回
		return
	}
	if len(table[index].list) > 0 {
		for _, v := range table[index].list {
			if v == word {
				return //数据已存在，无操作返回
			}
		}
	}
	// 优先将数据插入word项，非空则加到list里
	if table[index].word == "" {
		table[index].word = word
	} else {
		table[index].list = append(table[index].list, word)
	}
	length++
}

// 查询word是否存在hashtable中
func (table *HashTable) IsPresent(word string) bool {
	//根据word的hash值找到对应的index
	index := hash(word) % tableSize
	if table[index].word == word {
		return true
	}
	for _, v := range table[index].list {
		if v == word {
			return true
		}
	}
	return false
}

// 暴力搜索文本中是否出现给定字符串
func bruteSearch(txt, pattern string) (bool, int) {
	length := len(pattern)
	for i := 0; i < len(txt)-length; i++ {
		if txt[i:i+length] == pattern {
			return true, i
		}
	}
	return false, -1
}

func Hash(s string, Length int) uint64 {
	// Horner's method
	h := uint64(0)
	for i := 0; i < Length; i++ {
		h = (h*Radix + uint64(s[i])) % Q
	}
	return h
}
func Search(txt, pattern string) (bool, int) {
	strings.ToLower(txt)
	strings.ToLower(pattern)
	n := len(txt)
	m := len(pattern)
	patternHash := Hash(pattern, m)
	textHash := Hash(txt, m)
	if textHash == patternHash {
		return true, 0
	}
	PM := uint64(1)
	for i := 1; i <= m-1; i++ {
		PM = (Radix * PM) % Q
	}
	for i := m; i < n; i++ {
		textHash = (textHash + Q - PM*uint64(txt[i-m])%Q) % Q
		textHash = (textHash*Radix + uint64(txt[i])) % Q
		if (patternHash == textHash) && pattern == txt[(i-m+1):(i+1)] {
			return true, i - m + 1
		}
	}
	return false, -1
}

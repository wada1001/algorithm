package mmap

import (
	"errors"
	"math"
)

type HashMap struct {
	arr []*HashSet
	length int
}

type HashSet struct {
	key string
	val string
}

func MakeHashMap() *HashMap {
	return &HashMap{arr: make([]*HashSet, 500), length: 500}
}

// Cannot use...
func MakePoorHash(str string, max int) int {
	hash := 0
	for _, v := range str {
		hash += int(v)
	}
	return hash % max
}

func MakeHash(str string, max int) int {
	hash := 0
	for i, v := range str {
		hash += int(v) * int(math.Pow(256, float64(i % 8)))
	}
	return hash % max
}

func (h *HashMap) Set(key, val string) error {
	hash := MakeHash(key, h.length)
	if h.arr[hash] != nil && h.arr[hash].key != key {
		hash = h.ReHash(hash, key)
		if hash < 0 {
			return errors.New("cannot set")
		}
	}
	h.arr[hash] = &HashSet{key: key, val: val}
	return nil
}

func (h *HashMap) Get(key string) (string, error) {
	hash := MakeHash(key, h.length)
	if h.arr[hash] == nil {
		return "", errors.New("not set")
	}
	set := h.arr[hash]
	i := 1
	for set.key != key {
		hash += int(math.Pow(float64(i), float64(i)))
		set = h.arr[hash]
		i++
		if int(math.Pow(float64(i), float64(i))) > h.length {
			return "", errors.New("not set")
		}
	}
	
	return h.arr[hash].val, nil
}

// move set position. hash + i * i
func (h *HashMap) ReHash(first int, key string) int {
	i := 1
	tmp := 0
	for h.arr[first] != nil && h.arr[first].key != key {
		pow := int(math.Pow(float64(i), float64(i)))

		tmp = first + pow
		if h.arr[tmp] == nil {
			return tmp
		}

		if h.arr[tmp].key == key {
			return tmp
		}

		i++
		if pow > h.length / 2{
			break
		}
	}
	return -1
}


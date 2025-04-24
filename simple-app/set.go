package main

import "fmt"

type Set[T comparable] map[T]bool

func (s Set[T]) Add(v T) {
	s[v] = true
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Delete(v T) {
	delete(s, v)
}

func (s Set[T]) String() string {
	res := "Set["

	for k := range s {
		res += fmt.Sprint(k) + " "
	}
	return res + "]"
}

package main

import (
	"golang.org/x/exp/slices"
	"reflect"
)

// type Set interface {
// 	Insert(any)
// 	Contains(any) bool
// 	// Equal(Set) bool
// }

type MapSet struct {
	ds map[any]any
}

func (s *MapSet) Insert(elem any) {
	s.ds[elem] = struct{}{}
}

func (s *MapSet) Contains(elem any) bool {
	_, ok := s.ds[elem]
	return ok
}

type SliceSetExp[T comparable] struct {
	ds []T
}

func (s *SliceSetExp[T]) Insert(elem T) {
	s.ds = append(s.ds, elem)
}

func (s *SliceSetExp[T]) Contains(elem T) bool {
	return slices.Contains(s.ds, elem)
}

type ReflectSliceSet struct {
	ds []any
}

func (s *ReflectSliceSet) Insert(elem any) {
	s.ds = append(s.ds, elem)
}

func (s *ReflectSliceSet) Contains(elem any) bool {
	for _, x := range s.ds {
		if reflect.DeepEqual(x, elem) {
			return true
		}
	}
	return false
}

func main() {
}

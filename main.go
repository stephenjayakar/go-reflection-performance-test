package main

import (
	"golang.org/x/exp/slices"
)

type MapSet struct {
	ds map[string]struct{}
}

func NewMapSet() MapSet {
	return MapSet{
		ds: map[string]struct{}{},
	}
}

func (s *MapSet) Insert(elem string) {
	s.ds[elem] = struct{}{}
}

func (s *MapSet) Contains(elem string) bool {
	_, ok := s.ds[elem]
	return ok
}

type SliceSetExp struct {
	ds []string
}

func NewSliceSetExp() SliceSetExp {
	return SliceSetExp{
		ds: []string{},
	}
}

func (s *SliceSetExp) Insert(elem string) {
	s.ds = append(s.ds, elem)
}

func (s *SliceSetExp) Contains(elem string) bool {
	return slices.Contains(s.ds, elem)
}

type SliceSet struct {
	ds []string
}

func NewSliceSet() SliceSet {
	return SliceSet{
		ds: []string{},
	}
}

func (s *SliceSet) Insert(elem string) {
	s.ds = append(s.ds, elem)
}

func (s *SliceSet) Contains(elem string) bool {
	for _, x := range s.ds {
		if x == elem {
			return true
		}
	}
	return false
}

func main() {
}

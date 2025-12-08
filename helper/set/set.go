package set

import "maps"

type Set[T comparable] struct {
	inner map[T]any
}

func NewEmptySet[T comparable]() *Set[T] {
	return &Set[T]{inner: make(map[T]any)}
}

func NewSet[T comparable](init ...T) *Set[T] {
	res := NewEmptySet[T]()
	for _, v := range init {
		res.inner[v] = nil
	}
	return res
}

func (set *Set[T]) Add(value T) {
	set.inner[value] = nil
}

func (set *Set[T]) Added(value T) *Set[T] {
	added := NewEmptySet[T]()
	maps.Copy(added.inner, set.inner)
	added.Add(value)
	return added
}

func (set *Set[T]) Remove(value T) {
	delete(set.inner, value)
}

func (set *Set[T]) Removed(value T) *Set[T] {
	removed := NewEmptySet[T]()
	maps.Copy(removed.inner, set.inner)
	removed.Remove(value)
	return removed
}

func (set *Set[T]) Contains(value T) bool {
	_, ok := set.inner[value]
	return ok
}

func (set *Set[T]) Size() int {
	return len(set.inner)
}

func (set *Set[T]) Clear() {
	set.inner = make(map[T]any)
}

func (set *Set[T]) Union(other *Set[T]) {
	for k := range other.inner {
		set.inner[k] = other.inner[k]
	}
}

func (set *Set[T]) United(other *Set[T]) *Set[T] {
	united := NewEmptySet[T]()
	united.Union(other)
	united.Union(set)
	return united
}

func (set *Set[T]) Intersect(other *Set[T]) {
	for k := range set.inner {
		if !other.Contains(k) {
			set.Remove(k)
		}
	}
}

func (set *Set[T]) Intersected(other *Set[T]) *Set[T] {
	intersected := NewEmptySet[T]()
	for k := range set.inner {
		if other.Contains(k) {
			intersected.Add(k)
		}
	}
	return intersected
}

func (set *Set[T]) Difference(other *Set[T]) {
	for k := range other.inner {
		set.Remove(k)
	}
}

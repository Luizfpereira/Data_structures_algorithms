package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) AddElement(element int) {
	if !set.containsElement(element) {
		set.integerMap[element] = true
	}
}

func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

func (set *Set) containsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]
	return exists
}

func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	var value int
	for value = range set.integerMap {
		if anotherSet.containsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	var value int
	for value = range set.integerMap {
		unionSet.AddElement(value)
	}
	for value = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}

func main() {
	set := &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set.containsElement(1))
	var anotherSet *Set
	anotherSet = &Set{}
	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)
	fmt.Println("Set: ", set)
	fmt.Println("Another Set: ", anotherSet)
	fmt.Println("Intersect:", set.Intersect(anotherSet))
	fmt.Println("Union: ", set.Union(anotherSet))
}

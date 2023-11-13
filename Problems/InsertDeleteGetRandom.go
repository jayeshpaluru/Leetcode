package main

// Implement the RandomizedSet class:
// RandomizedSet() Initializes the RandomizedSet object.
// bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
// bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
// int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called).
// Each element must have the same probability of being returned.
// You must implement the functions of the class such that each function works in average O(1) time complexity.

// RandomizedSet is a struct with a map and an array
import "math/rand"

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		a: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.a)
	this.a = append(this.a, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	i := this.m[val]
	this.a[i] = this.a[len(this.a)-1]
	this.m[this.a[i]] = i
	this.a = this.a[:len(this.a)-1]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.a[rand.Intn(len(this.a))]
}

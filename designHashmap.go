package main

type MyHashMap struct {
	value []*Node
}

type Node struct {
	key   int
	value int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]*Node, 0)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	for _, node := range this.value {
		if node.key == key {
			node.value = value
			return
		}
	}
	this.value = append(this.value, &Node{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	for _, node := range this.value {
		if node.key == key {
			return node.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	for i, node := range this.value {
		if node.key == key {
			this.value = append(this.value[:i], this.value[i+1:]...)
			return
		}
	}
}

/**
* Your MyHashMap object will be instantiated and called as such:
* obj := Constructor();
* obj.Put(key,value);
* param_2 := obj.Get(key);
* obj.Remove(key);
 */

package main

import "fmt"

type LRUCache struct {
	m         map[int]int
	um        map[int]int
	capacity  int
	uCapacity int
	u         int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:        make(map[int]int, capacity),
		um:       make(map[int]int, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("BEFORE Get, key:", key, " this:", this)

	if value, exist := this.m[key]; !exist {
		fmt.Println("AFTER Get, key:", key, " this:", this)

		return -1
	} else {
		this.u = this.u + 1
		this.um[key] = this.u

		fmt.Println("AFTER Get, key:", key, " this:", this)
		return value
	}

}

func (this *LRUCache) Put(key int, value int) {
	fmt.Println("BEFORE Put, key:", key, " value: ", value, " this: ", this)
			if _, exist := this.m[key]; exist {
			this.m[key] = value
			//this.u = this.u + 1
			this.um[key] = this.u

			return
		}

	if this.uCapacity >= this.capacity {

		var minV int
		var minK int
		var i = 0
		for k, v := range this.um {
			if i == 0 {
				minK = k
				minV = v
			} else {
				if v < minV {
					minK = k
					minV = v
				}
			}

			i = i + 1
		}

		delete(this.m, minK)
		delete(this.um, minK)
		this.m[key] = value
		//this.u = this.u + 1
		this.um[key] = this.u

		fmt.Println("AFTER Put, key:", key, " value: ", value, " this: ", this)
		return
	}

	//this.u = this.u + 1
	this.m[key] = value
	this.um[key] = this.u
	this.uCapacity = this.uCapacity + 1

	fmt.Println("AFTER Put, key:", key, " value: ", value, " this: ", this)
}

func main() {

	lru := Constructor(2)
	lru.Put(1, 1) // cache is {1=1}
	fmt.Println()
	lru.Put(2, 2) // cache is {1=1, 2=2}
	fmt.Println()
	fmt.Println(lru.Get(1)) // return 1
	fmt.Println()
	lru.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println()
	fmt.Println(lru.Get(2)) // returns -1 (not found)
	fmt.Println()
	lru.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println()
	fmt.Println(lru.Get(1)) // return -1 (not found)

	fmt.Println(lru.Get(3)) // return 3
	fmt.Println(lru.Get(4)) // return 4

}

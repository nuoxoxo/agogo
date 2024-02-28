package main

import "fmt"

/************ PUBLIC ************/

type Heapq struct {
	Arr []int // actually we use a slice, not array
}

func Heapify(Slice []int) (hq *Heapq) {
	res := &Heapq{}
	for _, n := range Slice {
		res.Heappush(n)
	}
	return res
}

func (hq *Heapq) Heappush(key int) {
	hq.Arr = append(hq.Arr, key)
	hq.heapifyUp(len(hq.Arr) - 1)
}

// pop the largest key
func (hq *Heapq) Heappop() int {
	top := hq.Arr[0] // val?
	N := len(hq.Arr)
	if N == 0 {
		panic("popped an empty hq")
	}
	hq.Arr[0] = hq.Arr[N-1]
	hq.Arr = hq.Arr[:N-1]
	hq.heapifyDown(0)
	return top
}

func (hq *Heapq) Size() int {
	return len(hq.Arr)
}

/************ PRIVATE ************/

func (hq *Heapq) heapifyUp(index int) {
	// heapify from ground up
	for hq.Arr[parentIndex(index)] < hq.Arr[index] {
		hq.swap(parentIndex(index), index)
		index = parentIndex(index)
	}
}

func (hq *Heapq) heapifyDown(index int) {
	// heapify top to bottom
	N := len(hq.Arr)
	L := left(index)
	R := right(index)

	// cond. 1 - left child is the only child
	// cond. 2 - left child is larger
	for L <= N-1 {
		var to_compare int
		if L == N-1 || hq.Arr[L] > hq.Arr[R] {
			to_compare = L
		} else {
			// cond. 3 - right child is larger
			to_compare = R
		}
		if hq.Arr[index] < hq.Arr[to_compare] {
			hq.swap(index, to_compare)
			index = to_compare
			L, R = left(index), right(index)
		} else {
			return
		}
	}

}

func parentIndex(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

// swap keys in .Arr
// will modify the Arr
func (hq *Heapq) swap(L, R int) {
	hq.Arr[L], hq.Arr[R] = hq.Arr[R], hq.Arr[L]
}

func main() {
	Slice := []int{1, 3, 5, 7, 9, 15, 30, 42, 99, 1024}
	fmt.Println("/init", Slice, len(Slice))

	// testing two code blocks
	choice := 1

	// block 1 : Heappush 1 item at a time, no Heapify
	if choice == 0 {
		testq := &Heapq{}
		fmt.Println("/no Heapify", testq, testq.Size())
		for _, n := range Slice {
			testq.Heappush(n)
			fmt.Println("/Heappush", testq, testq.Size())
		}
		i := 0
		for i < 5 {
			i++
			testq.Heappop()
			fmt.Println("/Pop", testq)
		}
		// block 2 : using Heapify
	} else {
		testq := Heapify(Slice)
		fmt.Println("/w. Heapify", testq, testq.Size())
		i := 0
		for i < 5 {
			i++
			testq.Heappop()
			fmt.Println("/Pop", testq, testq.Size())
		}
	}
}

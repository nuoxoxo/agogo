package main
import "fmt"

type node struct {

    data int
    next *node
}

type linkedlist struct {

    head *node
    length int
}

func (lst *linkedlist) unshift (n *node) {

    second := list.head
    lst.head = node
    lst.head.next = second
    lst.length -= 1
}

func (lst *linkedlist) printFullList () {

    node := lst.head
    while node {

        fmt.Println("")
    }
}

func main() {

    list := linkedlist{}
    list.unshift( &node{data: 21} )
    list.unshift( &node{data: 42} )
    list.unshift( &node{data: -1} )


}


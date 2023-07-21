package main

import (
    "fmt"
    "_colors"
)

type Node struct {

    data int
    next *Node
}

type Linkedlist struct {

    head *Node
    length int
}

func (lst *Linkedlist) unshift (node *Node) {

    second := lst.head
    lst.head = node
    lst.head.next = second
    lst.length -= 1
}

func (lst *Linkedlist) printFullList () {

    fmt.Println(Cyan + "\n(printer : begin)\n" + Reset)

    node := lst.head
    n := 0
    for node != nil {

        fmt.Printf("node %d : %v \n", n, node.data)
        node = node.next
        n += 1
    }

    fmt.Println(Cyan + "\n(printer : end)\n" + Reset)
}

func main() {

    {

        list := Linkedlist{}

        fmt.Println(list)

        list.unshift( &Node{data: 42} )
        list.unshift( &Node{data: 21} )
        list.unshift( &Node{data: 0} )
        list.unshift( &Node{data: -1} )

        fmt.Println(list)

        list.printFullList()
    }
    {

        list := Linkedlist{}

        fmt.Println(list)

        n1 := &Node{data: 42}
        n2 := &Node{data: 21}
        n3 := &Node{data: 0}
        n4 := &Node{data: -1}

        list.unshift( n1 )
        list.unshift( n2 )
        list.unshift( n3 )
        list.unshift( n4 )


        fmt.Println(list)

        list.printFullList()
    }
}


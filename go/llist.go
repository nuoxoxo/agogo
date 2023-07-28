package main

import (
    "fmt"
)

const Cyan = "\033[0;36m"
const Reset = "\033[0m"

type Node struct {

    Data int
    Next *Node
}

type Linkedlist struct {

    Head *Node
    Length int
}

// func (lst *Linkedlist) deleteValue () {

func (lst *Linkedlist) unshift (node *Node) {

    second := lst.Head
    lst.Head = node
    lst.Head.Next = second
    // lst.Length -= 1 // buggy
    lst.Length += 1
}

func (lst *Linkedlist) print () {

    fmt.Printf(Cyan + "\n(printer starts, Length = %v)\n" + Reset, lst.Length)

    // fmt.Printf(" - Length: %d \n", lst.Length)
    node := lst.Head
    n := 0
    for node != nil {

        fmt.Printf("node %d : %v \n", n, node.Data)
        node = node.Next
        n += 1
    }

    fmt.Println(Cyan + "\n(printer ends)\n" + Reset)
}

func main() {

    {

        list := Linkedlist{}

        fmt.Println(list)

        list.unshift( &Node{Data: 42} )
        list.unshift( &Node{Data: 21} )
        list.unshift( &Node{Data: 0} )
        list.unshift( &Node{Data: -1} )

        fmt.Println(list)

        list.print()
    }
    {

        list := Linkedlist{}

        fmt.Println(list)

        n1 := &Node{Data: 42}
        n2 := &Node{Data: 21}
        n3 := &Node{Data: 0}
        n4 := &Node{Data: -1}

        list.unshift( n1 )
        list.unshift( n2 )
        list.unshift( n3 )
        list.unshift( n4 )


        fmt.Println(list)

        list.print()
    }
}


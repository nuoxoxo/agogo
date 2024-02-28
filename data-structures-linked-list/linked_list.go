package main

import (
    "fmt"
)

const Cyan = "\033[0;36m"
const Yell = "\033[33m"
const Reset = "\033[0m"

type Node struct {
    Data int
    Next *Node
}

type Linkedlist struct {
    Head *Node
    Length int
}

func (lst *Linkedlist) deleteByValue (targetVal int) {
    if lst.Length == 0 {
        return
    }
    // if target node is the head
    if lst.Head.Data == targetVal {
        lst.Head = lst.Head.Next
        lst.Length -= 1
        return
    }
    // (if no node matches : to handle inside traversal)
    // find the prev node of the target node
    prev := lst.Head
    for prev.Next != nil && prev.Next.Data != targetVal {
        // if not found
        if prev.Next.Next == nil {
            fmt.Println(Yell + "(target to delete not found)" + Reset)
            return
        }
        prev = prev.Next
    }
    prev.Next = prev.Next.Next
    lst.Length -= 1
}

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
        fmt.Println("try deleting the node {42}")
        list.deleteByValue(42)
        fmt.Println(list)
        fmt.Println("try deleting a node that is not there {1024}")
        list.deleteByValue(1024)
        fmt.Println(list)
        list.print()
    }
}

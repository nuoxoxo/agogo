package main
import "fmt"

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

    fmt.Println()

    node := lst.head
    n := 0
    for node != nil {

        fmt.Printf("node %d : %v \n", n, node.data)
        node = node.next
        n += 1
    }
}

func main() {

    list := Linkedlist{}
    list.unshift( &Node{data: 21} )
    list.unshift( &Node{data: 42} )
    list.unshift( &Node{data: -1} )

    list.printFullList()
}


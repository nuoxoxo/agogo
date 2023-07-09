package main

import (
    "fmt"
)

type Graph struct {

    vertices []*Vertex
}

type Vertex struct {

    key int
    adjacent []*Vertex
}

func (g *Graph) setVertex(k int) {
    g.vertices = append(g.vertices, & Vertex{ key: k })
}


//  Todo
//      getter/setter vertex
//      setter edge
//      .contains


func (e *Graph) printGraph() {
    for _, v := range e.vertices {
        fmt.Printf("\nvertex %v : ", v.key)
        for _, v := range v.adjacent {
            fmt.Printf("%v", v.key)
        }
    }
    fmt.Println()
}


//  Drive

func main() {
    e := & Graph{}


    for i := 0; i < 12; i++ {
        e.setVertex(i)
    }

    fmt.Println(e)
    e.printGraph()
}


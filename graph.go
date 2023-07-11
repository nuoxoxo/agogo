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

func (g * Graph) getVertex(k int) *Vertex {
    for i, v := range g.vertices {
        if v.key == k {
            return g.vertices[i]
        }
    }
    return nil
}

func (g *Graph) setVertex(k int) {
    // check if key already exists - .contains
    if contains(g.vertices, k) {
        err := fmt.Errorf("key existed: %v", k)
        fmt.Println(err.Error())
        return
    }
    g.vertices = append(g.vertices, & Vertex{ key: k })
}

func contains(vs []*Vertex, k int) bool {
    for _, v := range vs {
        if k == v.key {
            return true
        }
    }
    return false
}

//  Todo
//      getter/setter vertex
//      setter edge
//      .contains



func (g *Graph) setEdge(from, to int) {
    // get vertex
    fromVertex := g.getVertex(from)
    toVertex := g.getVertex(to)

    // check error
    if fromVertex == nil || toVertex == nil {
        err := fmt.Errorf("invalid setter: (%v-->%v)", from, to)
        fmt.Println(err.Error())
        return
    }

    // add/set edge
    fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}


func (e *Graph) printGraph() {
    for _, v := range e.vertices {
        fmt.Printf("\nvertex %v : ", v.key)
        for _, v := range v.adjacent {
            fmt.Printf("%v ", v.key)
        }
    }
    fmt.Println()
}


//  Drive

func main() {

    e := & Graph{}
    fmt.Println()


    for i := 0; i < 12; i++ {
        e.setVertex(i)
    }

    fmt.Println(e)

    /*
    // duplication-check
    e.setVertex(0)
    e.setVertex(0)
    */

    e.setEdge(1, 2)
    e.setEdge(12 + 1, 2) // 'from' is invalid
    e.printGraph()

}


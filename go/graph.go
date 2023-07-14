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

//  Method : Contain

func contains(vs []*Vertex, k int) bool {
    for _, v := range vs {
        if k == v.key {
            return true
        }
    }
    return false
}

//  Vertex

func (g * Graph) getVertex(k int) *Vertex {
    for i, v := range g.vertices {
        if v.key == k {
            return g.vertices[i]
        }
    }
    return nil
}

func (g *Graph) addVertex(k int) {
    // check if key already exists - .contains
    if contains(g.vertices, k) {
        err := fmt.Errorf("key existed:\t%v", k)
        fmt.Println(err.Error())
        return
    }
    g.vertices = append(g.vertices, & Vertex{ key: k })
}

//  Edges

func (g *Graph) addEdge(from, to int) {
    // get vertex
    fromVertex := g.getVertex(from)
    toVertex := g.getVertex(to)

    // check error
    if fromVertex == nil || toVertex == nil {
        // case : either `from` or `to` is wrong
        err := fmt.Errorf("invalid tunnel:\t(%v-->%v)", from, to)
        fmt.Println(err.Error())
        return
    } else if contains(fromVertex.adjacent, to) {
        err := fmt.Errorf("existing edge:\t(%v-->%v)", from, to)
        fmt.Println(err.Error())
        return
    }

    // add/add edge
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
        e.addVertex(i)
    }

    // print entire list of pointers
    // fmt.Println(e)

    /*
    // duplication-check
    e.addVertex(0)
    e.addVertex(0)
    */

    e.addEdge(4, 2)
    e.addEdge(4, 2)
    e.addEdge(12 + 1, 2) // 'from' is invalid
    e.printGraph()

}


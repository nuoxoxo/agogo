package main

import {
    "fmt"
}

type Graph struct {

    vertices []*Vertex
}

type Vertex struct {

    key int
    adjacent []*Vertex
}

//  Todo
//      getter/setter vertex
//      setter edge
//      .contains


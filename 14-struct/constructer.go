package main

import (
	"fmt"
	"time"
)

// like what we do in java , javascript classes object here in go is struct

// example of graph getting
type Edge struct{
	src int
	dest int
	cost int
	time time.Time // nanosecond precision
}
// constructor function
func NewEdge(src, dest, cost int, time time.Time)Edge {
	 return Edge{
		src : src,
		dest : dest,
		cost : cost,
		time : time,
	}
}

func constructor() {
	edge := NewEdge(1,2,3,time.Now())
	fmt.Println(edge)
	fmt.Println(edge.src)
	fmt.Println(edge.dest)
	fmt.Println(edge.cost)
	fmt.Println(edge.time)
}
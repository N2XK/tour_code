package main

import "fmt"

type Vertex4 struct {
	X int
	Y int
}

func main() {
	v := Vertex4{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

package main

import "fmt"

/** shape **/
type shape interface {
    getArea() float64
	//printArea()
}

func printArea(s shape) {
    fmt.Println(s.getArea())
}

/** End shape **/

/** triangle **/

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

/** End triangle **/

/** square **/

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

/** End square **/

func main() {
    triangle := triangle{
        base : 2,
        height : 4,
    }

    printArea(triangle)

    square := square{
        sideLength: 3,
    }

    printArea(square)
}

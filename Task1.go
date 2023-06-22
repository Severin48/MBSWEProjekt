package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type rectangle struct {
	length int
	width  int
}

type square struct {
	length int
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (s square) area() int {
	return s.length * s.length
}

type shape interface {
	area() int
}

func sumArea(x, y shape) int {
	return x.area() + y.area()
}

// Introducing unique function names for overloaded methods

func area_Rec(r rectangle) int {
	return r.length * r.width
}

func area_Sq(s square) int {
	return s.length * s.length
}

// "value" method implies "pointer" method
func area_RecPtr(r *rectangle) int {
	return area_Rec(*r)
}

func area_SqPtr(s *square) int {
	return area_Sq(*s)
}

// Run-time method lookup

func area_Lookup(x interface{}) int {
	var y int

	switch v := x.(type) {
	case square:
		y = area_Sq(v)
	case rectangle:
		y = area_Rec(v)
	}
	return y

}

func sumArea_Lookup(x, y interface{}) int {
	return area_Lookup(x) + area_Lookup(y)
}

// Dictionary translation

type shape_Value struct {
	val  interface{}
	area func(interface{}) int
}

func sumArea_Dict(x, y shape_Value) int {
	return x.area(x.val) + y.area(y.val)
}

func generateShapes(amount int) ([]shape, []shape_Value) {
	rand.Seed(time.Now().UnixNano())
	shapes := make([]shape, amount)
	shape_values := make([]shape_Value, amount)
	var rect_counter float64
	rect_counter = 0
	var square_counter float64
	square_counter = 0

	for i := 0; i < amount; i++ {
		// Determine whether to create a rectangle or a square
		shapeType := rand.Intn(2)

		// Create a rectangle or square with random dimensions
		if shapeType == 0 {
			r := rectangle{rand.Intn(100) + 1, rand.Intn(100) + 1}
			shapes[i] = r
			shape_values[i] = shape_Value{
				val: r,
				area: func(v interface{}) int {
					return v.(rectangle).area()
				},
			}
			rect_counter++
		} else {
			s := square{rand.Intn(100) + 1}
			shapes[i] = s
			shape_values[i] = shape_Value{
				val: s,
				area: func(v interface{}) int {
					return v.(square).area()
				},
			}
			square_counter++
		}
	}

	fmt.Printf("Rectangles: %d, Squares: %d, Ratio: %f : %f\n", int(rect_counter),
		int(square_counter), rect_counter/square_counter,
		square_counter/rect_counter)

	return shapes, shape_values
}

func Task1() {
	shape_amount := 10000000

	shapes, shape_values := generateShapes(shape_amount)

	start := time.Now()

	// Calculate areas with run-time method lookup
	for i := 0; i < shape_amount-1; i++ {
		sumArea_Lookup(shapes[i], shapes[i+1])
	}

	elapsed := time.Since(start)
	log.Printf("Run-time method lookup took %d ms", elapsed.Milliseconds())

	start = time.Now()

	// Calculate areas with run-time method lookup
	for i := 0; i < shape_amount-1; i++ {
		sumArea_Dict(shape_values[i], shape_values[i+1])
	}

	elapsed = time.Since(start)
	log.Printf("Dictionary translation took %d ms", elapsed.Milliseconds())
}

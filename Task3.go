package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Gegeben
var area_SqPtr_Wrapper = func(v interface{}) int {
	return area_SqPtr(v.(*square))
}

func sumAreaVariant(x, y shape) int {
	if z, ok := y.(square); ok {
		return x.area() + y.area() + z.length
	}
	return x.area() + y.area()
}

type shape_Value_Variant struct {
	val    interface{}
	area   func(interface{}) int
	length func(interface{}) int
}

func square_length(v interface{}) int {
	// Run-time type check
	if sq, ok := v.(square); ok {
		return sq.length
	}
	return 0
}

func sumAreaVariant_Dict(x, y shape_Value_Variant) int {
	if x.area == nil || y.area == nil {
		log.Fatal("Area function is nil")
	}

	length := 0
	if y.length != nil {
		length = y.length(y.val)
	}

	return x.area(x.val) + y.area(y.val) + length
}

func generateShapesT3(amount int) ([]shape, []shape_Value_Variant) {
	rand.Seed(time.Now().UnixNano())
	shapes := make([]shape, amount)
	shape_value_variants := make([]shape_Value_Variant, amount)
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
			shape_value_variants[i] = shape_Value_Variant{
				val: r,
				area: func(v interface{}) int {
					return v.(rectangle).area()
				},
			}
			rect_counter++
			// println("Rectangle with length", s.length, "and width", s.width)
		} else {
			s := square{rand.Intn(100) + 1}
			shapes[i] = s
			shape_value_variants[i] = shape_Value_Variant{
				val: s,
				area: func(v interface{}) int {
					return v.(square).area()
				},
				length: func(v interface{}) int {
					return v.(square).length
				},
			}
			square_counter++
			// println("Square with size", s.length)
		}
	}

	fmt.Printf("Rectangles: %d, Squares: %d, Ratio: %f : %f\n", int(rect_counter),
		int(square_counter), rect_counter/square_counter,
		square_counter/rect_counter)

	return shapes, shape_value_variants
}

func Task3() {
	shape_amount := 10000000
	shapes, shape_value_variants := generateShapesT3(shape_amount)

	start := time.Now()

	// Calculate areas with run-time method lookup
	for i := 0; i < shape_amount-1; i++ {
		sumAreaVariant(shapes[i], shapes[i+1])
	}

	elapsed := time.Since(start)
	log.Printf("Run-time method lookup took %d ms", elapsed.Milliseconds())

	start = time.Now()

	// Calculate areas with run-time method lookup
	for i := 0; i < shape_amount-1; i++ {
		sumAreaVariant_Dict(shape_value_variants[i], shape_value_variants[i+1])
	}

	elapsed = time.Since(start)
	log.Printf("Dictionary translation took %d ms", elapsed.Milliseconds())
}

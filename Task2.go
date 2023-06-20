package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

type body interface {
	volume() float32
}

type cube struct {
	length float32
	width  float32
	height float32
}

type sphere struct {
	radius float32
}

func (c cube) volume() float32 {
	return c.length * c.width * c.height
}

func (b sphere) volume() float32 {
	return 4 / 3 * math.Pi * b.radius
}

func sumVolume(x, y body) float32 {
	return x.volume() + y.volume()
}

// Introducing unique function names for overloaded methods

func vol_Cube(c cube) float32 {
	return c.length * c.width * c.height
}

func vol_Sphere(b sphere) float32 {
	return 4 / 3 * math.Pi * b.radius
}

// Run-time method lookup

func volume_Lookup(x interface{}) float32 {
	var y float32

	switch v := x.(type) {
	case cube:
		y = vol_Cube(v)
	case sphere:
		y = vol_Sphere(v)
	}
	return y

}

func sumVolume_Lookup(x, y interface{}) float32 {
	return volume_Lookup(x) + volume_Lookup(y)
}

// Dictionary translation

type body_Value struct {
	val    interface{}
	volume func(interface{}) float32
}

func sumVolume_Dict(x, y body_Value) float32 {
	return x.volume(x.val) + y.volume(y.val)
}

func generateBodies(amount int) ([]body, []body_Value) {
	rand.Seed(time.Now().UnixNano())
	bodies := make([]body, amount)
	body_values := make([]body_Value, amount)

	for i := 0; i < amount; i++ {
		// Determine whether to create a cube or sphere
		bodyType := rand.Intn(2)

		// Create a cube or sphere with random dimensions
		if bodyType == 0 {
			c := cube{rand.Float32() + 1, rand.Float32() + 1, rand.Float32() + 1}
			bodies[i] = c
			body_values[i] = body_Value{
				val: c,
				volume: func(v interface{}) float32 {
					return v.(cube).volume()
				},
			}
		} else {
			b := sphere{rand.Float32() + 1}
			bodies[i] = b
			body_values[i] = body_Value{
				val: b,
				volume: func(v interface{}) float32 {
					return v.(sphere).volume()
				},
			}
		}
	}

	return bodies, body_values
}

func Task2() {
	fmt.Printf("\nStarting task 2\n\n")
	body_amount := 10000000
	var cube_counter float64
	cube_counter = 0
	var sphere_counter float64
	sphere_counter = 0
	bodies, body_values := generateBodies(body_amount)
	for _, body := range bodies {
		// switch s := shape.(type) {
		switch body.(type) {
		case cube:
			cube_counter++
			// println("Cube with length=", s.length, ",width=", s.width, " and height=", s.height)
		case sphere:
			sphere_counter++
			// println("Sphere with size", s.length)
		}
	}

	fmt.Printf("Cubes: %d, Spheres: %d, Ratio: %f : %f\n", int(cube_counter),
		int(sphere_counter), cube_counter/sphere_counter,
		sphere_counter/cube_counter)

	start := time.Now()

	// Calculate volumes with run-time method lookup
	for i := 0; i < body_amount-1; i++ {
		sumVolume_Lookup(bodies[i], bodies[i+1])
	}

	elapsed := time.Since(start)
	log.Printf("Run-time method lookup took %d ms", elapsed.Milliseconds())

	start = time.Now()

	// Calculate volumes with run-time method lookup
	for i := 0; i < body_amount-1; i++ {
		sumVolume_Dict(body_values[i], body_values[i+1])
	}

	elapsed = time.Since(start)
	log.Printf("Dictionary translation took %d ms", elapsed.Milliseconds())

}

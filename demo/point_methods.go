package demo

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

type Point3 struct {
	X, Y, Z float64
}

type Polar struct {
	R, T float64
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

func (p *Polar) Abs() float64 {
	return math.Sqrt(float64(p.R*p.R + p.T*p.T))
}

func InitPointMethods() {
	p1 := &Point{10, 4}
	fmt.Printf("Point abs :%v\n", p1.Abs())
	p2 := &Point3{10, 4, 2}
	fmt.Printf("Point3 abs :%v\n", p2.Abs())
}

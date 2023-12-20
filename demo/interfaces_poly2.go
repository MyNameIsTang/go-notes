package demo

type Circle struct {
	radial float32
}

func (c *Circle) Area2() float32 {
	return 3.14 * c.radial * c.radial
}

type Shape3 struct{}

package sprint

type Circle struct {

	Radius float32
	Diameter float32
	Area float32
	Perimeter float32

}

func GetCircle(r float32) Circle {

	c := Circle{}

	c.Radius = r
	c.Diameter = r * 2
	c.Perimeter = 3.14 * c.Diameter
	c.Area = 3.14 * r * r

	return c

}
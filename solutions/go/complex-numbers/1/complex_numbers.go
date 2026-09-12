package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	real      float64
	imaginary float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		real:      n1.real + n2.real,
		imaginary: n1.imaginary + n2.imaginary,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		real:      n1.real - n2.real,
		imaginary: n1.imaginary - n2.imaginary,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	a, b := n1.real, n1.imaginary
	c, d := n2.real, n2.imaginary
	return Number{
		real:      a*c - b*d,
		imaginary: b*c + a*d,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		real:      n.real * factor,
		imaginary: n.imaginary * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	a, b := n1.real, n1.imaginary
	c, d := n2.real, n2.imaginary
	return Number{
		real:      (a*c + b*d) / (c*c + d*d),
		imaginary: (b*c - a*d) / (c*c + d*d),
	}
}

func (n Number) Conjugate() Number {
	a, b := n.real, n.imaginary
	return Number{
		real:      a,
		imaginary: -b,
	}
}

func (n Number) Abs() float64 {
	a, b := n.real, n.imaginary
	return math.Sqrt(a*a + b*b)
}

func (n Number) Exp() Number {
	a, b := n.real, n.imaginary
	return Number{
		real:      math.Exp(a) * (math.Cos(b)),
		imaginary: math.Exp(a) * (math.Sin(b)),
	}
}

package main

import "math"

type Vec3 [3]float64

type AffineTransform [2]Vec3

var affineTransformIdentity = AffineTransform{
	{1, 0, 0},
	{0, 1, 0},
}

func AffineTransformIdentity() *AffineTransform {
	return &affineTransformIdentity
}

func (p *Point) Transform(a *AffineTransform) Point {
	return Point{
		a[0][0]*p.X + a[0][1]*p.Y + a[0][1]*1,
		a[1][0]*p.X + a[1][1]*p.Y + a[1][1]*1,
	}
}

func (a0 *AffineTransform) Mul(a1 *AffineTransform) *AffineTransform {
	return &AffineTransform{
		{
			a0[0][0]*a1[0][0] + a0[0][1]*a1[1][0],
			a0[0][0]*a1[0][1] + a0[0][1]*a1[1][1],
			a0[0][0]*a1[0][2] + a0[0][1]*a1[1][2] + a0[0][2],
		},
		{
			a0[1][0]*a1[0][0] + a0[1][1]*a1[1][0],
			a0[1][0]*a1[0][1] + a0[1][1]*a1[1][1],
			a0[1][0]*a1[0][2] + a0[1][1]*a1[1][2] + a0[1][2],
		},
	}
}

func (a *AffineTransform) Inverse() *AffineTransform {

	d := a[0][0]*a[1][1] - a[0][1]*a[1][0]

	return &AffineTransform{
		{
			a[1][1] / d,
			-a[0][1] / d,
			(a[0][1]*a[1][2] - a[1][1]*a[0][2]) / d,
		},
		{
			-a[1][0] / d,
			a[0][0] / d,
			(a[1][0]*a[0][2] - a[0][0]*a[1][2]) / d,
		},
	}
}

func (a *AffineTransform) Scale(x, y float64) *AffineTransform {

	a1 := &AffineTransform{
		{x, 0, 0},
		{0, y, 0},
	}

	return a1.Mul(a)

	// return &AffineTransform{
	// 	{
	// 		a[0][0] * x,
	// 		a[0][1] * y,
	// 		a[0][2],
	// 	},
	// 	{
	// 		a[1][0] * x,
	// 		a[1][1] * y,
	// 		a[1][2],
	// 	},
	// }
}

func (a *AffineTransform) Translate(x, y float64) *AffineTransform {

	a1 := &AffineTransform{
		{1, 0, x},
		{0, 1, y},
	}

	return a1.Mul(a)

	// return &AffineTransform{
	// 	{
	// 		a[0][0],
	// 		a[0][1],
	// 		a[0][2] + a[0][0]*x + a[0][1]*y,
	// 	},
	// 	{
	// 		a[1][0],
	// 		a[1][1],
	// 		a[1][2] + a[1][0]*x + a[1][1]*y,
	// 	},
	// }
}

// RotateRadian 左回りの回転
func (a *AffineTransform) RotateRadian(r float64) *AffineTransform {

	s := math.Sin(r)
	c := math.Cos(r)

	a1 := &AffineTransform{
		{c, -s, 0},
		{s, c, 0},
	}

	return a1.Mul(a)
}

// RotateDegree 左回りの回転
func (a *AffineTransform) RotateDegree(r float64) *AffineTransform {
	rad := r * math.Pi / 180
	return a.RotateRadian(rad)
}

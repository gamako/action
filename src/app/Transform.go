package main

import (
	"fmt"
	"math"
)

// Point 2D座標を表す構造体
type Point struct {
	X float64
	Y float64
}

// Eq イコール
func (p0 *Point) Eq(p1 *Point) bool {
	return p0.X == p1.X && p0.Y == p1.Y
}

const eps = 1e-10

// Approximately おおよそのイコール
func (p0 *Point) Approximately(p1 *Point) bool {
	return math.Abs(p0.X-p1.X) < eps && math.Abs(p0.Y-p1.Y) < eps
}

// Size サイズを表す構造体
type Size struct {
	W float64
	H float64
}

// Transform 2Dtransformを表す構造体
type Transform struct {
	Point
	Scale Point
	Angle float64
}

// TransformIdentity 基本transform
var TransformIdentity = Transform{Point{0, 0}, Point{1, 1}, 0}

func Plus(p1, p2 *Point) Point {
	return Point{
		p1.X + p2.X,
		p1.Y + p2.Y,
	}
}

func (p1 *Point) Add(p2 *Point) Point {
	return Point{
		p1.X + p2.X,
		p1.Y + p2.Y,
	}
}

func Mul(p1 *Point, v float64) Point {
	return Point{
		p1.X * v,
		p1.Y * v,
	}
}

func (p1 *Point) Mul(v float64) Point {
	return Point{
		p1.X * v,
		p1.Y * v,
	}
}

func (p *Point) Normalized() Point {
	l := math.Sqrt(p.X*p.X + p.Y*p.Y)
	if l == 0 {
		return Point{0, 0}
	}

	return Point{
		p.X / l,
		p.Y / l,
	}
}

func (a *AffineTransform) Transform(p *Point) *Point {
	return &Point{
		a[0][0]*p.X + a[0][1]*p.Y + a[0][2]*1,
		a[1][0]*p.X + a[1][1]*p.Y + a[1][2]*1,
	}
}

func (t *Transform) GetAffineTransform() *AffineTransform {

	a := AffineTransformIdentity()

	return a.
		Scale(t.Scale.X, t.Scale.Y).
		RotateDegree(t.Angle).
		Translate(t.X, t.Y)
}

func CreateTransform(affine *AffineTransform) *Transform {

	a := affine[0][0]
	b := affine[0][1]
	c := affine[1][0]
	d := affine[1][1]

	var cosTh, r, w, h float64

	if math.Abs(a) > eps {

		if math.Abs(c) < eps {
			// sinTh = 0, cosTh = 1 or -1
			if a >= 0 {
				// cosTh = 1
				w = a
				h = d
				r = 0
			} else {
				// cosTh = -1
				w = -a
				h = -d
				r = 180
			}
		} else {

			if a >= 0 {
				cosTh = math.Sqrt((a * a) / (a*a + c*c))
			} else {
				cosTh = -math.Sqrt((a * a) / (a*a + c*c))
			}
			if cosTh >= 0 {
				if (a >= 0 && c >= 0) || (a < 0 && c < 0) {
					r = math.Acos(cosTh) / math.Pi * 180
				} else {
					r = 360 - math.Acos(cosTh)/math.Pi*180
				}
			} else {
				if (a >= 0 && c >= 0) || (a < 0 && c < 0) {
					r = 360 - math.Acos(cosTh)/math.Pi*180
				} else {
					r = math.Acos(cosTh) / math.Pi * 180
				}
			}
			w = a / cosTh
			h = d / cosTh
		}

	} else if math.Abs(c) > eps {
		// aが0だとすると、cos(theta) * Scale-w = 0
		// Scale-w != 0 であれば、cosTh = 0 である

		if c >= 0 {
			// sinTh = 1
			w = c
			h = -b
			r = float64(90)

		} else {
			// sinTh = -1
			w = -c
			h = b
			r = float64(270)
		}

	} else { //if math.Abs(e) > eps {
		// w = 0なので、ほぼ表示なしと思っても問題ない

		w = 0
		h = 0
		r = 0
	}

	t := &Transform{
		Point{affine[0][2], affine[1][2]},
		Point{w, h},
		r,
	}

	return t
}

func (t *Transform) String() string {
	return fmt.Sprintf("{{%f,%f},{%f,%f},%f}", t.X, t.Y, t.Scale.X, t.Scale.Y, t.Angle)
}

func (t0 *Transform) Eq(t1 *Transform) bool {
	return t0.X == t1.X &&
		t0.Y == t1.Y &&
		t0.Scale.X == t1.Scale.X &&
		t0.Scale.Y == t1.Scale.Y &&
		t0.Angle == t1.Angle
}

func Approximately(a, b float64) bool {
	return math.Abs(a-b) < eps
}

func (t0 *Transform) Approximately(t1 *Transform) bool {
	return t0.Point.Approximately(&t1.Point) &&
		t0.Scale.Approximately(&t1.Scale) &&
		Approximately(t0.Angle, t1.Angle)

}

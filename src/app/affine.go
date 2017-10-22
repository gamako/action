package main

import "math"

// Vec3 AffineTransformで使用する目的の行列の１行
type Vec3 [3]float64

// AffineTransform アフィン変換行列
type AffineTransform [2]Vec3

// AffineTransformIdentity 単位行列
var AffineTransformIdentity = &AffineTransform{
	{1, 0, 0},
	{0, 1, 0},
}

// MulPoint Point座業に変換行列を適用する
func (a0 *AffineTransform) MulPoint(p *Point) *Point {
	return &Point{
		a0[0][0]*p.X + a0[0][1]*p.Y + a0[0][1]*1,
		a0[1][0]*p.X + a0[1][1]*p.Y + a0[1][1]*1,
	}
}

// Mul 変換行列の合成
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

// Inverse 逆行列の作成
func (a0 *AffineTransform) Inverse() *AffineTransform {

	d := a0[0][0]*a0[1][1] - a0[0][1]*a0[1][0]

	return &AffineTransform{
		{
			a0[1][1] / d,
			-a0[0][1] / d,
			(a0[0][1]*a0[1][2] - a0[1][1]*a0[0][2]) / d,
		},
		{
			-a0[1][0] / d,
			a0[0][0] / d,
			(a0[1][0]*a0[0][2] - a0[0][0]*a0[1][2]) / d,
		},
	}
}

// Scale 拡大縮小
func (a0 *AffineTransform) Scale(x, y float64) *AffineTransform {

	return &AffineTransform{
		{
			a0[0][0] * x,
			a0[0][1] * y,
			a0[0][2],
		},
		{
			a0[1][0] * x,
			a0[1][1] * y,
			a0[1][2],
		},
	}
}

// Translate 移動
func (a0 *AffineTransform) Translate(x, y float64) *AffineTransform {

	a1 := &AffineTransform{
		{1, 0, x},
		{0, 1, y},
	}

	return a1.Mul(a0)
}

// RotateRadian 左回りの回転
func (a0 *AffineTransform) RotateRadian(r float64) *AffineTransform {

	s := math.Sin(r)
	c := math.Cos(r)

	a1 := &AffineTransform{
		{c, -s, 0},
		{s, c, 0},
	}

	return a1.Mul(a0)
}

// RotateDegree 左回りの回転
func (a0 *AffineTransform) RotateDegree(r float64) *AffineTransform {
	rad := r * math.Pi / 180
	return a0.RotateRadian(rad)
}

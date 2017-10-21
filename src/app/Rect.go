package main

// Rect 矩形
type Rect struct {
	X, Y, W, H float64
}

// MinX 左端
func (r *Rect) MinX() float64 { return r.X }

// MaxX 右端
func (r *Rect) MaxX() float64 { return r.X + r.W }

// MinY 上端
func (r *Rect) MinY() float64 { return r.Y }

// MaxY 下端
func (r *Rect) MaxY() float64 { return r.Y + r.H }

// Intersect 矩形同士の重なり反映
func Intersect(r0, r1 *Rect) bool {
	return !(r0.MaxX() < r1.MinX() ||
		r1.MaxX() < r0.MinX() ||
		r0.MaxY() < r1.MinY() ||
		r1.MaxY() < r0.MinY())
}

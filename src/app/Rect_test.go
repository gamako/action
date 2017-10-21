package main

import (
	"testing"
)

type InersectTestCase struct {
	r0, r1 Rect
	expect bool
}

var testCase = []InersectTestCase{
	// X軸方向の比較
	InersectTestCase{
		Rect{3, 0, 2, 2},
		Rect{0, 0, 2, 2},
		false,
	},
	InersectTestCase{
		Rect{3, 0, 2, 2},
		Rect{1, 0, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{3, 0, 2, 2},
		Rect{2, 0, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{3, 0, 2, 2},
		Rect{3, 0, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{3, 0, 2, 2},
		Rect{4, 0, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{3, 0, 2, 2},
		Rect{5, 0, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{3, 0, 2, 2},
		Rect{6, 0, 2, 2},
		false,
	},
	// Y軸方向の比較
	InersectTestCase{
		Rect{0, 3, 2, 2},
		Rect{0, 0, 2, 2},
		false,
	},
	InersectTestCase{
		Rect{0, 3, 2, 2},
		Rect{0, 1, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{0, 3, 2, 2},
		Rect{0, 2, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{0, 3, 2, 2},
		Rect{0, 3, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{0, 3, 2, 2},
		Rect{0, 4, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{0, 3, 2, 2},
		Rect{0, 5, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{0, 3, 2, 2},
		Rect{0, 6, 2, 2},
		false,
	},
	// 斜め
	InersectTestCase{
		Rect{3, 3, 2, 2},
		Rect{0, 0, 2, 2},
		false,
	},
	InersectTestCase{
		Rect{3, 3, 2, 2},
		Rect{6, 0, 2, 2},
		false,
	},
	InersectTestCase{
		Rect{3, 3, 2, 2},
		Rect{0, 6, 2, 2},
		false,
	},
	InersectTestCase{
		Rect{3, 3, 2, 2},
		Rect{6, 6, 2, 2},
		false,
	},
	InersectTestCase{
		Rect{3, 3, 2, 2},
		Rect{1, 1, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{3, 3, 2, 2},
		Rect{4, 1, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{3, 3, 2, 2},
		Rect{1, 4, 2, 2},
		true,
	},
	InersectTestCase{
		Rect{3, 3, 2, 2},
		Rect{4, 4, 2, 2},
		true,
	},
	// 完全に内方している
	InersectTestCase{
		Rect{3, 3, 3, 3},
		Rect{4, 4, 1, 1},
		true,
	},
	// 完全に内方している
	InersectTestCase{
		Rect{4, 4, 1, 1},
		Rect{3, 3, 3, 3},
		true,
	},
}

func TestIntersect(t *testing.T) {
	for _, v := range testCase {
		actual := Intersect(&v.r0, &v.r1)
		if actual != v.expect {
			t.Fatalf("%#v", v)
		}
	}
}

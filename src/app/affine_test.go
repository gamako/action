package main

import (
	"testing"
)

type convertSet struct {
	name   string
	expect *Transform
	f      func() *AffineTransform
}

var converttestset = []convertSet{
	convertSet{"Identity", &Transform{Point{0, 0}, Point{1, 1}, 0},
		func() *AffineTransform {
			return AffineTransformIdentity
		}},
	convertSet{"Inverse", &Transform{Point{0, 0}, Point{1, 1}, 0},
		func() *AffineTransform {
			return AffineTransformIdentity.Inverse()
		}},
	convertSet{"Translate", &Transform{Point{1, 0}, Point{1, 1}, 0},
		func() *AffineTransform {
			return AffineTransformIdentity.Translate(1, 0)
		}},
	convertSet{"Translate", &Transform{Point{1, 2}, Point{1, 1}, 0},
		func() *AffineTransform {
			return AffineTransformIdentity.Translate(1, 2)
		}},
	//
	convertSet{"Translate & Translate", &Transform{Point{11, 102}, Point{1, 1}, 0},
		func() *AffineTransform {
			return AffineTransformIdentity.Translate(1, 2).Translate(10, 100)
		}},
	convertSet{"Scale", &Transform{Point{0, 0}, Point{2, 1}, 0},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 1)
		}},
	convertSet{"Scale & Translate", &Transform{Point{1, 2}, Point{10, 100}, 0},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(10, 100).Translate(1, 2)
		}},
	convertSet{"Rotate & Translate", &Transform{Point{1, 2}, Point{1, 1}, 90},
		func() *AffineTransform {
			return AffineTransformIdentity.RotateDegree(90).Translate(1, 2)
		}},
	//
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 30},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(30).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 45},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(45).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 60},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(60).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 90},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(90).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 120},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(120).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 135},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(135).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 150},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(150).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 180},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(180).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 210},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(210).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 225},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(225).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 240},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(240).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 270},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(270).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 300},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(300).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 315},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(315).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 330},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(330).Translate(1, 2)
		}},
	convertSet{"Rotate & Scale & Translate", &Transform{Point{1, 2}, Point{2, 4}, 0},
		func() *AffineTransform {
			return AffineTransformIdentity.Scale(2, 4).RotateDegree(360).Translate(1, 2)
		}},
}

func TestAffine(t *testing.T) {
	for _, v := range converttestset {
		a := v.f()
		actual := CreateTransform(a)
		expext := v.expect

		if !actual.Approximately(expext) {
			t.Fatalf("failed test \"%s\" %#v %#v (%#v)", v.name, expext, actual, a)
		}
	}
}

type pointTransformationSet struct {
	expect *Point
	f      func() *Point
}

var zeroPoint = &Point{0, 0}

var transformationtset = []pointTransformationSet{
	// Identity
	pointTransformationSet{&Point{0, 0},
		func() *Point {
			return zeroPoint
		}},
	// Translate
	pointTransformationSet{&Point{10, 20},
		func() *Point {
			return AffineTransformIdentity.Translate(10, 20).Transform(&Point{0, 0})
		}},
	pointTransformationSet{&Point{11, 22},
		func() *Point {
			return AffineTransformIdentity.Translate(10, 20).Transform(&Point{1, 2})
		}},
	// Scale
	pointTransformationSet{&Point{0, 0},
		func() *Point {
			return AffineTransformIdentity.Scale(2, 3).Transform(&Point{0, 0})
		}},
	pointTransformationSet{&Point{2, 6},
		func() *Point {
			return AffineTransformIdentity.Scale(2, 3).Transform(&Point{1, 2})
		}},
	// Rotate
	pointTransformationSet{&Point{0, 0},
		func() *Point {
			return AffineTransformIdentity.RotateDegree(90).Transform(&Point{0, 0})
		}},
	// Rotate
	pointTransformationSet{&Point{0, 0},
		func() *Point {
			return AffineTransformIdentity.RotateDegree(45).Transform(&Point{0, 0})
		}},
	// Rotate
	pointTransformationSet{&Point{0, 1},
		func() *Point {
			return AffineTransformIdentity.RotateDegree(90).Transform(&Point{1, 0})
		}},
}

func TestTransformation(t *testing.T) {
	for _, v := range transformationtset {
		actual := v.f()
		expext := v.expect

		if !actual.Approximately(expext) {
			t.Fatalf("failed test %#v %#v", actual, expext)
		}
	}
}

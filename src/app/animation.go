package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Animation アニメーションの基底インターフェース
type Animation interface {
	Start()
	Draw(r *sdl.Renderer, t float64, transform *Transform)
}

// FiniteAnimation 時間権限のあるアニメーション
type FiniteAnimation interface {
	Draw(r *sdl.Renderer, t float64, transform *Transform)
	Duration() float64
}

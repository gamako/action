package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Animation アニメーションの基底インターフェース
type Animation interface {
	Start()
	Draw(r *sdl.Renderer, t float64, x float64, y float64, scale float64, angle float64)
}

// FiniteAnimation 時間権限のあるアニメーション
type FiniteAnimation interface {
	DrawRatio(r *sdl.Renderer, t float64, x float64, y float64, scale float64, angle float64)
	Duration() float64
}

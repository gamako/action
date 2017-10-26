package main

import "github.com/veandco/go-sdl2/sdl"
import "math"

// LoopAnimation FiniteAnimationの繰り返し
type LoopAnimation struct {
	animation FiniteAnimation
	startTime float64
	isStarted bool
	isEnabled bool
}

// CreateLoopAnimation 生成
func CreateLoopAnimation(animation FiniteAnimation) *LoopAnimation {
	return &LoopAnimation{
		animation,
		0,
		false,
		false}
}

// Start 開始
func (a *LoopAnimation) Start() {
	a.isEnabled = true
}

// Draw LoopAnimationの描画
func (a *LoopAnimation) Draw(r *sdl.Renderer, now float64, t *Transform) {

	if !a.isEnabled {
		return
	}

	if !a.isStarted {
		a.isStarted = true
		a.startTime = now
	}

	elapsed := now - a.startTime

	ratio := math.Mod(elapsed, a.animation.Duration())

	a.animation.Draw(r, ratio, t)
}

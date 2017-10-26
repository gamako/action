package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// FrameAnimation パラパラアニメの定義
type FrameAnimation struct {
	textures []*Texture
	duration float64
	unit     float64
}

// CreateFrameAnimation 生成
func CreateFrameAnimation(textures []*Texture, duration float64) *FrameAnimation {
	unit := duration / float64(len(textures))

	return &FrameAnimation{
		textures,
		duration,
		unit}
}

// FiniteAnimationのインターフェース

// Duration 長さ
func (a *FrameAnimation) Duration() float64 {
	return a.duration
}

// Draw 描画
// t は 0~Durationの間
func (a *FrameAnimation) Draw(r *sdl.Renderer, time float64, transform *Transform) {

	index := int(time / a.unit)

	if index >= len(a.textures) {
		index = len(a.textures) - 1
	}

	tex := a.textures[index]

	w := tex.Surface.W
	h := tex.Surface.H

	srcRect := sdl.Rect{W: w, H: h}
	dstW := float64(w) * transform.Scale.X
	dstH := float64(h) * transform.Scale.Y

	dstRect := sdl.Rect{X: int32(transform.X - dstW/2), Y: int32(transform.Y - dstH/2), W: int32(dstW), H: int32(dstH)}

	// dstRectで拡大率
	// angleで回転
	r.CopyEx(tex.Texture, &srcRect, &dstRect, transform.Angle, nil, 0)
}

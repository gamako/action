package main

import "github.com/veandco/go-sdl2/sdl"

// SpriteNode
type SpriteNode struct {
	Transform
	texture  *Texture
	children []Node
}

// CreateSpriteNode 生成
func CreateSpriteNode(texture *Texture, children []Node) *SpriteNode {

	return &SpriteNode{TransformIdentity, texture, children}
}

// GetTransform Transform情報
func (n *SpriteNode) GetTransform() *Transform {
	return &n.Transform
}

// Update Update
func (n *SpriteNode) Update(now float64) {
}

// Draw Draw
func (n *SpriteNode) Draw(r *sdl.Renderer, parentTransform *AffineTransform, _ float64) {

	tex := n.texture

	w := tex.Surface.W
	h := tex.Surface.H

	srcRect := sdl.Rect{W: w, H: h}
	dstW := float64(w) * n.Scale.X
	dstH := float64(h) * n.Scale.Y

	dstRect := sdl.Rect{X: int32(n.X - dstW/2), Y: int32(n.Y - dstH/2), W: int32(dstW), H: int32(dstH)}

	// dstRectで拡大率
	// angleで回転
	r.CopyEx(tex.Texture, &srcRect, &dstRect, n.Angle, nil, 0)
}

// Chilidren 子
func (n *SpriteNode) Chilidren() []Node {
	return n.children
}

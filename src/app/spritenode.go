package main

import "github.com/veandco/go-sdl2/sdl"

// SpriteNode
type SpriteNode struct {
	NodeBase
	texture *Texture
}

// CreateSpriteNode 生成
func CreateSpriteNode(name string, texture *Texture, children []Node) *SpriteNode {

	return &SpriteNode{
		*CreateNodeBase(name),
		texture,
	}
}

// GetTransform Transform情報
func (n *SpriteNode) GetTransform() *Transform {
	return &n.Transform
}

// Update Update
func (n *SpriteNode) Update(now float64) {
	n.NodeBase.Update(now)
}

// Draw Draw
func (n *SpriteNode) Draw(r *sdl.Renderer, parentTransform *AffineTransform, now float64) {

	tex := n.texture

	a := parentTransform.Mul(n.GetAffineTransform())
	transform := CreateTransform(a)

	w := tex.Surface.W
	h := tex.Surface.H

	srcRect := sdl.Rect{W: w, H: h}
	dstX := transform.X
	dstY := transform.Y
	dstW := float64(w) * transform.Scale.X
	dstH := float64(h) * transform.Scale.Y

	dstRect := sdl.Rect{X: int32(dstX - dstW/2), Y: int32(dstY - dstH/2), W: int32(dstW), H: int32(dstH)}

	// dstRectで拡大率
	// angleで回転
	r.CopyEx(tex.Texture, &srcRect, &dstRect, n.Angle, nil, 0)

	DrawChildren(r, n, a, now)
}

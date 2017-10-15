package main

import "github.com/veandco/go-sdl2/sdl"

// SpriteNode
type EmptyNode struct {
	Transform
	Children []Node
}

func (n *EmptyNode) GetTransform() *Transform {
	return &n.Transform
}
func (n *EmptyNode) Update(now float64) {
}

func (n *EmptyNode) Draw(r *sdl.Renderer, parentTransform *AffineTransform, now float64) {

	a := parentTransform.Mul(n.GetAffineTransform())

	DrawChildren(r, n, a, now)
}

func (n *EmptyNode) GetChildren() []Node {
	return n.Children
}

func (n *EmptyNode) AddChild(child Node) {
	n.Children = append(n.Children, child)
}

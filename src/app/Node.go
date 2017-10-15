package main

import "github.com/veandco/go-sdl2/sdl"

// Node ノードインターフェース
// 表示キャラクタの管理と描画に関するインターフェース
type Node interface {
	GetTransform() *Transform
	Update(float64)
	Draw(*sdl.Renderer, *AffineTransform, float64)
	GetChildren() []Node
}

func Update(node Node, now float64) {

	node.Update(now)

	UpdateChildren(node, now)
}

func UpdateChildren(node Node, now float64) {
	for _, child := range node.GetChildren() {
		Update(child, now)
	}
}

func DrawChildren(r *sdl.Renderer, node Node, parentTransform *AffineTransform, now float64) {

	for _, child := range node.GetChildren() {
		child.Draw(r, parentTransform, now)
	}
}

type NodeBase struct {
	Transform
	Children []Node
}

func (n *NodeBase) GetTransform() *Transform {
	return &n.Transform
}

func (n *NodeBase) Update(float64) {

}

func (n *NodeBase) Draw(r *sdl.Renderer, parentTransform *AffineTransform, now float64) {
	a := parentTransform.Mul(n.GetAffineTransform())

	DrawChildren(r, n, a, now)
}

func (n *NodeBase) GetChildren() []Node {
	return n.Children
}

func (n *NodeBase) AddChild(child Node) {
	n.Children = append(n.Children, child)
}

func (n *NodeBase) DeleteChild(child Node) {

	result := []Node{}
	for _, v := range n.Children {
		if v != child {
			result = append(result, v)
		}
	}
	n.Children = result
}

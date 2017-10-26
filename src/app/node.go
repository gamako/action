package main

import "github.com/veandco/go-sdl2/sdl"

// Node ノードインターフェース
// 表示キャラクタの管理と描画に関するインターフェース
type Node interface {
	GetName() string
	GetTransform() *Transform
	Update(float64)
	Draw(*sdl.Renderer, *AffineTransform, float64)
	GetChildren() []Node
}

func Update(node Node, now float64) {

	// fmt.Printf("Update: %#v\n", node)
	node.Update(now)
	UpdateChildren(node, now)
}

func UpdateChildren(node Node, now float64) {

	// fmt.Printf("UpdateChildren: %#v (%d)\n", node, len(node.GetChildren()))
	for _, child := range node.GetChildren() {
		Update(child, now)
	}
}

func DrawChildren(r *sdl.Renderer, node Node, parentTransform *AffineTransform, now float64) {

	for _, child := range node.GetChildren() {
		child.Draw(r, parentTransform, now)
	}
}

// NodeBase ノードの基本的な実装がされたstruct
type NodeBase struct {
	Name string
	Transform
	Children []Node
}

// CreateNodeBase 生成
func CreateNodeBase(name string) *NodeBase {
	return &NodeBase{
		name,
		TransformIdentity,
		[]Node{},
	}
}

// GetName 名前の取得
func (n *NodeBase) GetName() string {
	return n.Name
}

// GetTransform 位置情報
func (n *NodeBase) GetTransform() *Transform {
	return &n.Transform
}

// Update フレーム毎処理
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

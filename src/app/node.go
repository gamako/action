package main

import "github.com/veandco/go-sdl2/sdl"

// Node ノードインターフェース
// 表示キャラクタの管理と描画に関するインターフェース
type Node interface {
	GetName() string
	GetTransform() *Transform
	Update(float64)
	Draw(*sdl.Renderer, *AffineTransform, float64)
	EachChild(func(Node))
}

// Update nodeの毎フレーム更新
func Update(node Node, now float64) {

	// fmt.Printf("Update: %#v\n", node)
	node.Update(now)
	UpdateChildren(node, now)
}

// UpdateChildren 子のnodeの毎フレーム更新
func UpdateChildren(node Node, now float64) {

	node.EachChild(func(child Node) {
		Update(child, now)
	})

}

// DrawChildren 子のDraw
func DrawChildren(r *sdl.Renderer, node Node, parentTransform *AffineTransform, now float64) {

	node.EachChild(func(child Node) {
		child.Draw(r, parentTransform, now)
	})
}

// NodeBase ノードの基本的な実装がされたstruct
type NodeBase struct {
	Name string
	Transform
	Children      []child
	ChildrenDirty bool
}

type child struct {
	Node
	valid bool
}

// CreateNodeBase 生成
func CreateNodeBase(name string) *NodeBase {
	return &NodeBase{
		name,
		TransformIdentity,
		[]child{},
		false,
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
	n.FlushChild()
}

// Draw 描画
func (n *NodeBase) Draw(r *sdl.Renderer, parentTransform *AffineTransform, now float64) {
	a := parentTransform.Mul(n.GetAffineTransform())
	DrawChildren(r, n, a, now)
}

// EachChild 子要素のiterate
func (n *NodeBase) EachChild(f func(Node)) {
	for _, v := range n.Children {
		if v.valid {
			f(v)
		}
	}
}

// AddChild 要素の削除
func (n *NodeBase) AddChild(c Node) {
	n.Children = append(n.Children, child{c, true})
}

// DeleteChild 要素の削除
// この中ではマークするだけ
// 実際のエントリの削除はFlushChildのときに行われる
func (n *NodeBase) DeleteChild(child Node) {
	for i := range n.Children {
		if n.Children[i].Node == child {
			n.Children[i].valid = false
			n.ChildrenDirty = true
		}
	}
}

// FlushChild DeleteChildを呼ばれたノードを実際に削除する
func (n *NodeBase) FlushChild() {
	if n.ChildrenDirty {
		return
	}

	result := []child{}
	for _, v := range n.Children {
		if v.valid {
			result = append(result, v)
		}
	}
	n.Children = result

}

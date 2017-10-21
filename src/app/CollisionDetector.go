package main

// CollisionDetector 衝突判定
// Nodeオブジェクトとして動く
type CollisionDetector struct {
	NodeBase
	collisionNodes        []collisionNodeInfo
	tagFilter             CollisionTagFilter
	collisionNodesIsDirty bool
}

type collisionNodeInfo struct {
	node      CollisonNode
	onCollide OnDetectCollision
	valid     bool
}

// CollisonNode 衝突判定の対象ノードの満たすべきインターフェース
type CollisonNode interface {
	CollisonTag() int
	ColliderRect() *Rect
	GetNode() Node
}

// OnDetectCollision 衝突が発生したときのコールバックインターフェース
type OnDetectCollision interface {
	OnCollision(other CollisonNode)
}

// CollisionTagFilter 衝突判定を行うかどうかを制御するためのフィルタ
type CollisionTagFilter interface {
	CollisionFilter(tag1, tag2 int) bool
}

// AllCollideTags すべてヒット対象とするフィルタ
type AllCollideTags struct {
}

// CreateCollisionDetector 生成
func CreateCollisionDetector(filter CollisionTagFilter) *CollisionDetector {
	return &CollisionDetector{
		*CreateNodeBase("CollisionDetector"),
		[]collisionNodeInfo{},
		filter,
		false,
	}
}

// CollisionFilter AllCollideTagsのフィルタ条件
func (o AllCollideTags) CollisionFilter(tag1, tag2 int) bool {
	return true
}

// AddCollisionNode 衝突判定対象の追加
func (d *CollisionDetector) AddCollisionNode(node CollisonNode, onCollide OnDetectCollision) {

	newInfo := collisionNodeInfo{node, onCollide, true}
	d.collisionNodes = append(d.collisionNodes, newInfo)
}

// DeleteCollisionNode 衝突判定対象の削除
// 削除マークをつけておく
func (d *CollisionDetector) DeleteCollisionNode(node CollisonNode) {

	for i := range d.collisionNodes {
		if d.collisionNodes[i].node == node {
			d.collisionNodes[i].valid = false
			d.collisionNodesIsDirty = true
		}
	}
}

func (d *CollisionDetector) detectCollisions() {
	l := len(d.collisionNodes)
	for i, n0 := range d.collisionNodes {
		if !n0.valid {
			continue
		}
		for j := i + 1; j < l; j++ {
			n1 := d.collisionNodes[j]
			if !n1.valid {
				continue
			}
			if Intersect(n0.node.ColliderRect(), n1.node.ColliderRect()) {
				n0.onCollide.OnCollision(n1.node)
				n1.onCollide.OnCollision(n0.node)
			}
		}
	}
}

func (d *CollisionDetector) flush() {

	nodes := []collisionNodeInfo{}
	for _, n := range d.collisionNodes {
		if n.valid {
			nodes = append(nodes, n)
		}
	}
	d.collisionNodes = nodes
}

// Update 毎フレームの処理
func (d *CollisionDetector) Update(now float64) {

	d.detectCollisions()

	if d.collisionNodesIsDirty {
		d.flush()
		d.collisionNodesIsDirty = false
	}

}

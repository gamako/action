package main

// Bullet 弾オブジェクト
type Bullet struct {
	SpriteNode
	Delta        Point
	ColliderSize Size
}

// CreateBullet 生成
func CreateBullet(tex *Texture, delta Point, colliderSize Size) *Bullet {

	b := &Bullet{
		*CreateSpriteNode("Bullet", tex, []Node{}),
		delta,
		colliderSize,
	}
	b.Transform.Scale = Point{4, 4}

	return b
}

// Update 毎フレームの更新
func (b *Bullet) Update(now float64) {

	b.X += b.Delta.X
	b.Y += b.Delta.Y

	if (b.X < 0 || b.X > 800) || (b.Y < 0 || b.Y > 600) {
		// 画面外に出たので消える
		GlobalNode.DeleteChild(b)
		GlobalCollisionDetecter.DeleteCollisionNode(b)
	}
}

func (e *Bullet) CollisonTag() int {
	return TAG_BULLET
}

func (e *Bullet) ColliderRect() *Rect {
	r := Rect{
		e.Transform.Point.X - e.ColliderSize.W/2*e.Transform.Scale.X,
		e.Transform.Point.Y - e.ColliderSize.H/2*e.Transform.Scale.Y,
		e.ColliderSize.W * e.Transform.Scale.Y,
		e.ColliderSize.H * e.Transform.Scale.Y,
	}
	return &r
}

func (e *Bullet) GetNode() Node {
	return e
}

func (e *Bullet) OnCollision(other CollisonNode) {
	GlobalCollisionDetecter.DeleteCollisionNode(e)
	GlobalNode.DeleteChild(e)
}

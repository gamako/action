package main

import (
	"fmt"
)

// Bullet 弾オブジェクト
type Bullet struct {
	SpriteNode
	Delta Point
}

// CreateBullet 生成
func CreateBullet(tex *Texture, delta Point) *Bullet {

	return &Bullet{
		*CreateSpriteNode(tex, []Node{}),
		delta,
	}
}

// Update 毎フレームの更新
func (b *Bullet) Update(now float64) {

	b.X += b.Delta.X
	b.Y += b.Delta.Y

	if (b.X < 0 || b.X > 800) || (b.Y < 0 || b.Y > 600) {
		// 画面外に出たので消える
		GlobalNode.DeleteChild(b)

		fmt.Println("Global count", len(GlobalNode.Children))
	}
}

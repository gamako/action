package main

import (
	"math"
)

// Enemy 敵
type Enemy struct {
	SpriteNode
	targetPlayer *Player
	speed        float64
	ColliderSize Size
}

// CreateEnemy 生成
func CreateEnemy(name string, tex *Texture, colliderSize Size) *Enemy {
	e := &Enemy{
		*CreateSpriteNode(name, tex, []Node{}),
		nil,
		1,
		colliderSize,
	}

	e.Transform.Scale = Point{2, 2}
	return e
}

// Update 更新
func (e *Enemy) Update(now float64) {
	if e.targetPlayer == nil {
		// 一番近いPlayerを探す
		ps := GetActivePlayers()
		neaestDistance := math.MaxFloat64
		for _, p := range ps {
			if !p.enabled {
				break
			}
			dx := e.X - p.X
			dy := e.Y - p.Y
			d := math.Sqrt(dx*dx + dy*dy)

			if neaestDistance > d {
				neaestDistance = d
				e.targetPlayer = p
			}
		}
	}

	if e.targetPlayer == nil {
		// Playerがいない場合は何もしない
		return
	}

	v := (&Point{
		e.targetPlayer.Point.X - e.Point.X,
		e.targetPlayer.Point.Y - e.Point.Y,
	}).Normalized()

	dv := v.Mul(e.speed)

	e.Point = (&e.Point).Add(&dv)
}

func (e *Enemy) CollisonTag() int {
	return TAG_ENEMY
}

func (e *Enemy) ColliderRect() *Rect {
	r := Rect{
		e.Transform.Point.X - e.ColliderSize.W/2*e.Transform.Scale.X,
		e.Transform.Point.Y - e.ColliderSize.H/2*e.Transform.Scale.Y,
		e.ColliderSize.W,
		e.ColliderSize.H,
	}
	return &r
}

func (e *Enemy) GetNode() Node {
	return e
}

func (e *Enemy) OnCollision(other CollisonNode) {
	if other.CollisonTag() == TAG_BULLET {
		GlobalCollisionDetecter.DeleteCollisionNode(e)
	}
	GlobalNode.DeleteChild(e)
}

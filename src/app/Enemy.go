package main

import (
	"math"
)

type EnemyManager struct {
	NodeBase
}

func (n *EnemyManager) Update(now float64) {

}

type Enemy struct {
	SpriteNode
	targetPlayer *Player
	speed        float64
}

func CreateEnemy(name string, tex *Texture) *Enemy {
	e := &Enemy{
		*CreateSpriteNode(name, tex, []Node{}),
		nil,
		1,
	}

	e.Transform.Scale = Point{2, 2}
	return e
}

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
			d := math.Sqrt(dx*dx - dy*dy)
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

package main

import (
	"math/rand"
)

// GameLevelManager レベルデザインを担当するクラス
type GameLevelManager struct {
	NodeBase
	nextEnemyTime float64
	screenSize    Point
	textures      []*Texture
}

// CreateGameLevelManager 生成
func CreateGameLevelManager(screenSize Point, textures []*Texture) *GameLevelManager {
	return &GameLevelManager{
		nextEnemyTime: 10,
		textures:      textures,
	}
}

func randInRange(min, max float64) float64 {
	v := rand.Float64()

	return v*(max-min) + min
}

// Update 毎フレーム処理
func (m *GameLevelManager) Update(now float64) {

	if now > m.nextEnemyTime {
		m.nextEnemyTime += 5
		// 敵の発生
		var startPoint Point
		if rand.Float32() > 0.5 {
			var x float64
			if rand.Float32() > 0.5 {
				x = 0
			} else {
				x = m.screenSize.X
			}
			startPoint = Point{x, randInRange(0, m.screenSize.Y)}
		} else {
			var y float64
			if rand.Float32() > 0.5 {
				y = 0
			} else {
				y = m.screenSize.Y
			}
			startPoint = Point{randInRange(0, m.screenSize.X), y}
		}

		newEnemy := CreateEnemy("enemy", m.textures[0])
		newEnemy.Transform.Point = startPoint

		GlobalNode.AddChild(newEnemy)

	}
}

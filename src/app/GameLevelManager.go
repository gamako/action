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
	spawnPoints   []Point
}

// CreateGameLevelManager 生成
func CreateGameLevelManager(screenSize Point, textures []*Texture, spawnPoints []Point) *GameLevelManager {

	rand.Seed(0)

	return &GameLevelManager{
		nextEnemyTime: 10,
		textures:      textures,
		spawnPoints:   spawnPoints,
	}
}

func randInRange(min, max float64) float64 {
	v := rand.Float64()

	return v*(max-min) + min
}

func randInRangeInt(min, max int) int {
	v := rand.Float64()

	return int(v*(float64(max-min))) + min
}

// Update 毎フレーム処理
func (m *GameLevelManager) Update(now float64) {

	if now > m.nextEnemyTime {
		m.nextEnemyTime++

		// 敵のスポーンポイントをランダムに選ぶ
		index := randInRangeInt(0, len(m.spawnPoints))
		spawnPoint := m.spawnPoints[index]

		// 敵の発生
		newEnemy := CreateEnemy("enemy", m.textures[0])
		newEnemy.Transform.Point = spawnPoint

		GlobalNode.AddChild(newEnemy)

	}
}

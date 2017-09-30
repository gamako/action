package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// CotrolerManager コントローラー管理
type CotrolerManager struct {
	// NumOfControllers  int
	OpenedControllers []*GameController
}

// GameController CotrolerManagerが管理しているControllerの情報
type GameController struct {
	*sdl.GameController
}

// SDLController sdlのコントローラー情報の取得
func (c *GameController) SDLController() *sdl.GameController {
	return c.GameController
}

// GetNewGameController 新しく接続されたコントローラーオブジェクトがあれば取得
func (m *CotrolerManager) GetNewGameController() *GameController {

	num := sdl.NumJoysticks()

	if len(m.OpenedControllers) == num {
		// コントローラーが増えてない場合
		return nil
	}

	if len(m.OpenedControllers) > num {
		// Open中に無効になる場合があるのか？
		fmt.Println("len(m.OpenedControllers) < num ..", len(m.OpenedControllers), "/", num)
		return nil
	}

	// 新しいコントローラーオブジェクトをOpen
	newIndex := len(m.OpenedControllers)
	sdlController := sdl.GameControllerOpen(newIndex)

	fmt.Printf("New Controller (%d) : %s\n", newIndex, sdlController.Name())

	// 内部型に変換
	controller := &GameController{sdlController}

	// 管理用リストに追加
	m.OpenedControllers = append(m.OpenedControllers, controller)

	fmt.Println("len(m.OpenedControllers) = ", len(m.OpenedControllers))

	return controller
}

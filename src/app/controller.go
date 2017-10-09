package main

import "github.com/veandco/go-sdl2/sdl"
import "math"

// Controller キーボードを含んだコントローラーの抽象化インターフェース
type Controller interface {
	Name() string
	GetAxis(axis sdl.GameControllerAxis) int16
	GetButton(btn sdl.GameControllerButton) byte
	Close()
}

// KeybordController キーボードをgameControllerと同じインターフェースで扱うためのstruct
type KeybordController struct {
}

// Name 名前の取得
func (k *KeybordController) Name() string {
	return "Keybord"
}

// GetAxis スティック情報
func (k *KeybordController) GetAxis(axis sdl.GameControllerAxis) int16 {
	keyboardState := sdl.GetKeyboardState()

	if axis == sdl.CONTROLLER_AXIS_LEFTX || axis == sdl.CONTROLLER_AXIS_RIGHTX {
		if keyboardState[sdl.SCANCODE_RIGHT] != 0 {
			return math.MaxInt16
		} else if keyboardState[sdl.SCANCODE_LEFT] != 0 {
			return math.MinInt16
		} else {
			return 0
		}
	}
	if axis == sdl.CONTROLLER_AXIS_LEFTY || axis == sdl.CONTROLLER_AXIS_RIGHTY {
		if keyboardState[sdl.SCANCODE_DOWN] != 0 {
			return math.MaxInt16
		} else if keyboardState[sdl.SCANCODE_UP] != 0 {
			return math.MinInt16
		} else {
			return 0
		}
	}

	return 0
}

// GetButton ボタン情報
func (k *KeybordController) GetButton(btn sdl.GameControllerButton) byte {

	keyboardState := sdl.GetKeyboardState()

	if btn == sdl.CONTROLLER_BUTTON_A {
		if keyboardState[sdl.SCANCODE_SPACE] != 0 {
			return 1
		}
		return 0
	}
	return 0
}

// Close クローズ
func (k *KeybordController) Close() {
	// 特にする事がない
}

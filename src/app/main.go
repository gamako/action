package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Point struct {
	X int32
	Y int32
}

type Player struct {
	Point
}

func main() {
	// sdlの初期化
	sdl.Init(sdl.INIT_EVERYTHING)
	// 最後にsdlの終了
	defer sdl.Quit()

	// sdlで扱うWindowの生成
	window, err := sdl.CreateWindow("action", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		// あまり失敗しないと思われるので手抜きのエラー処理
		panic(err)
	}
	// 最後にwindowの後始末
	defer window.Destroy()

	// windowに描画するためのrendererオブジェクトを取得
	render, renderErr := sdl.CreateRenderer(window, -1, 0)
	if renderErr != nil {
		// あまり失敗しないと思われるので手抜きのエラー処理
		panic(err)
	}
	// 最後にrendererの後始末
	defer render.Destroy()

	t1, s1, err := loadTexture(render, "images/a.png")
	if err != nil {
		panic(err)
	}
	defer t1.Destroy()
	defer s1.Free()

	// 指定の色で全体をクリア
	render.SetDrawColor(0, 0, 0, 255)
	render.Clear()
	// renderをwindowに反映
	render.Present()

	// コンソールへの出力を確認するためにPrintしてみる
	fmt.Println("start!!")

	var p1 = Player{Point{X: 0, Y: 0}}

	for i := 0; i < 10000; i++ {

		// var dx int = 0

		sdl.PumpEvents()

		// Player情報を更新
		keyboardState := sdl.GetKeyboardState()
		if keyboardState[sdl.SCANCODE_RIGHT] != 0 {
			p1.X += 1
		} else if keyboardState[sdl.SCANCODE_LEFT] != 0 {
			p1.X += -1
		}

		if keyboardState[sdl.SCANCODE_UP] != 0 {
			p1.Y += -1
		} else if keyboardState[sdl.SCANCODE_DOWN] != 0 {
			p1.Y += 1
		}

		// 毎回の画面の更新

		// 指定の色で全体をクリア
		render.SetDrawColor(0, 0, 0, 255)
		render.Clear()

		srcRect := sdl.Rect{W: s1.W, H: s1.H}
		dstRect := sdl.Rect{X: p1.X, Y: p1.Y, W: s1.W, H: s1.H}

		angle := float64(i % 360)
		// dstRectで拡大率
		// 次の引数で
		render.CopyEx(t1, &srcRect, &dstRect, angle, nil, 0)

		// renderをwindowに反映
		render.Present()

		// ちょっとだけ待つ
		sdl.Delay(10 * 1)

	}

	// しばし表示している
	sdl.Delay(1000 * 2)
}

func loadTexture(r *sdl.Renderer, name string) (*sdl.Texture, *sdl.Surface, error) {
	s, err := img.Load(name)
	if err != nil {
		return nil, nil, err
	}
	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		s.Free()
		return nil, nil, err
	}
	return t, s, nil
}

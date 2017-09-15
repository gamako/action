package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

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

	// 指定の色で全体をクリア
	render.SetDrawColor(0, 0, 0, 255)
	render.Clear()
	// renderをwindowに反映
	render.Present()

	for i := 0; i < 10000; i++ {

		// 毎回の画面の更新

		// 指定の色で全体をクリア
		render.SetDrawColor(0, 0, 0, 255)
		render.Clear()

		// 指定の色、場所に四角を描く
		// だんだんと動くように、カウンタを元に座標計算
		render.SetDrawColor(255, 0, 0, 255)
		rect := sdl.Rect{X: 0, Y: 0, W: int32(i % 800), H: 200}
		render.FillRect(&rect)

		// renderをwindowに反映
		render.Present()

		// ちょっとだけ待つ
		sdl.Delay(10 * 1)

	}

	// しばし表示している
	sdl.Delay(1000 * 2)
}

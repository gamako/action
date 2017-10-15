package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var imageFiles = []string{
	"images/anim/character0000.png",
	"images/anim/character0001.png",
	"images/anim/character0002.png",
	"images/bullet1.png",
}

const (
	IndexChar0 int = iota
	IndexChar1
	IndexChar2
	IndexBullet
)

var GlobalNode = &EmptyNode{
	NodeBase{
		Transform{Point{0, 0}, Point{1, 1}, 0},
		[]Node{},
	},
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
	renderer, renderErr := sdl.CreateRenderer(window, -1, 0)
	if renderErr != nil {
		// あまり失敗しないと思われるので手抜きのエラー処理
		panic(err)
	}
	// 最後にrendererの後始末
	defer renderer.Destroy()

	ts, err := LoadTextures(renderer, imageFiles)
	if err != nil {
		panic(err)
	}
	defer func() {
		for _, v := range ts {
			v.Free()
		}
	}()

	// 指定の色で全体をクリア
	renderer.SetDrawColor(128, 128, 128, 255)
	renderer.Clear()
	// renderをwindowに反映
	renderer.Present()

	// コンソールへの出力を確認するためにPrintしてみる
	fmt.Println("start!!")

	man := CotrolerManager{}

	startTime := time.Now()

	for {
		now := time.Since(startTime).Seconds()

		{
			cs := man.GetNewGameController()
			for _, c := range cs {
				p := CreatePlayer(ts, c)
				p.animation.Start()

				GlobalNode.AddChild(p)
			}
		}

		// var dx int = 0

		sdl.PumpEvents()
		// sdl.GameControllerUpdate()

		// Nodeをそれぞれ更新
		Update(GlobalNode, now)

		// 毎回の画面の更新

		// 指定の色で全体をクリア
		renderer.SetDrawColor(128, 128, 128, 255)
		renderer.Clear()

		GlobalNode.Draw(renderer, AffineTransformIdentity(), now)

		// renderをwindowに反映
		renderer.Present()

		// ちょっとだけ待つ
		sdl.Delay(10 * 1)

		if sdl.GetKeyboardState()[sdl.SCANCODE_Q] != 0 {
			// Qキーで終了
			break
		}
	}

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

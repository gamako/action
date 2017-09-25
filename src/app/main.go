package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Point 2D座標を表す構造体
type Point struct {
	X float64
	Y float64
}

// Node ノードインターフェース
// 表示キャラクタの管理と描画に関するインターフェース
type Node interface {
	Update()
	Draw(*sdl.Renderer)
}

// Player プレイヤー情報
type Player struct {
	Point
	angle float64

	t *sdl.Texture
	s *sdl.Surface
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

	t1, s1, err := loadTexture(renderer, "images/a.png")
	if err != nil {
		panic(err)
	}
	defer t1.Destroy()
	defer s1.Free()

	// 指定の色で全体をクリア
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	// renderをwindowに反映
	renderer.Present()

	// コンソールへの出力を確認するためにPrintしてみる
	fmt.Println("start!!")

	nodes := []Node{}

	p1 := Player{Point{0, 0}, 0, t1, s1}

	nodes = append(nodes, &p1)

	for i := 0; i < 10000; i++ {

		// var dx int = 0

		sdl.PumpEvents()

		// Player情報を更新
		p1.Update()

		// 毎回の画面の更新

		// 指定の色で全体をクリア
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		for _, n := range nodes {
			n.Draw(renderer)
		}

		// renderをwindowに反映
		renderer.Present()

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

// Update 毎フレームの更新関数
func (p *Player) Update() {
	keyboardState := sdl.GetKeyboardState()

	space := keyboardState[sdl.SCANCODE_SPACE]

	rad := p.angle * float64(math.Pi) / 180

	if space != 0 {
		if keyboardState[sdl.SCANCODE_RIGHT] != 0 {
			p.Y += math.Sin(rad) * 1
			p.X += math.Cos(rad) * 1
		} else if keyboardState[sdl.SCANCODE_LEFT] != 0 {
			p.Y += math.Sin(rad) * -1
			p.X += math.Cos(rad) * -1
		}
	} else {
		if keyboardState[sdl.SCANCODE_RIGHT] != 0 {
			p.angle++
		} else if keyboardState[sdl.SCANCODE_LEFT] != 0 {
			p.angle--
		}
	}

	if keyboardState[sdl.SCANCODE_UP] != 0 {
		p.Y += math.Cos(rad) * -1
		p.X += math.Sin(rad) * 1
	} else if keyboardState[sdl.SCANCODE_DOWN] != 0 {
		p.Y += math.Cos(rad) * 1
		p.X += math.Sin(rad) * -1
	}
}

// Draw 描画
func (p *Player) Draw(r *sdl.Renderer) {

	srcRect := sdl.Rect{W: p.s.W, H: p.s.H}
	dstRect := sdl.Rect{X: int32(p.X), Y: int32(p.Y), W: p.s.W, H: p.s.H}

	// dstRectで拡大率
	// angleで回転
	r.CopyEx(p.t, &srcRect, &dstRect, p.angle, nil, 0)
}

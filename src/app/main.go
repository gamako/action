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

	controller Controller

	t []*Texture

	enabled bool
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

	// t1, s1, err := loadTexture(renderer, "images/a.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer t1.Destroy()
	// defer s1.Free()

	ts, err := LoadTextures(renderer, []string{
		"images/anim/character0000.png",
		"images/anim/character0001.png",
		"images/anim/character0002.png"})
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

	nodes := []Node{}

	for {

		{
			cs := man.GetNewGameController()
			for _, c := range cs {
				p1 := Player{Point{0, 0}, 0, c, ts, false}
				nodes = append(nodes, &p1)
			}
		}

		// var dx int = 0

		sdl.PumpEvents()
		// sdl.GameControllerUpdate()

		// Nodeをそれぞれ更新
		for _, n := range nodes {
			n.Update()
		}

		// 毎回の画面の更新

		// 指定の色で全体をクリア
		renderer.SetDrawColor(128, 128, 128, 255)
		renderer.Clear()

		for _, n := range nodes {
			n.Draw(renderer)
		}

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

// Update 毎フレームの更新関数
func (p *Player) Update() {

	leftX := p.controller.GetAxis(sdl.CONTROLLER_AXIS_LEFTX)
	leftY := p.controller.GetAxis(sdl.CONTROLLER_AXIS_LEFTY)
	rightX := p.controller.GetAxis(sdl.CONTROLLER_AXIS_RIGHTX)
	rightY := p.controller.GetAxis(sdl.CONTROLLER_AXIS_RIGHTY)
	//leftStick := p.controller.GetButton(sdl.CONTROLLER_BUTTON_LEFTSTICK)
	//leftShoulder := p.controller.GetButton(sdl.CONTROLLER_BUTTON_LEFTSHOULDER)
	abutton := p.controller.GetButton(sdl.CONTROLLER_BUTTON_A)
	//buttonDUp := p.controller.GetButton(sdl.CONTROLLER_BUTTON_DPAD_UP)

	// ボタンを押したコントローラーが有効になる
	if !p.enabled {
		if abutton != 0 {
			p.enabled = true
		} else {
			return
		}
	}

	// fmt.Printf("left: %d, %d right: %d, %d lstick:%d lshoulder:%d a:%d dup:%d\n",
	// 	leftX, leftY, rightX, rightY, leftStick, leftShoulder, abutton, buttonDUp)

	// 回転角度の決定 左スティックによる
	{
		x := float64(rightX) / -math.MinInt16
		y := float64(rightY) / -math.MinInt16

		if math.Abs(x) >= 0.2 || math.Abs(y) >= 0.2 {
			if y == 0 {
				if x > 0 {
					p.angle = 90
				} else if x < 0 {
					p.angle = 270
				} else {
					// ボタンを押していない間はangleをキープ
				}
			} else {
				angle := math.Atan(-x/y) / math.Pi * 180

				if y > 0 {
					angle += 180
				}

				p.angle = angle
			}

		}
	}

	// 自身の移動
	{
		x := float64(leftX) / -math.MinInt16
		y := float64(leftY) / -math.MinInt16

		if math.Abs(x) >= 0.2 || math.Abs(y) >= 0.2 {
			// たぶんここに速度係数を掛ける
			p.X += x
			p.Y += y
		}
	}

}

// Draw 描画
func (p *Player) Draw(r *sdl.Renderer) {

	if !p.enabled {
		return
	}

	w := p.t[0].Surface.W
	h := p.t[0].Surface.H

	srcRect := sdl.Rect{W: w, H: h}
	dstRect := sdl.Rect{X: int32(p.X), Y: int32(p.Y), W: w, H: h}

	// dstRectで拡大率
	// angleで回転
	r.CopyEx(p.t[0].Texture, &srcRect, &dstRect, p.angle, nil, 0)
}

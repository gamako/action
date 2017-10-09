package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Point 2D座標を表す構造体
type Point struct {
	X float64
	Y float64
}

// Size サイズを表す構造体
type Size struct {
	W float64
	H float64
}

// Transform 2Dtransformを表す構造体
type Transform struct {
	Point
	Scale float64
	Angle float64
}

// TransformIdentity 基本transform
var TransformIdentity = Transform{Point{0, 0}, 1, 0}

// Node ノードインターフェース
// 表示キャラクタの管理と描画に関するインターフェース
type Node interface {
	GetTransform() *Transform
	Update(float64)
	Draw(*sdl.Renderer, *Transform, float64)
	Chilidren() []Node
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

	startTime := time.Now()

	for {
		now := time.Since(startTime).Seconds()

		{
			cs := man.GetNewGameController()
			for _, c := range cs {
				p := CreatePlayer(ts, c)
				p.animation.Start()

				nodes = append(nodes, &p)
			}
		}

		// var dx int = 0

		sdl.PumpEvents()
		// sdl.GameControllerUpdate()

		// Nodeをそれぞれ更新
		for _, n := range nodes {
			n.Update(now)
		}

		// 毎回の画面の更新

		// 指定の色で全体をクリア
		renderer.SetDrawColor(128, 128, 128, 255)
		renderer.Clear()

		for _, n := range nodes {
			n.Draw(renderer, &TransformIdentity, now)
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

func Update(node Node, now float64) {

	node.Update(now)

	UpdateChildren(node, now)
}

func UpdateChildren(node Node, now float64) {
	for _, child := range node.Chilidren() {
		Update(child, now)
	}
}

func Draw(r *sdl.Renderer, node Node, parentTransform *Transform, now float64) {

	node.Draw(r, parentTransform, now)

	DrawChildren(r, node, parentTransform, now)
}

func DrawChildren(r *sdl.Renderer, node Node, parentTransform *Transform, now float64) {

	t := mul(node.GetTransform(), parentTransform)
	for _, child := range node.Chilidren() {
		Draw(r, child, t, now)
	}
}

func mul(t1, t2 *Transform) *Transform {
	// 簡易的に位置と大きさだけを扱う
	return &Transform{
		plus(&t1.Point, &t2.Point),
		t1.Scale + t2.Scale,
		t2.Angle}
}

func plus(p1, p2 *Point) Point {
	return Point{
		p1.X + p2.X,
		p1.Y + p2.Y}
}

package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// Player プレイヤー情報
type Player struct {
	Transform
	controller Controller
	animation  Animation
	enabled    bool
}

// CreatePlayer 生成
func CreatePlayer(ts []*Texture, c Controller) Player {
	f := CreateFrameAnimation([]*Texture{ts[0], ts[1]}, 1)
	animation := CreateLoopAnimation(f)

	return Player{TransformIdentity, c, animation, false}
}

// GetTransform Transform情報
func (p *Player) GetTransform() *Transform {
	return &p.Transform
}

// Update 毎フレームの更新関数
func (p *Player) Update(now float64) {

	leftX := p.controller.GetAxis(sdl.CONTROLLER_AXIS_LEFTX)
	leftY := p.controller.GetAxis(sdl.CONTROLLER_AXIS_LEFTY)
	rightX := p.controller.GetAxis(sdl.CONTROLLER_AXIS_RIGHTX)
	rightY := p.controller.GetAxis(sdl.CONTROLLER_AXIS_RIGHTY)
	// leftStick := p.controller.GetButton(sdl.CONTROLLER_BUTTON_LEFTSTICK)
	// leftShoulder := p.controller.GetButton(sdl.CONTROLLER_BUTTON_LEFTSHOULDER)
	abutton := p.controller.GetButton(sdl.CONTROLLER_BUTTON_A)
	// buttonDUp := p.controller.GetButton(sdl.CONTROLLER_BUTTON_DPAD_UP)

	// ボタンを押したコントローラーが有効になる
	if !p.enabled {
		if abutton != 0 {
			println("player enabled :", p.controller.Name())
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
					p.Transform.Angle = 90
				} else if x < 0 {
					p.Transform.Angle = 270
				} else {
					// ボタンを押していない間はangleをキープ
				}
			} else {
				angle := math.Atan(-x/y) / math.Pi * 180

				if y > 0 {
					angle += 180
				}

				p.Transform.Angle = angle
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
func (p *Player) Draw(r *sdl.Renderer, parentTransform *Transform, now float64) {

	if !p.enabled {
		return
	}

	p.animation.Draw(r, now, &p.Transform)

	DrawChildren(r, p, parentTransform, now)
}

// Chilidren 子
func (p *Player) Chilidren() []Node {
	return []Node{}
}

package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// Player プレイヤー情報
type Player struct {
	NodeBase
	controller Controller
	animation  Animation

	eyeSprite     *SpriteNode
	bulletTexture *Texture

	enabled   bool
	buttonOff bool
	shotAngle Point
}

// CreatePlayer 生成
func CreatePlayer(ts []*Texture, c Controller) *Player {
	f := CreateFrameAnimation([]*Texture{ts[0], ts[1]}, 1)
	animation := CreateLoopAnimation(f)

	eyeSprite := CreateSpriteNode("eye", ts[2], []Node{})

	p := Player{
		*CreateNodeBase("Player"),
		c,
		animation,
		eyeSprite,
		ts[IndexBullet],
		false,
		false,
		Point{},
	}

	p.AddChild(eyeSprite)

	p.Scale.X = 2
	p.Scale.Y = 2

	return &p
}

var ActivePlayers []*Player

func GetActivePlayers() []*Player {
	return ActivePlayers
}

// GetTransform Transform情報
func (p *Player) GetTransform() *Transform {
	return &p.Transform
}

// Update 毎フレームの更新関数
func (p *Player) Update(now float64) {

	leftX := p.controller.GetAxis(sdl.CONTROLLER_AXIS_LEFTX)
	leftY := p.controller.GetAxis(sdl.CONTROLLER_AXIS_LEFTY)
	// rightX := p.controller.GetAxis(sdl.CONTROLLER_AXIS_RIGHTX)
	// rightY := p.controller.GetAxis(sdl.CONTROLLER_AXIS_RIGHTY)
	// leftStick := p.controller.GetButton(sdl.CONTROLLER_BUTTON_LEFTSTICK)
	// leftShoulder := p.controller.GetButton(sdl.CONTROLLER_BUTTON_LEFTSHOULDER)
	abutton := p.controller.GetButton(sdl.CONTROLLER_BUTTON_A)
	// buttonDUp := p.controller.GetButton(sdl.CONTROLLER_BUTTON_DPAD_UP)

	// ボタンを押したコントローラーが有効になる
	if !p.enabled {
		if abutton != 0 {
			println("player enabled :", p.controller.Name())
			p.enabled = true
			ActivePlayers = append(ActivePlayers, p)
		} else {
			return
		}
	}

	// fmt.Printf("left: %d, %d right: %d, %d lstick:%d lshoulder:%d a:%d dup:%d\n",
	// 	leftX, leftY, rightX, rightY, leftStick, leftShoulder, abutton, buttonDUp)

	// 回転角度の決定 左スティックによる
	// {
	// 	x := float64(rightX) / -math.MinInt16
	// 	y := float64(rightY) / -math.MinInt16

	// 	if math.Abs(x) >= 0.2 || math.Abs(y) >= 0.2 {
	// 		if y == 0 {
	// 			if x > 0 {
	// 				p.Transform.Angle = 90
	// 			} else if x < 0 {
	// 				p.Transform.Angle = 270
	// 			} else {
	// 				// ボタンを押していない間はangleをキープ
	// 			}
	// 		} else {
	// 			angle := math.Atan(-x/y) / math.Pi * 180

	// 			if y > 0 {
	// 				angle += 180
	// 			}

	// 			p.Transform.Angle = angle
	// 		}

	// 	}
	// }

	// 自身の移動
	{
		x := float64(leftX) / -math.MinInt16
		y := float64(leftY) / -math.MinInt16

		if math.Abs(x) >= 0.2 || math.Abs(y) >= 0.2 {
			// たぶんここに速度係数を掛ける
			p.X += x
			p.Y += y

			p.eyeSprite.X = x
			p.eyeSprite.Y = y

			p1 := Point{x, y}
			p.shotAngle = p1.Normalized()
		} else {
			p.eyeSprite.X = 0
			p.eyeSprite.Y = 0
		}
	}

	if p.buttonOff {
		if abutton != 0 && !(p.shotAngle.X == 0 && p.shotAngle.Y == 0) {
			p.buttonOff = false
			// 弾を打つ
			b := CreateBullet(p.bulletTexture, p.shotAngle.Mul(4), p.bulletTexture.Size())
			b.Transform.Point = p.Transform.Point

			GlobalNode.AddChild(b)
			GlobalCollisionDetecter.AddCollisionNode(b, b.OnCollision)
		}
	} else {
		if abutton == 0 {
			p.buttonOff = true
		}
	}

}

// Draw 描画
func (p *Player) Draw(r *sdl.Renderer, parentTransform *AffineTransform, now float64) {

	if !p.enabled {
		return
	}

	a := parentTransform.Mul(p.GetAffineTransform())
	t := CreateTransform(a)

	p.animation.Draw(r, now, t)

	DrawChildren(r, p, a, now)
}

package main

import "github.com/veandco/go-sdl2/sdl"

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

    // windowに描画するためのsurfaceオブジェクトを取得
    surface, err := window.GetSurface()
    if err != nil {
        // あまり失敗しないと思われるので手抜きのエラー処理
        panic(err)
    }

    // 試しに四角を描いてみる
    rect := sdl.Rect{X: 0, Y: 0, W:200, H:200}
    surface.FillRect(&rect, 0xffff0000)
    // windowのsurfinceを更新することで画面を更新する
    window.UpdateSurface()

    // しばし表示している
    sdl.Delay(1000 * 2)
}

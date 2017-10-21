package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Texture テクスチャをロード、アンロードを簡単にするためのstruct
type Texture struct {
	*sdl.Texture
	*sdl.Surface
}

// LoadTexture ファイルからロード
func LoadTexture(r *sdl.Renderer, name string) (*Texture, error) {
	s, err := img.Load(name)
	if err != nil {
		return nil, err
	}
	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		s.Free()
		return nil, err
	}
	return &Texture{t, s}, nil
}

// LoadTextures LoadTextureの配列版
func LoadTextures(r *sdl.Renderer, names []string) ([]*Texture, error) {

	var texs = make([]*Texture, len(names))

	for i, name := range names {
		tex, err := LoadTexture(r, name)
		if err != nil {
			for _, v := range texs {
				if v != nil {
					v.Free()
				}
			}
			return nil, err
		}
		texs[i] = tex
	}
	return texs, nil
}

// Free Textuerの中身のClose
func (t *Texture) Free() {
	t.Texture.Destroy()
	t.Surface.Free()
}

// Size サイズの取得
func (t *Texture) Size() Size {
	return Size{
		float64(t.W),
		float64(t.H),
	}
}

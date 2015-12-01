package drip

import "github.com/veandco/go-sdl2/sdl"

type Texture struct {
	W int32 // Width; may not be modified
	H int32 // Height; may not be modified
	t *sdl.Texture
}

func newTexture(t *sdl.Texture, w, h int32) *Texture {
	return &Texture{
		t: t,
		W: w,
		H: h,
	}
}

func (t *Texture) Close() {
	t.t.Destroy()
}

type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

func (r *Rect) toSDL() *sdl.Rect {
	return &sdl.Rect{
		X: r.X,
		Y: r.Y,
		W: r.W,
		H: r.H,
	}
}

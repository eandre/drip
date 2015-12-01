package drip

import "github.com/veandco/go-sdl2/sdl"

type Window struct {
	*Renderer
	w *sdl.Window
}

func NewWindow(title string, width, height int) (*Window, error) {
	pos := sdl.WINDOWPOS_UNDEFINED
	w, err := sdl.CreateWindow(title, pos, pos, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	r, err := sdl.CreateRenderer(w, -1 /* first driver */, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}
	renderer := newRenderer(r)

	return &Window{
		Renderer: renderer,
		w:        w,
	}, nil
}

func (w *Window) Close() {
	// Close the renderer attached to this window before destroying the window
	w.Renderer.Close()
	w.w.Destroy()
}

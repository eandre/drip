package drip

import "github.com/veandco/go-sdl2/sdl"

type Renderer struct {
	r        *sdl.Renderer
	firstErr error
}

func (r *Renderer) Close() {
	r.r.Destroy()
}

func newRenderer(r *sdl.Renderer) *Renderer {
	return &Renderer{r: r}
}

func (r *Renderer) ResetError() {
	r.firstErr = nil
}

func (r *Renderer) Error() error {
	return r.firstErr
}

func (r *Renderer) Clear() error {
	err := r.r.Clear()
	return r.setError(err)
}

func (r *Renderer) CreateTextureFromSurface(s *Surface) (*Texture, error) {
	tex, err := r.r.CreateTextureFromSurface(s.s)
	if err != nil {
		r.setError(err)
		return nil, err
	}
	return newTexture(tex, s.W, s.H), nil
}

func (r *Renderer) DrawLine(x1, y1, x2, y2 int) error {
	err := r.r.DrawLine(x1, y1, x2, y2)
	return r.setError(err)
}

func (rr *Renderer) SetDrawColor(r, g, b, a uint8) error {
	err := rr.r.SetDrawColor(r, g, b, a)
	return rr.setError(err)
}

func (r *Renderer) Present() {
	r.r.Present()
}

func (r *Renderer) DrawTexture(t *Texture, src, dst Rect) error {
	err := r.r.Copy(t.t, src.toSDL(), dst.toSDL())
	return r.setError(err)
}

// setError sets the firstErr field on the renderer if
// the passed in err is not nil. It always returns the same
// error as was given, so that one can simply do this:
//
//		return r.setErr(err)
//
func (r *Renderer) setError(err error) error {
	if r.firstErr == nil && err != nil {
		r.firstErr = err
	}
	return err
}

package drip

import (
	"errors"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

type Surface struct {
	W int32 // Width; may not be modified
	H int32 // Height; may not be modified
	s *sdl.Surface
}

func ReadSurfaceFromFile(path string) (*Surface, error) {
	surf, err := img.Load(path)
	if err != nil {
		return nil, err
	}
	return newSurface(surf), nil
}

func LoadSurface(p []byte) (*Surface, error) {
	if len(p) == 0 {
		return nil, errors.New("drip: LoadSurface passed an empty slice")
	}

	ptr := unsafe.Pointer(&p[0])
	rw := sdl.RWFromMem(ptr, len(p))
	surf, err := img.Load_RW(rw, true)
	if err != nil {
		return nil, err
	}
	return newSurface(surf), nil
}

func newSurface(s *sdl.Surface) *Surface {
	return &Surface{
		W: s.W,
		H: s.H,
		s: s,
	}
}

func (s *Surface) Close() {
	s.s.Free()
}

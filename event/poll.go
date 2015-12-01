package event

import "github.com/veandco/go-sdl2/sdl"

func Poll() Event {
	e := sdl.PollEvent()
	if e == nil {
		return nil
	}
	return wrapEvent(e)
}

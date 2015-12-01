package event

import (
	"fmt"
	"time"

	"github.com/eandre/drip/internal/sdlinit"
	"github.com/eandre/drip/key"
	"github.com/veandco/go-sdl2/sdl"
)

type Event interface {
	fmt.Stringer
	Timestamp() time.Time
}

type Quit struct {
	ts time.Time
}

func (e *Quit) String() string       { return "Quit event" }
func (e *Quit) Timestamp() time.Time { return e.ts }

func fromQuitEvent(e *sdl.QuitEvent) *Quit {
	return &Quit{
		ts: sdlinit.TimestampToTime(e.Timestamp),
	}
}

type KeyDown struct {
	ts     time.Time
	Key    key.Key
	Repeat bool
}

type KeyUp struct {
	ts     time.Time
	Key    key.Key
	Repeat bool
}

func (e *KeyDown) String() string       { return e.Key.Name() + " key down" }
func (e *KeyDown) Timestamp() time.Time { return e.ts }
func (e *KeyUp) String() string         { return e.Key.Name() + " key up" }
func (e *KeyUp) Timestamp() time.Time   { return e.ts }

func fromKeyDownEvent(e *sdl.KeyDownEvent) *KeyDown {
	return &KeyDown{
		ts:     sdlinit.TimestampToTime(e.Timestamp),
		Key:    key.Key(e.Keysym.Sym),
		Repeat: e.Repeat != 0,
	}
}

func fromKeyUpEvent(e *sdl.KeyUpEvent) *KeyUp {
	return &KeyUp{
		ts:     sdlinit.TimestampToTime(e.Timestamp),
		Key:    key.Key(e.Keysym.Sym),
		Repeat: e.Repeat != 0,
	}
}

func wrapEvent(e sdl.Event) Event {
	switch e := e.(type) {
	default:
		// Unknown event; ignore it
		return nil
	case *sdl.QuitEvent:
		return fromQuitEvent(e)
	case *sdl.KeyDownEvent:
		return fromKeyDownEvent(e)
	case *sdl.KeyUpEvent:
		return fromKeyUpEvent(e)
	}
}

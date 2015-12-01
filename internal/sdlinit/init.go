package sdlinit

import (
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func InitAll() {
	runtime.LockOSThread()
	sdl.Init(sdl.INIT_EVERYTHING)
	setInitTime(time.Now())
}

func Quit() {
	sdl.Quit()
}

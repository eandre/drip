package main

import (
	"log"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/eandre/drip"
)

func main() {
	runtime.LockOSThread()

	w, err := drip.NewWindow("MOM, GET THE CAMERA!", 500, 500)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	surf, err := drip.ReadSurfaceFromFile("./bush.png")
	if err != nil {
		log.Fatal(err)
	}
	defer surf.Close()

	tex, err := w.CreateTextureFromSurface(surf)
	if err != nil {
		log.Fatal(err)
	}
	defer tex.Close()

	var frames uint64 = 0
	go frameCounter(&frames)

	for {
		w.SetDrawColor(64, 64, 64, 255)
		w.Clear()
		w.SetDrawColor(255, 255, 255, 255)
		w.DrawLine(0, 0, 100, 100)

		w.DrawTexture(tex,
			drip.Rect{0, 0, tex.W, tex.H},     // src
			drip.Rect{100, 100, tex.W, tex.H}, // dst
		)
		if err := w.Error(); err != nil {
			log.Println("Could not draw:", err)
		}
		w.Present()

		atomic.AddUint64(&frames, 1)
	}
}

func frameCounter(counter *uint64) {
	for {
		before := atomic.LoadUint64(counter)
		time.Sleep(1 * time.Second)
		after := atomic.LoadUint64(counter)
		log.Println("Rendered", after-before, "fps")
	}
}

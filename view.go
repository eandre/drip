package drip

// View provides a view into the world, decoupling world space from view
// space.  This allows for the view (or camera) to move around
// without having to move the whole world.
type View struct {
	Width  float32
	Height float32
	CX     float32 // center x
	CY     float32 // center y
}

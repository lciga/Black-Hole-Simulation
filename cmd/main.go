package main

import (
	"Black-Hole-Simulation/internal/bodies"
	"Black-Hole-Simulation/internal/engine"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	eng := engine.New2DEngine(800, 600, "Black Hole Simulator", rl.Black)
	bh := bodies.NewBlackHole(5.972e24, [2]int32{600, 300})
	r := bodies.NewRay([2]int32{100, 300}, [2]int32{1, 0})
	eng.Init(func() {
		bh.DrawBlackHole(color.RGBA{255, 0, 0, 255}, 5000)
		r.Update(100000000)
		r.DrawRay(color.RGBA{255, 255, 255, 255}, 10)
	})
}

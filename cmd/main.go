package main

import (
	"Black-Hole-Simulation/internal/bodies"
	"Black-Hole-Simulation/internal/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	eng := engine.New2DEngine(800, 600, "Black Hole Simulator", rl.Black)

	bh := bodies.NewBlackHole(5.972e24, [2]float64{400, 300})

	r := bodies.NewRay([2]float64{100, 300}, [2]float64{1, 0})

	// Единый масштаб пикселей на метр для рисования и проверки горизонта
	const pxPerMeter = 5000.0
	// Масштаб скорости, чтобы луч не «улетал» слишком быстро
	const speedScale = 1e8

	eng.Init(func() {
		bh.DrawBlackHole(rl.Red, pxPerMeter)
		r.Update(speedScale, pxPerMeter, bh)
		r.DrawRay(rl.White, 10)
	})
}

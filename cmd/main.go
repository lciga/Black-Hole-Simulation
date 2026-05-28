package main

import (
	"Black-Hole-Simulation/internal/bodies"
	"Black-Hole-Simulation/internal/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	eng := engine.New2DEngine(800, 600, "Black Hole Simulator", rl.Black)

	bh := bodies.NewBlackHole(5.972e24, [2]float64{400, 300})

	const pxPerMeter = 5000.0
	const speedScale = 1e8

	rays := make([]*bodies.Ray, 0)

	count := 10
	offset := 50.0
	total := offset * float64(count-1)

	for i := 0; i < count; i++ {
		y := 300.0 + float64(i)*offset - total/2

		ray := bodies.NewRay(
			[2]float64{100, y},
			[2]float64{1, 0},
		)

		rays = append(rays, ray)
	}

	eng.Init(func() {
		bh.DrawBlackHole(rl.Red, pxPerMeter)

		for _, ray := range rays {
			ray.Update(speedScale, pxPerMeter, bh)
			ray.DrawRay(rl.White)
		}
	})
}

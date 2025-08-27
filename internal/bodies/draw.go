// Пакет отвечает за моделирование и отрисовку тел
package bodies

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Функция отрисовки черной дыры
func (bh *BlackHole) DrawBlackHole(col rl.Color, scale float64) {
	radius := bh.rs * scale
	rl.DrawCircle(int32(bh.position[0]), int32(bh.position[1]), float32(radius), col)
}

// Функця отрисовки луча с перемещением по оси X
func (r *Ray) DrawRay(col rl.Color, count int) {
	if r.absorbed {
		return
	}

	offset := int32(50)
	total := offset * int32(count-1)
	startX := int32(r.position[0])

	for i := 0; i < count; i++ {
		dy := int32(i)*offset - total/2

		// Рисуем трейл
		for j := 1; j < len(r.trail); j++ {
			p0 := r.trail[j-1]
			p1 := r.trail[j]
			rl.DrawLine(int32(p0.X), int32(p0.Y)+dy, int32(p1.X), int32(p1.Y)+dy, col)
		}
		// Рисуем точку
		y := int32(r.position[1]) + dy
		rl.DrawCircle(startX, y, 0.5, col)
	}
}

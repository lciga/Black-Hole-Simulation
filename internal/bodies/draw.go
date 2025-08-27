// Пакет отвечает за моделирование и отрисовку тел
package bodies

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Функция отрисовки черной дыры
func (bh *BlackHole) DrawBlackHole(col rl.Color, scale float32) {
	radius := bh.Rs * scale
	rl.DrawCircle(bh.Position[0], bh.Position[1], radius, col)
}

// Функця отрисовки луча с перемещением по оси X
func (r *Ray) DrawRay(col rl.Color, count int) {
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
		y := int32(r.position[1]) + dy
		rl.DrawCircle(startX, y, 3, col)
	}
}

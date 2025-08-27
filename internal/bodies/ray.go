package bodies

import rl "github.com/gen2brain/raylib-go/raylib"

type Ray struct {
	position  [2]int32     // Двухмерный вектор, описывающий координаты луча
	direction [2]int32     // Двухмерный вектор, описывающий направление луча
	speed     int32        // Скорость луча
	trail     []rl.Vector2 // Вектор перемещения луча для отрисовки следа
}

// Создает новый луч с заданными параметрами
func NewRay(p, d [2]int32) *Ray {
	r := &Ray{
		position:  p,
		direction: d,
		speed:     c,
		trail:     make([]rl.Vector2, 0, 200),
	}
	r.trail = append(r.trail, rl.Vector2{X: float32(p[0]), Y: float32(p[1])})

	return r
}

// Обновляет позицию луча на основе его направления и скорости
func (r *Ray) Update(scale int32) {
	// Двигаем луч
	r.position[0] += r.direction[0] * r.speed / scale
	r.position[1] += r.direction[1] * r.speed / scale
	// Переписываем трейл
	r.trail = append(r.trail, rl.Vector2{X: float32(r.position[0]), Y: float32(r.position[1])})
	if len(r.trail) > 200 {
		r.trail = r.trail[1:]
	}
}

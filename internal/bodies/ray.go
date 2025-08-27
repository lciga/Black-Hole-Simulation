package bodies

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ray struct {
	position  [2]float64   // Текущая позиция луча
	prev      [2]float64   // Предыдущая позиция луча
	phi       float64      // Угол оси X, полярные координаты
	radius    float64      // Радиус от начала координат, полярные координаты
	direction [2]float64   // Двухмерный вектор, описывающий направление луча
	speed     float64      // Скорость луча
	trail     []rl.Vector2 // Вектор перемещения луча для отрисовки следа
	absorbed  bool         // Поглощен ли луч черной дырой
}

// Создает новый луч с заданными параметрами
func NewRay(p, d [2]float64) *Ray {
	r := &Ray{
		position:  p,
		direction: d,
		prev:      p,
		radius:    math.Hypot(p[0], p[1]),
		phi:       math.Atan2(p[1], p[0]),
		speed:     c,
		trail:     make([]rl.Vector2, 0, 200),
	}
	r.trail = append(r.trail, rl.Vector2{X: float32(p[0]), Y: float32(p[1])})

	return r
}

// Обновляет позицию луча на основе его направления и скорости.
// Проверяет, не поглощен ли луч черной дырой. pxPerMeter - такой же как масштаб чёрной дыры
func (r *Ray) Update(scale float64, pxPerMeter float64, bh *BlackHole) {
	// Проверяем, не поглощен ли луч черной дырой
	if r.absorbed {
		return
	}

	// Сохраняем предыдущую позицию
	r.prev = r.position

	// Двигаем луч
	r.position[0] += r.direction[0] * r.speed / scale
	r.position[1] += r.direction[1] * r.speed / scale

	// Проверяем на поглощение черной дырой
	rsPx := bh.rs * pxPerMeter
	if segmentIntersectsCircle(r.prev, r.position, bh.position, rsPx) {
		r.absorbed = true
		return
	}

	// Переписываем трейл
	r.trail = append(r.trail, rl.Vector2{X: float32(r.position[0]), Y: float32(r.position[1])})
	if len(r.trail) > 200 {
		r.trail = r.trail[1:]
	}
}

// Считаем расстояние от центра окружности до ближайшей точки отрезка
func segmentIntersectsCircle(p1, p2, c [2]float64, rad float64) bool {
	// Вектор d = p1 - p2
	dx, dy := p2[0]-p1[0], p2[1]-p1[1]
	// Если луч не двигается - проверка точечная
	den := dx*dx + dy*dy
	if den == 0 {
		return dist(p1, c) <= rad
	}
	// Проекция на отрезок
	fx, fy := p1[0]-c[0], p1[1]-c[1]
	t := -(fx*dx + fy*dy) / den
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}
	// Ближайшая точка на отрезке
	closestX := p1[0] + t*dx
	closestY := p1[1] + t*dy
	cx, cy := closestX-c[0], closestY-c[1]
	return cx*cx+cy*cy <= rad*rad
}

func dist(a, b [2]float64) float64 {
	return math.Hypot(a[0]-b[0], a[1]-b[1])
}

// Пакет отвечает за создание движка отрисвки графики
package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Структура описывающая 2D движок
type Engine2D struct {
	Width     uint
	Height    uint
	Title     string
	Bacground rl.Color
}

// Функция создания нового 2D движка с заданными параметрами
func New2DEngine(width, height uint, title string, color rl.Color) *Engine2D {
	return &Engine2D{
		Width:     width,
		Height:    height,
		Title:     title,
		Bacground: color,
	}
}

// Инициализация 2D движка
func (e *Engine2D) Init(drawFunc func()) {
	rl.InitWindow(int32(e.Width), int32(e.Height), e.Title)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.SetTargetFPS(60) // Ограничение до 60 кадров в секунду
		rl.ClearBackground(e.Bacground)
		drawFunc()

		rl.EndDrawing()
	}
}

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	InitWindowWidth  = 800
	InitWindowHeight = 450
	TargetFPS        = 1024
)

func main() {
	rl.InitWindow(InitWindowWidth, InitWindowHeight, "RhythmGameWithTheWhiteCube")
	rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()
	rl.SetTargetFPS(TargetFPS)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(0, 0, 100, 100, rl.RayWhite)

		rl.EndDrawing()
	}
}

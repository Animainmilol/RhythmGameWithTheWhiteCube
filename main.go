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

	cameraController := NewCameraController()

	for !rl.WindowShouldClose() {
		cameraController.Update()
		cameraController.HandleInput()

		rl.BeginDrawing()
		rl.BeginMode2D(cameraController.Camera)

		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(0, 0, 100, 100, rl.RayWhite)

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

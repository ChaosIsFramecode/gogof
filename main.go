package main

import (
	"math/rand"
	"time"

	"github.com/ChaosIsFramecode/gogof/config"
	"github.com/ChaosIsFramecode/gogof/grid"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func remove(slice []rl.Vector2, s int) []rl.Vector2 {
	return append(slice[:s], slice[s+1:]...)
}

func randGen(num int32, randGrid *rand.Rand) []rl.Vector2 {
	var arr []rl.Vector2
	for i := 0; i < int(num); i++ {
		arr = append(arr, rl.Vector2{X: float32(randGrid.Int31n(config.GridHeight)), Y: float32(randGrid.Int31n(config.GridWidth))})
	}
	return arr
}

func main() {
	rl.InitWindow(config.Width, config.Height, "gogof")
	defer rl.CloseWindow()
	rl.SetTargetFPS(config.FPS)

	// Setup game states
	var positions []rl.Vector2
	playing := false
	randGrid := rand.New(rand.NewSource(time.Now().Unix()))

	count := 0
	stepFreq := 30

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if playing {
			count++
		}
		if count >= stepFreq {
			count = 0
			positions = grid.Adjust(positions)
		}

		rl.ClearBackground(rl.Black)
		grid.Draw(positions)

		// Mouse Input
		if rl.IsMouseButtonPressed(0) {
			mousePos := rl.GetMousePosition()
			mousePos.X = float32(int32(mousePos.X) / config.TileSize)
			mousePos.Y = float32(int32(mousePos.Y) / config.TileSize)

			ok := false
			for i := 0; i < len(positions); i++ {
				if positions[i].X == mousePos.X && positions[i].Y == mousePos.Y {
					ok = true
					positions = remove(positions, i)
					break
				}
			}
			if !ok {
				positions = append(positions, mousePos)
			}
		}

		// Key input
		if rl.IsKeyPressed(rl.KeySpace) {
			playing = !playing
		}
		if rl.IsKeyPressed(rl.KeyC) {
			positions = []rl.Vector2{}
			playing = false
		}
		if rl.IsKeyPressed(rl.KeyG) {
			positions = randGen(randGrid.Int31n(6)*config.GridWidth, randGrid)
		}

		rl.EndDrawing()
	}
}

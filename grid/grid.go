package grid

import (
	"github.com/ChaosIsFramecode/gogof/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw(posiitons []rl.Vector2) {
	for _, pos := range posiitons {
		pos.X *= float32(config.TileSize)
		pos.Y *= float32(config.TileSize)
		rl.DrawRectangleV(pos, rl.Vector2{X: float32(config.TileSize), Y: float32(config.TileSize)}, rl.White)
	}

	// Row
	for i := 0; int32(i) < config.GridHeight; i++ {
		rl.DrawLine(0, int32(i)*config.TileSize, config.Width, int32(i)*config.TileSize, rl.White)
	}
	// Col
	for i := 0; int32(i) < config.GridHeight; i++ {
		rl.DrawLine(int32(i)*config.TileSize, 0, int32(i)*config.TileSize, config.Height, rl.White)
	}
}

func Adjust(positions []rl.Vector2) []rl.Vector2 {
	// Create a map for quick lookup of positions
	positionSet := make(map[rl.Vector2]struct{})
	for _, pos := range positions {
		positionSet[pos] = struct{}{}
	}

	// Filter function to get existing neighbors
	filterNeighbors := func(neighbors []rl.Vector2) []rl.Vector2 {
		filtered := []rl.Vector2{}
		for _, neighbor := range neighbors {
			if _, exists := positionSet[neighbor]; exists {
				filtered = append(filtered, neighbor)
			}
		}
		return filtered
	}

	allNeighbors := make(map[rl.Vector2]struct{})
	newPos := make(map[rl.Vector2]struct{})

	// Check current positions
	for _, v := range positions {
		neighbors := filterNeighbors(GetNeighbors(v))
		if len(neighbors) == 2 || len(neighbors) == 3 {
			newPos[v] = struct{}{}
		}
		for _, neighbor := range GetNeighbors(v) {
			allNeighbors[neighbor] = struct{}{}
		}
	}

	// Check all neighboring positions
	for v := range allNeighbors {
		if _, exists := positionSet[v]; !exists { // only check for positions not in the original set
			neighbors := filterNeighbors(GetNeighbors(v))
			if len(neighbors) == 3 {
				newPos[v] = struct{}{}
			}
		}
	}

	result := []rl.Vector2{}
	for v := range newPos {
		result = append(result, v)
	}

	return result
}

func GetNeighbors(pos rl.Vector2) []rl.Vector2 {
	var neighbors []rl.Vector2

	d := []int{-1, 0, 1}

	for _, dx := range d {
		if int(pos.X)+dx < 0 || int(pos.X)+dx > int(config.GridWidth) {
			continue
		}

		for _, dy := range d {
			if int(pos.Y)+dy < 0 || int(pos.Y)+dy > int(config.GridHeight) {
				continue
			}
			if dx == 0 && dy == 0 {
				continue
			}

			neighbors = append(neighbors, rl.Vector2{X: pos.X + float32(dx), Y: pos.Y + float32(dy)})
		}
	}

	return neighbors
}

package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func collisionCheck(a *Entity, b *Entity) bool {
	leftB := b.pos.X - BUFFER
	rightB := b.pos.X + b.size.X + BUFFER
	topB := b.pos.Y - BUFFER
	bottomB := b.pos.Y + b.size.Y + BUFFER

	if a.pos.X <= rightB &&
		a.pos.X >= leftB &&
		a.pos.Y <= bottomB &&
		a.pos.Y >= topB {

		return true
	}

	return false
}

func move(entity *Entity, direction float32) {
	entity.pos.Y += float32(direction) * (float32(entity.speed) * rl.GetFrameTime())
	entity.pos.Y = rl.Clamp(entity.pos.Y, 0, SCREEN_H - entity.size.Y)
}

func generateRandomDirection(entity *Entity) {
	entity.direction = rl.NewVector2(float32(rand.Intn(2)-1), float32(rand.Intn(2)-1))

	if entity.direction.X == 0 {
		entity.direction.X = 1
	}

	if entity.direction.Y == 0 {
		entity.direction.Y = 1
	}
}
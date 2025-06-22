package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func collisionCheck(a *Object, b *Entity, col *Collision) bool {
	// get the borders of the entity
	leftB := b.pos.X
	rightB := b.pos.X + b.size.X
	topB := b.pos.Y
	bottomB := b.pos.Y + b.size.Y

	// get the borders of the object
	leftA := a.pos.X - b.size.X + a.size.X
	rightA := a.pos.X + b.size.X - a.size.X
	topA := a.pos.Y + a.size.Y * 2
	bottomA := a.pos.Y - a.size.Y * 2
	

	if leftA <= rightB &&
	rightA >= leftB &&		
	bottomA <= bottomB && 
	topA >= topB {

		col.point.Y = ((b.pos.Y + b.size.Y / 2) - (a.pos.Y + a.size.Y / 2)) * 0.01

		fmt.Println(col.point.Y)

		return true
	}

	return false
}

func move(entity *Entity, direction float32) {
	entity.directionY = direction

	entity.pos.Y += float32(entity.directionY) * (float32(entity.speed) * rl.GetFrameTime())
	entity.pos.Y = rl.Clamp(entity.pos.Y, 0, SCREEN_H - entity.size.Y)
}

func generateRandomDirection(object *Object) {
	object.direction = rl.NewVector2(float32(rand.Intn(2)*2-1), float32(rand.Intn(2)*2-1))
	// object.direction = rl.NewVector2(float32(rand.Intn(2)*2-1), rand.Float32() * 2 - 1)

	// if object.direction.X == 0 {
	// 	object.direction.X = 1
	// }

	if object.direction.Y == 0 {
		object.direction.Y = float32(rand.Intn(2)*2-1)
	}
}
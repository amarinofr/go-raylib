package main

func collisionCheck(a *Entity, b *Entity) bool {
	return (a.pos.X-a.width <= b.pos.X+b.width &&
		a.pos.X+a.width >= b.pos.X &&
		a.pos.Y <= b.pos.Y+b.height &&
		a.pos.Y+a.height >= b.pos.Y)
}

func move(entity *Entity, direction float32) {
	entity.pos.Y += float32(direction) * float32(entity.speed)
}

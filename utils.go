package main

func collisionCheck(a *Entity, b *Entity) bool {
	return (a.pos.X-a.width <= b.pos.X+b.width &&
		a.pos.X+a.width >= b.pos.X &&
		a.pos.Y <= b.pos.Y+b.height &&
		a.pos.Y+a.height >= b.pos.Y)
}

func reset(a *Entity, b *Entity, c *Entity) {
	a.pos.X = SCREEN_W / 2
	a.pos.Y = SCREEN_H / 2
	b.pos.Y = SCREEN_H/2 - b.height/2
	c.pos.Y = SCREEN_H/2 - c.height/2
}

func move(entity *Entity, direction float32) {
	entity.pos.Y += float32(direction) * float32(entity.speed)
}
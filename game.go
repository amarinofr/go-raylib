package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Init() {
	rl.InitWindow(SCREEN_W, SCREEN_H, TITLE)
	rl.SetTargetFPS(FPS)

	ball.direction = float32(rand.Intn(2) - 1)
	if ball.direction == 0 {
		ball.direction = 1
	}

	g.running = true
	g.reset = false
}

func (g *Game) Close() {
	rl.CloseWindow()
}

func (g *Game) Loop() {
	for g.running {
		if rl.WindowShouldClose() {
			g.Close()
			g.running = false
			return
		}

		g.update()
		g.draw()

		input(&player, g)
	}
}

func reset(a *Entity, b *Entity, c *Entity) {
	a.pos.X = SCREEN_W / 2
	a.pos.Y = SCREEN_H / 2
	b.pos.Y = SCREEN_H/2 - b.height/2
	c.pos.Y = SCREEN_H/2 - c.height/2
}

func input(entity *Entity, g *Game) {
	if !g.reset {
		if rl.IsKeyDown(rl.KeyUp) {
			if !(entity.pos.Y < 0) {
				move(entity, float32(-1))
			}
		}

		if rl.IsKeyDown(rl.KeyDown) {
			if !(entity.pos.Y > SCREEN_H-entity.height) {
				move(entity, float32(1))
			}
		}
	}

	if g.reset {
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.reset = false
			ball.direction = float32(rand.Intn(2) - 1)

			if ball.direction == 0 {
				ball.direction = 1
			}
		}
	}
}

func checkBounds(ball *Entity, enemy *Entity, player *Entity, g *Game) {
	top := 0
	bottom := SCREEN_H
	left := 0
	right := SCREEN_W

	// if (collisionCheck(ball, player)) {
	// 	ball.direction = 1
	// }

	// if (collisionCheck(ball, enemy)) {
	// 	ball.direction = -1
	// }
}

func ballBehaviors(ball *Entity, enemy *Entity, player *Entity, g *Game) {

	if ball.pos.X < 0 {
		g.reset = true
		reset(ball, enemy, player)
		enemy.score++
	}

	if ball.pos.X > SCREEN_W-ball.width {
		g.reset = true
		reset(ball, enemy, player)
		player.score++
	}

	if g.reset {
		ball.direction = 0
	}

	ball.pos.X += float32(ball.direction) * float32(ball.speed)
}

func (g *Game) update() {
	rl.DrawFPS(10, SCREEN_H-20)
	ballBehaviors(&ball, &enemy, &player, g)
}

func (g *Game) draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	rl.DrawRectangle(int32(player.pos.X), int32(player.pos.Y), int32(player.width), int32(player.height), rl.Red)
	rl.DrawRectangle(int32(enemy.pos.X), int32(enemy.pos.Y), int32(enemy.width), int32(enemy.height), rl.Blue)
	rl.DrawCircle(int32(ball.pos.X), int32(ball.pos.Y), 10, rl.White)

	rl.DrawText(fmt.Sprintf("%d", player.score), 10, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("%d", enemy.score), SCREEN_W-20, 10, 20, rl.White)

	if g.reset {
		rl.DrawText("Press Enter to reset", SCREEN_W/2-100, 10, 20, rl.White)
	}

	rl.EndDrawing()
}

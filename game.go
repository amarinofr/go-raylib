package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Init() {
	rl.InitWindow(SCREEN_W, SCREEN_H, TITLE)
	rl.SetTargetFPS(FPS)

	generateRandomDirection(&ball)

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

func input(entity *Entity, g *Game) {
	if !g.reset {
		if rl.IsKeyDown(rl.KeyUp) {
			move(entity, float32(-1))
		}

		if rl.IsKeyDown(rl.KeyDown) {
			move(entity, float32(1))
		}
	}

	if g.reset {
		if rl.IsKeyPressed(rl.KeyEnter) {
			generateRandomDirection(&ball)

			g.reset = false
		}
	}
}

func reset(a *Entity, b *Entity, c *Entity, g *Game) {
	if !g.reset {
		g.reset = true
	}

	a.direction = rl.NewVector2(0, 0)
	a.pos.X = SCREEN_W / 2
	a.pos.Y = SCREEN_H / 2

	b.pos.Y = SCREEN_H/2 - b.size.Y/2
	c.pos.Y = SCREEN_H/2 - c.size.Y/2

}


func checkBounds(ball *Entity, enemy *Entity, player *Entity, g *Game) {
	if ball.pos.X <= BOUNDS_LEFT {
		// ball.direction.X = 1 // used for debbuging
		reset(ball, enemy, player, g)
		enemy.score++
	}

	if ball.pos.X >= BOUNDS_RIGHT {
		// ball.direction.X = -1 // used for debbuging
		reset(ball, enemy, player, g)
		player.score++
	}

	if ball.pos.Y - ball.size.Y <= BOUNDS_TOP {
		ball.direction.Y = 1
	}

	if ball.pos.Y + ball.size.Y * 2 >= BOUNDS_BOTTOM {
		ball.direction.Y = -1
	}
}

func ballBehaviors(ball *Entity, enemy *Entity, player *Entity) {
	if collisionCheck(ball, enemy) {
		ball.direction.X = -1
	}

	if collisionCheck(ball, player) {
		ball.direction.X = 1
	}

	ball.pos.X += ball.direction.X * float32(ball.speed) * rl.GetFrameTime()
	ball.pos.Y += ball.direction.Y * float32(ball.speed) * rl.GetFrameTime()
}

func moveEnemy(ball *Entity, enemy *Entity) {
	if ball.pos.Y > enemy.pos.Y + (enemy.size.Y/2 + 20) {
		move(enemy, 1)
	}

	if ball.pos.Y < enemy.pos.Y + (enemy.size.Y/2 - 20) {
		move(enemy, -1)
	}
}

func (g *Game) update() {
	rl.DrawFPS(10, SCREEN_H-20)
	ballBehaviors(&ball, &enemy, &player)
	checkBounds(&ball, &enemy, &player, g)
	moveEnemy(&ball, &enemy)
}

func (g *Game) draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	rl.DrawRectangle(int32(player.pos.X), int32(player.pos.Y), int32(player.size.X), int32(player.size.Y), rl.Red)
	rl.DrawRectangle(int32(enemy.pos.X), int32(enemy.pos.Y), int32(enemy.size.X), int32(enemy.size.Y), rl.Blue)
	rl.DrawCircle(int32(ball.pos.X), int32(ball.pos.Y), 10, rl.White)

	rl.DrawText(fmt.Sprintf("%d", player.score), 10, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("%d", enemy.score), SCREEN_W-20, 10, 20, rl.White)

	if g.reset {
		rl.DrawText("Press Enter to reset", SCREEN_W/2-100, 10, 20, rl.White)
	}

	rl.EndDrawing()
}

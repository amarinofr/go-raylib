package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	running bool
	reset   bool
}

type Entity struct {
	pos       rl.Vector2
	speed     int16
	size      rl.Vector2
	score     int
	center rl.Vector2
	directionY float32
}

type Object struct {
	pos       rl.Vector2
	speed     float64
	size      rl.Vector2
	direction rl.Vector2
	center rl.Vector2
}

type Collision struct {
	point rl.Vector2
}
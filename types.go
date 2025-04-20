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
	direction rl.Vector2
	score     int
	center rl.Vector2
}

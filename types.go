package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	running bool
	reset   bool
}

type Entity struct {
	pos       rl.Vector2
	speed     int16
	width     float32
	height    float32
	direction float32
	score     int
}

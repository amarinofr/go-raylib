package main

import rl "github.com/gen2brain/raylib-go/raylib"

var player = Entity{
	pos:    rl.NewVector2(0, SCREEN_H/2-50),
	speed:  5,
	width:  20,
	height: 100,
	center: rl.NewVector2(20/2, 100/2),
}

var enemy = Entity{
	pos:    rl.NewVector2(SCREEN_W-20, SCREEN_H/2-50),
	speed:  5,
	width:  20,
	height: 100,
	score:  0,
	center: rl.NewVector2(20/2, 100/2),
}

var ball = Entity{
	pos:    rl.NewVector2(SCREEN_W/2-5, SCREEN_H/2-5),
	speed:  7,
	width:  5,
	height: 5,
	score:  0,
	direction: rl.NewVector2(0,0),
	center: rl.NewVector2(5/2, 5/2),
}

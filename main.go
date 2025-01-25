package main

func main() {
	game := Game{}

	game.Init()
	defer game.Close()

	game.Loop()
}
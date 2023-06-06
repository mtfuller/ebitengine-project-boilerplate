package loaders

import (
	"log"
	"github.com/yourname/yourgame/engine"
)

var Player *engine.Spritesheet
var Wall *engine.Spritesheet

func Load() {
	var err error

	Player = engine.NewSpritesheet()

	sprite := engine.Rectangle{
		X: 0,
		Y: 0,
		W: 8,
		H: 8,
	}

	err = Player.AddSprite("test", "assets/sprites/man.png", sprite, 4, 4)
	if err != nil {
		log.Fatal("HELLO")
	}

	Wall = engine.NewSpritesheet()

	sprite = engine.Rectangle{
		X: 0,
		Y: 0,
		W: 32,
		H: 32,
	}

	err = Wall.AddSprite("test", "assets/sprites/wall.png", sprite, 1, 1)
	if err != nil {
		log.Fatal("HELLO2")
	}
}
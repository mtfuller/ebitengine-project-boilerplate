package engine

import (
	"fmt"
	"image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Rectangle struct {
	X, Y int
	W, H int
}

type Spritesheet struct {
	sprites map[string][]*ebiten.Image
}

func NewSpritesheet() *Spritesheet {
	spritesheet := &Spritesheet{}

	spritesheet.sprites = make(map[string][]*ebiten.Image)

	return spritesheet
}

func (s *Spritesheet) AddSprite(name string, src string, rect Rectangle, num int, scale int) error {
	var sprite []*ebiten.Image

	img, _, err := ebitenutil.NewImageFromFile(src)
	if err != nil {
		return err
	}

	for n := 0; n < num; n++ {
		x1 := rect.X + (n * rect.W)
		x2 := x1 + rect.W
		y1 := rect.Y
		y2 := y1 + rect.H

		fmt.Println(x1, y1, x2, y2)

		subimg := img.SubImage(image.Rect(x1,y1,x2,y2))

		i := ebiten.NewImageFromImage(subimg)
		newImg := ebiten.NewImage(rect.W * scale, rect.H * scale)
	
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Scale(float64(scale), float64(scale))
		newImg.DrawImage(i, op)
	
		sprite = append(sprite, newImg)
	}

	s.sprites[name] = sprite

	return nil
}

func (s Spritesheet) GetSprite(name string) []*ebiten.Image {
	return s.sprites[name]
}

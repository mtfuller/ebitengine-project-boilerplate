package framework

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

type Sprite struct {
	Frames    []*ebiten.Image
	Durations []int
}

type SpritesheetKey struct {
	EntityName string
	TileName   string
}

type Spritesheet struct {
	Sprites map[string]map[string]Sprite
	IdMap   map[uint32]SpritesheetKey
}

func getFrame(tileset *tiled.Tileset, tilesetImage *ebiten.Image, id uint32) *ebiten.Image {
	rect := tileset.GetTileRect(id)
	subImg := tilesetImage.SubImage(rect)
	frame := ebiten.NewImageFromImage(subImg)
	return frame
}

func LoadSpritesheetFromTileset(tileset *tiled.Tileset) Spritesheet {
	spritesheet := Spritesheet{
		Sprites: make(map[string]map[string]Sprite),
		IdMap:   make(map[uint32]SpritesheetKey),
	}

	filePath := tileset.GetFileFullPath(tileset.Image.Source)
	tilesetImage, _, err := ebitenutil.NewImageFromFile(filePath)
	if err != nil {
		panic(err)
	}

	for _, tile := range tileset.Tiles {
		tileType := tile.Type

		_, ok := spritesheet.Sprites[tileType]
		if !ok {
			spritesheet.Sprites[tileType] = make(map[string]Sprite)
		}

		frames := []*ebiten.Image{}
		durations := []int{}
		if len(tile.Animation) > 0 {
			for _, animation := range tile.Animation {
				frame := getFrame(tileset, tilesetImage, animation.TileID)
				frames = append(frames, frame)

				duration := int((100 / 1000.0) * float32(animation.Duration))
				durations = append(durations, duration)
			}
		} else {
			frame := getFrame(tileset, tilesetImage, tile.ID)

			frames = append(frames, frame)
			durations = append(durations, 1)
		}

		spriteId := fmt.Sprint(tile.ID)
		spritesheet.IdMap[tile.ID] = SpritesheetKey{
			EntityName: tileType,
			TileName:   spriteId,
		}

		spritesheet.Sprites[tileType][spriteId] = Sprite{
			Frames:    frames,
			Durations: durations,
		}
	}

	return spritesheet
}

package framework

import (
	"fmt"
	"sort"

	"github.com/lafriks/go-tiled"
)

type BoardMap struct {
	tilemap             *tiled.Map
	gidRanges           []int
	gidToSpritesheetMap map[uint32]*Spritesheet
}

type BoardMapTile struct {
	Type        string
	Spritesheet *Spritesheet
	SpriteName  string
	X           int
	Y           int
	W           int
	H           int
}

func LoadBoardMapFromTilemap(tilemapFilepath string) (BoardMap, error) {
	tilemap, err := tiled.LoadFile(tilemapFilepath)
	if err != nil {
		return BoardMap{}, err
	}

	ranges := []int{}
	spritesheetMap := make(map[uint32]*Spritesheet)
	for _, tileset := range tilemap.Tilesets {
		firstGid := tileset.FirstGID
		ranges = append(ranges, int(firstGid))
		spriteSheet := GetAssetManager().GetSpritesheetFromTileset(tileset)
		spritesheetMap[firstGid] = spriteSheet
	}
	sort.Ints(ranges)

	return BoardMap{
		tilemap:             tilemap,
		gidRanges:           ranges,
		gidToSpritesheetMap: spritesheetMap,
	}, nil
}

func (b BoardMap) GetAllTiles() []BoardMapTile {
	tiles := []BoardMapTile{}

	for row := 0; row < b.tilemap.Height; row++ {
		for col := 0; col < b.tilemap.Width; col++ {
			x := col * b.tilemap.TileWidth
			y := row * b.tilemap.TileHeight

			tileId := (row * b.tilemap.Width) + col
			tile := b.tilemap.Layers[0].Tiles[tileId]

			if !tile.IsNil() {
				tileset := tile.Tileset

				spriteSheet := GetAssetManager().GetSpritesheetFromTileset(tileset)

				_, ok := b.gidToSpritesheetMap[tileset.FirstGID]
				if !ok {
					b.gidToSpritesheetMap[tileset.FirstGID] = spriteSheet
				}

				tstile, err := tileset.GetTilesetTile(tile.ID)
				if err != nil {
					panic(err)
				}

				tiles = append(tiles, BoardMapTile{
					Type:        tstile.Type,
					Spritesheet: spriteSheet,
					SpriteName:  fmt.Sprint(tile.ID),
					X:           x,
					Y:           y,
					W:           b.tilemap.Width,
					H:           b.tilemap.Height,
				})
			}
		}
	}

	return tiles
}

func (b BoardMap) GetAllObjects() []BoardMapTile {
	objs := []BoardMapTile{}

	for _, obj := range b.tilemap.ObjectGroups[0].Objects {
		fmt.Println("ID:", obj.ID)
		fmt.Println("GID:", obj.GID)

		spritesheetGID := b.getSpritesheetGID(int(obj.GID))
		spritesheet := b.gidToSpritesheetMap[uint32(spritesheetGID)]

		ss := *spritesheet

		spriteId := obj.GID - spritesheetGID

		fmt.Println("Spritesheet GID:", spritesheetGID)
		fmt.Println("Sprite ID:", spriteId)
		fmt.Println("Spritesheet:", spritesheet)
		fmt.Println("Spritesheet:", ss.Sprites["CHAR1"]["0"])
		fmt.Println("Class:", obj.Class)
		fmt.Println("Type:", ss.IdMap[spriteId].EntityName)
		fmt.Println("X:", obj.X)
		fmt.Println("Y:", obj.Y)

		objs = append(objs, BoardMapTile{
			Type:        ss.IdMap[spriteId].EntityName,
			Spritesheet: spritesheet,
			SpriteName:  ss.IdMap[spriteId].TileName,
			X:           int(obj.X),
			Y:           int(obj.Y),
			W:           b.tilemap.Width,
			H:           b.tilemap.Height,
		})

	}

	return objs
}

func (b BoardMap) getSpritesheetGID(gid int) uint32 {
	for _, firstGid := range b.gidRanges {
		if gid <= firstGid {
			return uint32(firstGid)
		}
	}

	return uint32(b.gidRanges[len(b.gidRanges)-1])
}

// func (b BoardMap) getSpritesheetFromGID(gid int) (*Spritesheet, uint32) {
// 	var spritesheet *Spritesheet = nil

// 	for _, firstGid := range b.gidRanges {
// 		//fmt.Println("WOW:", firstGid)
// 		if gid > firstGid {
// 			continue
// 		} else {
// 			fmt.Println("WOW", firstGid)
// 			spritesheet = b.gidToSpritesheetMap[uint32(firstGid)]
// 			break
// 		}
// 	}

// 	return spritesheet
// }

// 1 181
// 196

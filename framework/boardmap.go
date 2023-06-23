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

type BoardMapEntity struct {
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

func (b BoardMap) GetAllEntities() []BoardMapEntity {
	tiles := b.getAllTiles()
	objects := b.getAllObjects()

	return append(tiles, objects...)
}

func (b BoardMap) getAllTiles() []BoardMapEntity {
	tiles := []BoardMapEntity{}

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

				tiles = append(tiles, BoardMapEntity{
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

func (b BoardMap) getAllObjects() []BoardMapEntity {
	objs := []BoardMapEntity{}

	for _, obj := range b.tilemap.ObjectGroups[0].Objects {
		spritesheetGID := b.getSpritesheetGID(int(obj.GID))
		spritesheet := b.gidToSpritesheetMap[uint32(spritesheetGID)]

		ss := *spritesheet

		spriteId := obj.GID - spritesheetGID

		objs = append(objs, BoardMapEntity{
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

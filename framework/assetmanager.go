package framework

import "github.com/lafriks/go-tiled"

type AssetManager struct {
	tilesetToSpritesheetMap map[string]*Spritesheet
}

var assetManager *AssetManager = nil

func GetAssetManager() AssetManager {
	if assetManager == nil {
		assetManager = &AssetManager{
			tilesetToSpritesheetMap: make(map[string]*Spritesheet),
		}
	}

	return *assetManager
}

func (a AssetManager) GetSpritesheetFromTileset(tileset *tiled.Tileset) *Spritesheet {
	_, ok := a.tilesetToSpritesheetMap[tileset.Source]
	if !ok {
		spritesheet := LoadSpritesheetFromTileset(tileset)
		a.tilesetToSpritesheetMap[tileset.Source] = &spritesheet
	}

	return a.tilesetToSpritesheetMap[tileset.Source]
}

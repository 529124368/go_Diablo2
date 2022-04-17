package maps

import (
	"game/mapCreator/d2interface"
	"game/mapCreator/ds1"
	"game/mapCreator/dt1"
	"game/tools"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func getTitleImage(tileData dt1.Tile, w d2interface.Palette) *ebiten.Image {
	tileYMinimum := int32(0)
	for _, block := range tileData.Blocks {
		tileYMinimum = tools.MinInt32(tileYMinimum, int32(block.Y))
	}
	tileYOffset := tools.AbsInt32(tileYMinimum)
	tileHeight := tools.AbsInt32(tileData.Height)
	indexData := make([]byte, tileData.Width*int32(tileHeight))
	dt1.DecodeTileGfxData(tileData.Blocks, &indexData, tileYOffset, tileData.Width)
	//加载调色板
	pixels := dt1.ImgIndexToRGBA(indexData, w)
	imgss := ebiten.NewImage(int(tileData.Width), int(tileHeight))
	imgss.ReplacePixels(pixels)
	return imgss
}

//获取wall层贴图
func getWallTitleImage(tileData dt1.Tile, tile *ds1.Tile, w d2interface.Palette) (*ebiten.Image, int) {

	tileMinY := int32(0)
	tileMaxY := int32(0)
	for _, block := range tileData.Blocks {

		tileMinY = tools.MinInt32(tileMinY, int32(block.Y))
		tileMaxY = tools.MaxInt32(tileMaxY, int32(block.Y+32))
	}

	realHeight := tools.MaxInt32(tools.AbsInt32(tileData.Height), tileMaxY-tileMinY)
	tileYOffset := -tileMinY

	if tile.Type == d2enum.TileRoof {
		tile.YAdjust = -int(tileData.RoofHeight)
	} else {
		tile.YAdjust = int(tileMinY) + 80
	}

	indexData := make([]byte, 160*realHeight)
	dt1.DecodeTileGfxData(tileData.Blocks, &indexData, tileYOffset, 160)
	//加载调色板
	pixels := dt1.ImgIndexToRGBA(indexData, w)
	imgss := ebiten.NewImage(160, int(realHeight))
	imgss.ReplacePixels(pixels)
	return imgss, tile.YAdjust
}

//根据ds1 获取对应dt1
func GetTiles(style, sequence int, tileType d2enum.TileType, m []dt1.Tile) []dt1.Tile {
	tiles := make([]dt1.Tile, 0)

	for idx := range m {
		if m[idx].Style != int32(style) || m[idx].Sequence != int32(sequence) ||
			m[idx].Type != int32(tileType) {
			continue
		}
		tiles = append(tiles, m[idx])
	}
	if len(tiles) == 0 {
		return nil
	}
	return tiles
}

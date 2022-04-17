package maps

import (
	"embed"
	"fmt"
	"game/mapCreator/d2interface"
	"game/mapCreator/dat"
	"game/mapCreator/ds1"
	"game/mapCreator/dt1"
	"game/tools"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	MAPOFFSETX float64 = -1800
	MAPOFFSETY float64 = -1300
	Scale      float64 = 1
)

type imgWall struct {
	img *ebiten.Image
	h   int
}

var img [][]*ebiten.Image

var img2 [][]imgWall

var img3 [][]imgWall

type MapBase struct {
	image   *embed.FS
	OpBg    *ebiten.DrawImageOptions
	BgImage *ebiten.Image
}

//Create Map Class
func NewMap(images *embed.FS) *MapBase {
	maps := &MapBase{
		image: images,
	}
	return maps
}

//加载地图图片
func (m *MapBase) LoadMap() {
	//加载静态地图
	//m.LoadStaticMap()
	//加载动态地图
	m.LoadDstMap()
}

//改变地图坐标
func (m *MapBase) ChangeMapTranslate(x, y float64) {
	m.OpBg.GeoM.Translate(x, y)
}

func (m *MapBase) Render(screen *ebiten.Image, offsetX, offsetY float64) {
	//screen.DrawImage(m.BgImage, m.OpBg)
	//floor
	sumX := 0
	startY := 0
	for i := 0; i < 41; i++ {
		startY += 40
		sumX = 0
		for j := 0; j < 57; j++ {
			s := img[i][j]
			sumX += 80
			if s != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY)
				op.GeoM.Scale(Scale, Scale)
				screen.DrawImage(s, op)
			}
		}
	}

	//补图
	sumX = 0
	startY = 0
	for i := 0; i < 41; i++ {
		startY += 40
		sumX = 0
		for j := 0; j < 57; j++ {
			s := img3[i][j].img
			sumX += 80
			if s != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY+float64(img3[i][j].h))
				op.GeoM.Scale(Scale, Scale)
				screen.DrawImage(s, op)
			}
		}
	}
	//wall
	sumX = 0
	startY = 0
	for i := 0; i < 41; i++ {
		startY += 40
		sumX = 0
		for j := 0; j < 57; j++ {
			s := img2[i][j].img
			sumX += 80
			if s != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY+float64(img2[i][j].h))
				op.GeoM.Scale(Scale, Scale)
				screen.DrawImage(s, op)
			}
		}
	}
}

//静态地图
func (m *MapBase) LoadStaticMap() {
	s2, _ := m.image.ReadFile("resource/bg/old.png")
	img := tools.GetEbitenImage(s2)
	m.BgImage = img
	m.OpBg = &ebiten.DrawImageOptions{}
	m.OpBg.Filter = ebiten.FilterLinear
	m.OpBg.GeoM.Translate(MAPOFFSETX, MAPOFFSETY)
}

//静态动态地图
func (m *MapBase) LoadDstMap() {
	//加载调色板
	r, _ := m.image.ReadFile(tools.ObjectPath + "/mapSucai/pal.dat")
	ww, _ := dat.Load(r)
	//加载地块dt1素材
	re, _ := m.image.ReadFile(tools.ObjectPath + "/mapSucai/floor.dt1")
	ss, _ := dt1.LoadDT1(re)
	re, _ = m.image.ReadFile(tools.ObjectPath + "/mapSucai/objects.dt1")
	ss1, _ := dt1.LoadDT1(re)

	re, _ = m.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/objects.dt1")
	ss2, _ := dt1.LoadDT1(re)
	re, _ = m.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/treegroups.dt1")
	ss3, _ := dt1.LoadDT1(re)

	re, _ = m.image.ReadFile(tools.ObjectPath + "/mapSucai/fence.dt1")
	ss4, _ := dt1.LoadDT1(re)

	re, _ = m.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/bridge.dt1")
	ss5, _ := dt1.LoadDT1(re)

	re, _ = m.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/stonewall.dt1")
	ss6, _ := dt1.LoadDT1(re)

	re, _ = m.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/river.dt1")
	ss7, err := dt1.LoadDT1(re)

	//wall
	ss2.Tiles = append(ss2.Tiles, ss1.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss3.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss4.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss5.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss6.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss7.Tiles...)

	//floor
	ss.Tiles = append(ss.Tiles, ss2.Tiles...)

	if err != nil {
		fmt.Println(err)
	}
	//读取DS1文件
	dd, _ := m.image.ReadFile(tools.ObjectPath + "/mapSucai/townE1.ds1")
	d, _ := ds1.Unmarshal(dd)

	//加载素材信息提取
	// for i := 0; i < len(d.Files); i++ {
	// 	fmt.Println(strings.ReplaceAll(d.Files[i], "tg1", "dt1"))
	// }

	//floor
	w, h := d.Floors[0].Size()
	img = make([][]*ebiten.Image, h)
	for i := 0; i < h; i++ {
		img[i] = make([]*ebiten.Image, w)
		for j := 0; j < w; j++ {
			ds1Tile := d.Floors[0].Tile(j, i)
			if !ds1Tile.Hidden() && ds1Tile.Prop1 != 0 {
				ds := GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), 0, ss.Tiles)
				if ds != nil {
					img[i][j] = getTitleImage(ds[ds1Tile.RandomIndex], ww)
				}
			}
		}
	}

	//wall
	img2 = make([][]imgWall, h)
	for i := 0; i < h; i++ {
		img2[i] = make([]imgWall, w)
		for j := 0; j < w; j++ {
			ds1Tile := d.Walls[0].Tile(j, i)
			if !ds1Tile.Hidden() && ds1Tile.Prop1 != 0 {
				ds := GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), ds1Tile.Type, ss2.Tiles)
				if ds != nil {
					if ds1Tile.Type == d2enum.TileRightPartOfNorthCornerWall {
						dss := GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), d2enum.TileLeftPartOfNorthCornerWall, ss2.Tiles)
						if dss != nil && dss[ds1Tile.RandomIndex].Height < ds[ds1Tile.RandomIndex].Height {
							m, h := getWallTitleImage(dss[ds1Tile.RandomIndex], ds1Tile, ww)
							img2[i][j].img = m
							img2[i][j].h = h
						} else {
							m, h := getWallTitleImage(ds[ds1Tile.RandomIndex], ds1Tile, ww)
							img2[i][j].img = m
							img2[i][j].h = h
						}
					} else {
						m, h := getWallTitleImage(ds[ds1Tile.RandomIndex], ds1Tile, ww)
						img2[i][j].img = m
						img2[i][j].h = h
					}
				}
			}
		}
	}

	//地图显示补图  hardcode
	img3 = make([][]imgWall, h)
	for i := 0; i < h; i++ {
		img3[i] = make([]imgWall, w)
		for j := 0; j < w; j++ {
			if j == 18 && i == 6 {
				img3[i][j].img = getTitleImage(ss4.Tiles[7], ww)
				img3[i][j].h = -170
			}
			if j == 30 && i == 11 {
				img3[i][j].img = getTitleImage(ss1.Tiles[22], ww)
				img3[i][j].h = -80
			}
			if j == 31 && i == 11 {
				img3[i][j].img = getTitleImage(ss1.Tiles[23], ww)
				img3[i][j].h = -110
			}

			if j == 32 && i == 11 {
				img3[i][j].img = getTitleImage(ss1.Tiles[24], ww)
				img3[i][j].h = -110
			}
			if j == 22 && i == 22 {
				img3[i][j].img = getTitleImage(ss1.Tiles[24], ww)
				img3[i][j].h = -110
			}
			if j == 33 && i == 11 {
				img3[i][j].img = getTitleImage(ss1.Tiles[21], ww)
				img3[i][j].h = -20
			}
			if j == 33 && i == 10 {
				img3[i][j].img = getTitleImage(ss1.Tiles[25], ww)
				img3[i][j].h = -80
			}
			if j == 33 && i == 9 {
				img3[i][j].img = getTitleImage(ss1.Tiles[26], ww)
				img3[i][j].h = -80
			}
			if j == 33 && i == 8 {
				img3[i][j].img = getTitleImage(ss1.Tiles[27], ww)
				img3[i][j].h = -40
			}
			if j == 33 && i == 7 {
				img3[i][j].img = getTitleImage(ss1.Tiles[28], ww)
				img3[i][j].h = -40
			}
			if j == 22 && i == 22 {
				img3[i][j].img = getTitleImage(ss1.Tiles[19], ww)
				img3[i][j].h = -80
			}
			if j == 21 && i == 24 {
				img3[i][j].img = getTitleImage(ss1.Tiles[16], ww)
				img3[i][j].h = -80
			}

		}
	}
}

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

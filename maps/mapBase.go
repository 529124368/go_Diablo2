package maps

import (
	"embed"
	"fmt"
	"game/mapCreator/dat"
	"game/mapCreator/ds1"
	"game/mapCreator/dt1"
	"game/status"
	"game/tools"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Scale float64 = 1
	img   [][]*ebiten.Image
	img2  [][]imgWall
	img3  [][]imgWall
)

type imgWall struct {
	Img *ebiten.Image
	h   int
}

type MapBase struct {
	image   *embed.FS
	BgImage *ebiten.Image
	status  *status.StatusManage //状态
}

//Create Map Class
func NewMap(images *embed.FS, s *status.StatusManage) *MapBase {
	maps := &MapBase{
		image:  images,
		status: s,
	}
	return maps
}

//加载地图图片
func (m *MapBase) LoadMap() {
	//加载动态地图
	m.LoadDstMap()
}

//改变地图坐标
func (m *MapBase) ChangeMapTranslate(x, y float64) {
	m.status.MoveOffsetX += x
}

//渲染地图的地砖
func (m *MapBase) RenderFloor(screen *ebiten.Image, offsetX, offsetY float64) {
	//floor
	sumX := 0
	startY := 0
	for i := 0; i < m.status.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < m.status.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			//视野剔除
			if j > m.status.MapTitleX-m.status.MapZoom && j < m.status.MapTitleX+m.status.MapZoom && i > m.status.MapTitleY-m.status.MapZoom && i < m.status.MapTitleY+m.status.MapZoom {
				s := img[i][j]
				if s != nil {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(3280+float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY)
					op.GeoM.Scale(Scale, Scale)
					screen.DrawImage(s, op)
					// debug  info
					//ebitenutil.DebugPrintAt(screen, "·"+strconv.Itoa(j)+","+strconv.Itoa(i), i*(-80)+sumX+int(offsetX)+74, startY+j*40+int(offsetY)+img3[i][j].h+37)
				}
			}
		}
	}

}

//渲染地图的建筑
func (m *MapBase) RenderWall(screen *ebiten.Image, offsetX, offsetY float64) {
	//补图
	sumX := 0
	startY := 0
	for i := 0; i < m.status.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < m.status.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			s := img3[i][j].Img
			//视野剔除
			if j > m.status.MapTitleX-m.status.MapZoom && j < m.status.MapTitleX+m.status.MapZoom && i > m.status.MapTitleY-m.status.MapZoom && i < m.status.MapTitleY+m.status.MapZoom {
				if s != nil {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(3280+float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY+float64(img3[i][j].h))
					op.GeoM.Scale(Scale, Scale)
					screen.DrawImage(s, op)
				}
			}
		}
	}
	//wall
	sumX = 0
	startY = 0
	for i := 0; i < m.status.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < m.status.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			//视野剔除
			if j > m.status.MapTitleX-m.status.MapZoom && j < m.status.MapTitleX+m.status.MapZoom && i > m.status.MapTitleY-m.status.MapZoom && i < m.status.MapTitleY+m.status.MapZoom {
				s := img2[i][j].Img
				if s != nil {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(3280+float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY+float64(img2[i][j].h))
					op.GeoM.Scale(Scale, Scale)
					screen.DrawImage(s, op)
				}
			}
		}
	}
}

//动态地图解析和加载
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

	w, h := d.Floors[0].Size()
	//保存地图大小
	m.status.ReadMapSizeWidth = w
	m.status.ReadMapSizeHeight = h
	//floor
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
							img2[i][j].Img = m
							img2[i][j].h = h
						} else {
							m, h := getWallTitleImage(ds[ds1Tile.RandomIndex], ds1Tile, ww)
							img2[i][j].Img = m
							img2[i][j].h = h
						}
					} else {
						m, h := getWallTitleImage(ds[ds1Tile.RandomIndex], ds1Tile, ww)
						img2[i][j].Img = m
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
				img3[i][j].Img = getTitleImage(ss4.Tiles[7], ww)
				img3[i][j].h = -170
			}
			if j == 30 && i == 11 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[22], ww)
				img3[i][j].h = -80
			}
			if j == 31 && i == 11 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[23], ww)
				img3[i][j].h = -110
			}

			if j == 32 && i == 11 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[24], ww)
				img3[i][j].h = -110
			}
			if j == 22 && i == 22 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[24], ww)
				img3[i][j].h = -110
			}
			if j == 33 && i == 11 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[21], ww)
				img3[i][j].h = -20
			}
			if j == 33 && i == 10 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[25], ww)
				img3[i][j].h = -80
			}
			if j == 33 && i == 9 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[26], ww)
				img3[i][j].h = -80
			}
			if j == 33 && i == 8 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[27], ww)
				img3[i][j].h = -40
			}
			if j == 33 && i == 7 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[28], ww)
				img3[i][j].h = -40
			}
			if j == 22 && i == 22 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[19], ww)
				img3[i][j].h = -80
			}
			if j == 21 && i == 24 {
				img3[i][j].Img = getTitleImage(ss1.Tiles[16], ww)
				img3[i][j].h = -80
			}

		}
	}
}

//获取墙体区域
func (m *MapBase) GetBlock1Aera() [][]imgWall {
	return img2
}

//获取墙体区域
func (m *MapBase) GetBlock2Aera() [][]imgWall {
	return img3
}

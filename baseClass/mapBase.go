package baseClass

import (
	"embed"
	"game/fonts"
	"game/interfaces"
	"game/mapCreator/ds1"
	"game/mapCreator/dt1"
	"game/maps"
	"game/status"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Scale float64 = 1
)

type ImgWall struct {
	Img *ebiten.Image
	H   int
}

type MapBase struct {
	Image        *embed.FS
	BgImage      *ebiten.Image
	Img_Floor    [][]*ebiten.Image
	Img_Wall     [][]ImgWall
	Img_Wall_Add [][]ImgWall
	DS1          *ds1.DS1
	DT1LIST      []dt1.Tile
	PA           interfaces.Palette
	Fonts        *fonts.FontBase
}

//加载地图图片
func (m *MapBase) LoadMap() {
	//定时任务
	go func() {
		for {
			<-time.After(time.Second * 20)
			m.ClearMap()
		}
	}()
}

//改变地图坐标
func (m *MapBase) ChangeMapTranslate(x, y float64) {
	status.Config.CamerOffsetX += x
	status.Config.CamerOffsetY += y
}

//渲染地图的地砖
func (m *MapBase) RenderFloor(screen *ebiten.Image, offsetX, offsetY float64) {
	//floor
	sumX := 0
	startY := 0
	for i := 0; i < status.Config.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < status.Config.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			//视野剔除
			if j > status.Config.MapTitleX-status.Config.MapZoom && j < status.Config.MapTitleX+status.Config.MapZoom && i > status.Config.MapTitleY-status.Config.MapZoom && i < status.Config.MapTitleY+status.Config.MapZoom {
				s := m.Img_Floor[i][j]
				if s == nil {
					s = m.GetFloor(i, j)
					m.Img_Floor[i][j] = s
				}
				if s != nil {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(3280+float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY)
					op.GeoM.Scale(Scale, Scale)
					screen.DrawImage(s, op)
					// debug  info
					//ebitenutil.DebugPrintAt(screen, "·"+strconv.Itoa(j)+","+strconv.Itoa(i), i*(-80)+sumX+int(offsetX)+74, startY+j*40+int(offsetY)+m.Img3[i][j].h+37)
				}
			}
		}
	}
}

//渲染地图的建筑
func (m *MapBase) RenderWall(screen *ebiten.Image, offsetX, offsetY float64) {
	//wall
	sumX := 0
	startY := 0
	for i := 0; i < status.Config.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < status.Config.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			//视野剔除
			if j > status.Config.MapTitleX-status.Config.MapZoom && j < status.Config.MapTitleX+status.Config.MapZoom && i > status.Config.MapTitleY-status.Config.MapZoom && i < status.Config.MapTitleY+status.Config.MapZoom {
				s := m.Img_Wall[i][j].Img
				if s == nil {
					s, h := m.GetWall(i, j)
					m.Img_Wall[i][j].Img = s
					m.Img_Wall[i][j].H = h
				}
				if s != nil {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(3280+float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY+float64(m.Img_Wall[i][j].H))
					op.GeoM.Scale(Scale, Scale)
					screen.DrawImage(s, op)
				}
			}
		}
	}
}

//获取墙体区域
func (m *MapBase) GetBlock1Aera(x, y int) bool {
	return m.Img_Wall[y][x].Img == nil
}

//动态生成墙体素材
func (m *MapBase) GetWall(i, j int) (*ebiten.Image, int) {
	ds1Tile := m.DS1.Walls[0].Tile(j, i)
	if !ds1Tile.Hidden() && ds1Tile.Prop1 != 0 {
		ds := maps.GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), ds1Tile.Type, m.DT1LIST)
		if ds != nil {
			if ds1Tile.Type == d2enum.TileRightPartOfNorthCornerWall {
				dss := maps.GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), d2enum.TileLeftPartOfNorthCornerWall, m.DT1LIST)
				if dss != nil && dss[ds1Tile.RandomIndex].Height < ds[ds1Tile.RandomIndex].Height {
					m, h := maps.GetWallTitleImage(dss[ds1Tile.RandomIndex], ds1Tile, m.PA)
					return m, h
				} else {
					m, h := maps.GetWallTitleImage(ds[ds1Tile.RandomIndex], ds1Tile, m.PA)
					return m, h
				}
			} else {
				m, h := maps.GetWallTitleImage(ds[ds1Tile.RandomIndex], ds1Tile, m.PA)
				return m, h
			}
		}
	}
	return nil, 0
}

//动态生成地面素材
func (m *MapBase) GetFloor(i, j int) *ebiten.Image {
	ds1Tile := m.DS1.Floors[0].Tile(j, i)
	if !ds1Tile.Hidden() && ds1Tile.Prop1 != 0 {
		ds := maps.GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), 0, m.DT1LIST)
		if ds != nil {
			return maps.GetTitleImage(ds[ds1Tile.RandomIndex], m.PA)
		}
	}
	return nil
}

//清理不需要的地图数据
func (m *MapBase) ClearMap() {
	for i := 0; i < status.Config.ReadMapSizeHeight; i++ {
		for j := 0; j < status.Config.ReadMapSizeWidth; j++ {
			if j > status.Config.MapTitleX-status.Config.MapZoom-3 && j < status.Config.MapTitleX+status.Config.MapZoom+3 &&
				i > status.Config.MapTitleY-status.Config.MapZoom-3 && i < status.Config.MapTitleY+status.Config.MapZoom+3 {
				continue
			} else {
				m.Img_Wall[i][j].Img = nil
				m.Img_Wall[i][j].H = 0
				m.Img_Floor[i][j] = nil
			}
		}
	}
}

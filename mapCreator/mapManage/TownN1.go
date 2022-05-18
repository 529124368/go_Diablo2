package mapManage

import (
	"embed"
	"errors"
	"fmt"
	"game/baseClass"
	"game/mapCreator/dat"
	"game/mapCreator/ds1"
	"game/mapCreator/dt1"
	"game/maps"
	"game/status"
	"game/tools"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/hajimehoshi/ebiten/v2"
)

type TownN1 struct {
	baseClass.MapBase                 //继承
	anmiList          []*ebiten.Image //火把动画图集
	wayList           []*ebiten.Image //瞬间移动动画图集
	huodui            []*ebiten.Image //火堆动画图集
	NPC               []*ebiten.Image //NPC图集
	dropAnm           []*ebiten.Image //掉落动画图集
	dropItemsList     []dropItem      //掉落物品一栏
	op                []*ebiten.DrawImageOptions
	//xyPos         [17]postion
	image *embed.FS //静态资源获取
}

func NewN1(images *embed.FS) *TownN1 {
	a := &TownN1{
		anmiList:      make([]*ebiten.Image, 0),
		wayList:       make([]*ebiten.Image, 0),
		huodui:        make([]*ebiten.Image, 0),
		NPC:           make([]*ebiten.Image, 0),
		dropAnm:       make([]*ebiten.Image, 0),
		dropItemsList: make([]dropItem, 0),
		op:            make([]*ebiten.DrawImageOptions, 0),
		image:         images,
	}
	a.Image = images
	return a
}

//加载动画资源
func (t *TownN1) LoadAnm() {
	t.LoadXyList()
	for i := 0; i < 17; i++ {
		o, _ := t.image.ReadFile("resource/itemsdrop/c_" + strconv.Itoa(i) + ".png")
		t.dropAnm = append(t.dropAnm, tools.GetEbitenImage(o))
	}

}

//加载动画坐标
func (t *TownN1) LoadXyList() {

}

//渲染掉落物品
func (t *TownN1) RenderDropItems(screen *ebiten.Image, offsetX, offsetY float64, playX, playY float64, mx, my int) {
	//掉落物品
	for i := 0; i < len(t.dropItemsList); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(t.dropItemsList[i].pos.x+offsetX, t.dropItemsList[i].pos.y+offsetY)
		screen.DrawImage(t.dropAnm[16], op)
	}
}
func (t *TownN1) SortLayer(mapX, mapY int) {
	if mapY >= 26 || mapY == 16 && mapX >= 47 || mapX == 46 || mapX == 47 {
		status.Config.DisplaySort = true
	} else {
		status.Config.DisplaySort = false
	}
}

//渲染地图上物体
func (t *TownN1) Render(screen *ebiten.Image, frameIndexFor20, frameIndexFor12 int, offsetX, offsetY float64) {

}

//暗黑新手村专用 根据逻辑坐标 求具体坐标
func (t *TownN1) GetCellXY(x, y int) (float64, float64, error) {
	if x < 0 || x > 57 || y < 0 || y > 41 {
		str := "坐标范围不正确"
		return 0, 0, errors.New(str)
	}
	startY := 0
	sumX := 0
	for i := 0; i < status.Config.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < status.Config.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			if j == x && y == i {
				return float64(3280 + i*(-80) + sumX), float64(startY + j*40), nil
			}

		}
	}
	str := "没有找到匹配的位置"
	return 0, 0, errors.New(str)
}

//播放丢物品动画
func (t *TownN1) PlayDropItemAnm(screen *ebiten.Image, x, y float64, name string) {
	go func() {
		countForMap := 0
		countsForMap := 0
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear
		op.GeoM.Translate(float64(tools.LAYOUTX/2-220), float64(tools.LAYOUTY/2)-50)
		for {
			countForMap++
			screen.DrawImage(t.dropAnm[countsForMap], op)
			//切换图
			if countForMap > 4000 {
				countsForMap++
				countForMap = 0
				if countsForMap >= 16 {
					t.InsertOnLoadItesm(name, x, y)
					return
				}
			}
		}
	}()
}

//记录掉落在地面的物品信息
func (t *TownN1) InsertOnLoadItesm(name string, x, y float64) {
	var i dropItem
	i.name = name
	i.pos.x = x - 40
	i.pos.y = y - 130 + 40
	t.dropItemsList = append(t.dropItemsList, i)
}

//加载地图图片
func (t *TownN1) LoadMap() {
	//加载动态地图
	//加载调色板
	r, _ := t.image.ReadFile(tools.ObjectPath + "/mapSucai/pal.dat")
	ww, _ := dat.Load(r)
	//加载地块dt1素材
	re, _ := t.image.ReadFile(tools.ObjectPath + "/mapSucai/floor.dt1")
	ss, _ := dt1.LoadDT1(re)
	re, _ = t.image.ReadFile(tools.ObjectPath + "/mapSucai/objects.dt1")
	ss1, _ := dt1.LoadDT1(re)

	re, _ = t.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/objects.dt1")
	ss2, _ := dt1.LoadDT1(re)
	re, _ = t.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/treegroups.dt1")
	ss3, _ := dt1.LoadDT1(re)

	re, _ = t.image.ReadFile(tools.ObjectPath + "/mapSucai/fence.dt1")
	ss4, _ := dt1.LoadDT1(re)

	re, _ = t.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/bridge.dt1")
	ss5, _ := dt1.LoadDT1(re)

	re, _ = t.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/stonewall.dt1")
	ss6, _ := dt1.LoadDT1(re)

	re, _ = t.image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/river.dt1")
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
	dd, _ := t.image.ReadFile(tools.ObjectPath + "/mapSucai/townN1.ds1")
	d, _ := ds1.Unmarshal(dd)

	w, h := d.Floors[0].Size()
	//保存地图大小
	status.Config.ReadMapSizeWidth = w
	status.Config.ReadMapSizeHeight = h
	//floor
	t.Img_Floor = make([][]*ebiten.Image, h)
	for i := 0; i < h; i++ {
		t.Img_Floor[i] = make([]*ebiten.Image, w)
		for j := 0; j < w; j++ {
			ds1Tile := d.Floors[0].Tile(j, i)
			if !ds1Tile.Hidden() && ds1Tile.Prop1 != 0 {
				ds := maps.GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), 0, ss.Tiles)
				if ds != nil {
					t.Img_Floor[i][j] = maps.GetTitleImage(ds[ds1Tile.RandomIndex], ww)
				}
			}
		}
	}

	//wall
	t.Img_Wall = make([][]baseClass.ImgWall, h)
	for i := 0; i < h; i++ {
		t.Img_Wall[i] = make([]baseClass.ImgWall, w)
		for j := 0; j < w; j++ {
			ds1Tile := d.Walls[0].Tile(j, i)
			if !ds1Tile.Hidden() && ds1Tile.Prop1 != 0 {
				ds := maps.GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), ds1Tile.Type, ss2.Tiles)
				if ds != nil {
					if ds1Tile.Type == d2enum.TileRightPartOfNorthCornerWall {
						dss := maps.GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), d2enum.TileLeftPartOfNorthCornerWall, ss2.Tiles)
						if dss != nil && dss[ds1Tile.RandomIndex].Height < ds[ds1Tile.RandomIndex].Height {
							m, h := maps.GetWallTitleImage(dss[ds1Tile.RandomIndex], ds1Tile, ww)
							t.Img_Wall[i][j].Img = m
							t.Img_Wall[i][j].H = h
						} else {
							m, h := maps.GetWallTitleImage(ds[ds1Tile.RandomIndex], ds1Tile, ww)
							t.Img_Wall[i][j].Img = m
							t.Img_Wall[i][j].H = h
						}
					} else {
						m, h := maps.GetWallTitleImage(ds[ds1Tile.RandomIndex], ds1Tile, ww)
						t.Img_Wall[i][j].Img = m
						t.Img_Wall[i][j].H = h
					}
				}
			}
		}
	}

}

//渲染地图的建筑
func (t *TownN1) RenderWall(screen *ebiten.Image, offsetX, offsetY float64) {

	t.MapBase.RenderWall(screen, offsetX, offsetY)
}

//后期加入的可行走区域
func (t *TownN1) GetBlock1AeraUpdate(x, y int) bool {
	return false
}

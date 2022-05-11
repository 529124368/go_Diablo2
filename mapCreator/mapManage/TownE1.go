package mapManage

import (
	"embed"
	"errors"
	"fmt"
	"game/baseClass"
	"game/mapCreator/dat"
	"game/mapCreator/ds1"
	"game/mapCreator/dt1"
	"game/mapCreator/mapManage/npc"
	"game/maps"
	"game/status"
	"game/storage"
	"game/tools"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/hajimehoshi/ebiten/v2"
)

var Scale float64 = 1

type postion struct {
	x float64
	y float64
}

//掉落物品
type dropItem struct {
	name string
	pos  postion
}
type TownE1 struct {
	baseClass.MapBase                 //继承
	anmiList          []*ebiten.Image //火把动画图集
	wayList           []*ebiten.Image //瞬间移动动画图集
	huodui            []*ebiten.Image //火堆动画图集
	dropAnm           []*ebiten.Image //掉落动画图集
	dropItemsList     []dropItem      //掉落物品一览
	op                []*ebiten.DrawImageOptions
	xyPos             [17]postion
	bag               *storage.Bag
	NPCAI             [4]*npc.NpcAI //AI NPC
}

func NewE1(images *embed.FS, sta *status.StatusManage, b *storage.Bag) *TownE1 {
	a := &TownE1{
		anmiList:      make([]*ebiten.Image, 0),
		wayList:       make([]*ebiten.Image, 0),
		huodui:        make([]*ebiten.Image, 0),
		dropAnm:       make([]*ebiten.Image, 0),
		dropItemsList: make([]dropItem, 0),
		op:            make([]*ebiten.DrawImageOptions, 0),
		bag:           b,
	}
	a.Image = images
	a.Status = sta
	return a
}

//加载动画资源
func (t *TownE1) LoadAnm() {
	t.LoadXyList()
	for i := 0; i < 20; i++ {
		o, _ := t.Image.ReadFile(tools.ObjectPath + "/fire/frame_" + strconv.Itoa(i) + ".png")
		t.anmiList = append(t.anmiList, tools.GetEbitenImage(o))
		o, _ = t.Image.ReadFile(tools.ObjectPath + "/huodui/frame_" + strconv.Itoa(i) + ".png")
		t.huodui = append(t.huodui, tools.GetEbitenImage(o))
	}
	for i := 0; i < 15; i++ {
		o, _ := t.Image.ReadFile(tools.ObjectPath + "/waypoint/frame_" + strconv.Itoa(i) + ".png")
		t.wayList = append(t.wayList, tools.GetEbitenImage(o))

	}
	for i := 0; i < 17; i++ {
		o, _ := t.Image.ReadFile("resource/itemsdrop/box_" + strconv.Itoa(i) + ".png")
		t.dropAnm = append(t.dropAnm, tools.GetEbitenImage(o))
	}
	//设置NPC DC
	t.NPCAI[0] = npc.NewPlayerAI(4580, 2041, 0, 0, t.Status, t.Image)
	t.NPCAI[0].LoadImages("DC", "/NPC/", 1)
	aiPath := make([]npc.AIEndPoint, 3)
	aiPath = append(aiPath, npc.AIEndPoint{X: 4674, Y: 1947, Dir: 2})
	aiPath = append(aiPath, npc.AIEndPoint{X: 4780, Y: 2053, Dir: 3})
	aiPath = append(aiPath, npc.AIEndPoint{X: 4580, Y: 2053, Dir: 5})
	t.NPCAI[0].SetAIPath(aiPath, 100)

	//设置NPC PS
	t.NPCAI[1] = npc.NewPlayerAI(6003, 2048, 0, 4, t.Status, t.Image)
	t.NPCAI[1].LoadImages("PS", "/NPC/", 1)
	aiPath = make([]npc.AIEndPoint, 2)
	aiPath = append(aiPath, npc.AIEndPoint{X: 6200, Y: 2048, Dir: 7})
	aiPath = append(aiPath, npc.AIEndPoint{X: 6003, Y: 2048, Dir: 5})
	t.NPCAI[1].SetAIPath(aiPath, 100)

	//设置NPC PS
	t.NPCAI[2] = npc.NewPlayerAI(4823, 1691, 0, 4, t.Status, t.Image)
	t.NPCAI[2].LoadImages("RC", "/NPC/", 1)
	aiPath = make([]npc.AIEndPoint, 2)
	aiPath = append(aiPath, npc.AIEndPoint{X: 4823, Y: 1855, Dir: 4})
	aiPath = append(aiPath, npc.AIEndPoint{X: 4918, Y: 1759, Dir: 2})
	aiPath = append(aiPath, npc.AIEndPoint{X: 4823, Y: 1855, Dir: 0})
	aiPath = append(aiPath, npc.AIEndPoint{X: 4823, Y: 1691, Dir: 6})
	t.NPCAI[2].SetAIPath(aiPath, 100)

	//设置NPC GH
	t.NPCAI[3] = npc.NewPlayerAI(3230, 1905, 0, 3, t.Status, t.Image)
	t.NPCAI[3].LoadImages("GH", "/NPC/", 1)
	aiPath = make([]npc.AIEndPoint, 2)
	aiPath = append(aiPath, npc.AIEndPoint{X: 3482, Y: 1905, Dir: 7})
	aiPath = append(aiPath, npc.AIEndPoint{X: 3354, Y: 2005, Dir: 0})
	aiPath = append(aiPath, npc.AIEndPoint{X: 3482, Y: 1905, Dir: 2})
	aiPath = append(aiPath, npc.AIEndPoint{X: 3230, Y: 1905, Dir: 5})
	t.NPCAI[3].SetAIPath(aiPath, 100)
}

//加载动画坐标
func (t *TownE1) LoadXyList() {
	syList := [17]string{
		"21,23", "23,21", "22,16", "21,11", "24,9", "27,13", "33,11", "34,7", "38,8", "41,9", "44,7", "41,12", "46,17", "46,14", "43,25", "35,10", "31,16",
	}
	for i, k := range syList {
		re := strings.Split(k, ",")
		x, _ := strconv.Atoi(re[0])
		y, _ := strconv.Atoi(re[1])
		xx, yy, _ := t.GetCellXY(x, y)
		t.xyPos[i].x = xx
		t.xyPos[i].y = yy
		opr := &ebiten.DrawImageOptions{}
		opr.GeoM.Translate(xx, yy)
		opr.GeoM.Scale(Scale, Scale)
		t.op = append(t.op, opr)
	}
}

//渲染掉落物品
func (t *TownE1) RenderDropItems(screen *ebiten.Image, offsetX, offsetY float64, playX, playY float64) {
	//掉落物品
	for i := 0; i < len(t.dropItemsList); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(t.dropItemsList[i].pos.x+offsetX, t.dropItemsList[i].pos.y+offsetY)
		screen.DrawImage(t.dropAnm[16], op)
		if tools.Distance(int64(playX), int64(playY), int64(t.dropItemsList[i].pos.x), int64(t.dropItemsList[i].pos.y+130)) <= 30 {
			if t.bag.InsertBag(t.dropItemsList[i].name) {
				if i != len(t.dropItemsList)-1 {
					t.dropItemsList = append(t.dropItemsList[:i], t.dropItemsList[i+1:]...)
				} else {
					t.dropItemsList = t.dropItemsList[:i]
				}
			}
		}
	}
}

//切换渲染顺序
func (t *TownE1) SortLayer(mapX, mapY int) {
	if mapY >= 26 || mapY == 16 && mapX >= 47 || mapX == 46 || mapX == 47 || mapY == 8 && mapX == 30 || mapY == 7 && mapX == 30 || mapY == 6 && mapX == 30 {
		t.Status.DisplaySort = true
	} else {
		t.Status.DisplaySort = false
	}
}

//渲染地图上物体
func (t *TownE1) Render(screen *ebiten.Image, frameIndexFor20, frameIndexFor12 int, offsetX, offsetY float64) {

	//普通火台
	for i, k := range t.op {
		if i <= 14 {
			//shadow
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Rotate(-0.5)
			op.GeoM.Scale(1, 0.5)
			op.ColorM.Scale(0, 0, 0, 1)
			op.GeoM.Translate(float64(t.xyPos[i].x)+offsetX-50, float64(t.xyPos[i].y)+offsetY+50)
			screen.DrawImage(t.anmiList[frameIndexFor20], op)
			//
			k.GeoM.Reset()
			k.GeoM.Translate(float64(t.xyPos[i].x)+offsetX, float64(t.xyPos[i].y)+offsetY)
			k.Filter = ebiten.FilterLinear
			k.GeoM.Scale(Scale, Scale)
			screen.DrawImage(t.anmiList[frameIndexFor20], k)
		}
	}
	//传送点火焰
	t.op[15].GeoM.Reset()
	t.op[15].GeoM.Translate(float64(t.xyPos[15].x)+100+offsetX, float64(t.xyPos[15].y)+offsetY)
	t.op[15].Filter = ebiten.FilterLinear
	t.op[15].CompositeMode = ebiten.CompositeModeLighter
	t.op[15].GeoM.Scale(Scale, Scale)
	if frameIndexFor20 > 14 {
		screen.DrawImage(t.wayList[frameIndexFor20-15], t.op[15])
	} else {
		screen.DrawImage(t.wayList[frameIndexFor20], t.op[15])
	}
	//火堆
	t.op[16].GeoM.Reset()
	t.op[16].GeoM.Translate(float64(t.xyPos[16].x)+20+offsetX, float64(t.xyPos[16].y)-30+offsetY)
	t.op[16].Filter = ebiten.FilterLinear
	t.op[16].GeoM.Scale(Scale, Scale)
	screen.DrawImage(t.huodui[frameIndexFor20], t.op[16])

	//AI NPC
	for _, v := range t.NPCAI {
		v.Render(screen)
	}
}

//暗黑新手村专用 根据逻辑坐标 求具体坐标
func (t *TownE1) GetCellXY(x, y int) (float64, float64, error) {
	if x < 0 || x > 57 || y < 0 || y > 41 {
		str := "坐标范围不正确"
		return 0, 0, errors.New(str)
	}
	startY := 0
	sumX := 0
	for i := 0; i < t.Status.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < t.Status.ReadMapSizeWidth; j++ {
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
func (t *TownE1) PlayDropItemAnm(screen *ebiten.Image, x, y float64, name string, countsFor17 int) bool {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	op.GeoM.Translate(float64(tools.LAYOUTX/2-220), float64(tools.LAYOUTY/2)-50)
	screen.DrawImage(t.dropAnm[countsFor17], op)
	//判断是否掉落动画最后一帧
	if countsFor17 == 16 {
		t.InsertOnLoadItesm(name, x, y)
		return true
	} else {
		return false
	}
}

//记录掉落在地面的物品信息
func (t *TownE1) InsertOnLoadItesm(name string, x, y float64) {
	var i dropItem
	i.name = name
	i.pos.x = x - 40
	i.pos.y = y - 130 + 40
	t.dropItemsList = append(t.dropItemsList, i)
}

//加载地图图片
func (t *TownE1) LoadMap() {
	//加载动态地图
	//加载调色板
	r, _ := t.Image.ReadFile(tools.ObjectPath + "/mapSucai/pal.dat")
	ww, _ := dat.Load(r)
	//加载地块dt1素材
	re, _ := t.Image.ReadFile(tools.ObjectPath + "/mapSucai/floor.dt1")
	ss, _ := dt1.LoadDT1(re)
	re, _ = t.Image.ReadFile(tools.ObjectPath + "/mapSucai/objects.dt1")
	ss1, _ := dt1.LoadDT1(re)

	re, _ = t.Image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/objects.dt1")
	ss2, _ := dt1.LoadDT1(re)
	re, _ = t.Image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/treegroups.dt1")
	ss3, _ := dt1.LoadDT1(re)

	re, _ = t.Image.ReadFile(tools.ObjectPath + "/mapSucai/fence.dt1")
	ss4, _ := dt1.LoadDT1(re)

	re, _ = t.Image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/bridge.dt1")
	ss5, _ := dt1.LoadDT1(re)

	re, _ = t.Image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/stonewall.dt1")
	ss6, _ := dt1.LoadDT1(re)

	re, _ = t.Image.ReadFile(tools.ObjectPath + "/mapSucai/outdoor/river.dt1")
	ss7, err := dt1.LoadDT1(re)

	ss2.Tiles = append(ss2.Tiles, ss1.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss3.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss4.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss5.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss6.Tiles...)
	ss2.Tiles = append(ss2.Tiles, ss7.Tiles...)
	ss.Tiles = append(ss.Tiles, ss2.Tiles...)

	if err != nil {
		fmt.Println(err)
	}
	//读取DS1文件
	dd, _ := t.Image.ReadFile(tools.ObjectPath + "/mapSucai/townE1.ds1")
	d, _ := ds1.Unmarshal(dd)

	w, h := d.Floors[0].Size()
	//保存地图大小
	t.Status.ReadMapSizeWidth = w
	t.Status.ReadMapSizeHeight = h

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
				ds := maps.GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), ds1Tile.Type, ss.Tiles)
				if ds != nil {
					if ds1Tile.Type == d2enum.TileRightPartOfNorthCornerWall {
						dss := maps.GetTiles(int(ds1Tile.Style), int(ds1Tile.Sequence), d2enum.TileLeftPartOfNorthCornerWall, ss.Tiles)
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

	//地图显示补图  hardcode
	t.Img_Wall_Add = make([][]baseClass.ImgWall, h)
	for i := 0; i < h; i++ {
		t.Img_Wall_Add[i] = make([]baseClass.ImgWall, w)
		for j := 0; j < w; j++ {
			if j == 18 && i == 6 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss4.Tiles[7], ww)
				t.Img_Wall_Add[i][j].H = -170
			}
			if j == 30 && i == 11 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[22], ww)
				t.Img_Wall_Add[i][j].H = -80
			}
			if j == 31 && i == 11 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[23], ww)
				t.Img_Wall_Add[i][j].H = -110
			}

			if j == 32 && i == 11 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[24], ww)
				t.Img_Wall_Add[i][j].H = -110
			}
			if j == 22 && i == 22 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[24], ww)
				t.Img_Wall_Add[i][j].H = -110
			}
			if j == 33 && i == 11 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[21], ww)
				t.Img_Wall_Add[i][j].H = -20
			}
			if j == 33 && i == 10 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[25], ww)
				t.Img_Wall_Add[i][j].H = -80
			}
			if j == 33 && i == 9 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[26], ww)
				t.Img_Wall_Add[i][j].H = -80
			}
			if j == 33 && i == 8 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[27], ww)
				t.Img_Wall_Add[i][j].H = -40
			}
			if j == 33 && i == 7 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[28], ww)
				t.Img_Wall_Add[i][j].H = -40
			}
			if j == 22 && i == 22 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[19], ww)
				t.Img_Wall_Add[i][j].H = -80
			}
			if j == 21 && i == 24 {
				t.Img_Wall_Add[i][j].Img = maps.GetTitleImage(ss1.Tiles[16], ww)
				t.Img_Wall_Add[i][j].H = -80
			}

		}
	}
	//GC
	ss2 = nil
	ss = nil
	d = nil
	ww = nil
}

//渲染地图的建筑
func (t *TownE1) RenderWall(screen *ebiten.Image, offsetX, offsetY float64) {
	//补图
	sumX := 0
	startY := 0
	for i := 0; i < t.Status.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < t.Status.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			//视野剔除
			if j > t.Status.MapTitleX-t.Status.MapZoom && j < t.Status.MapTitleX+t.Status.MapZoom && i > t.Status.MapTitleY-t.Status.MapZoom && i < t.Status.MapTitleY+t.Status.MapZoom {
				s := t.Img_Wall_Add[i][j].Img
				if s != nil {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(3280+float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY+float64(t.Img_Wall_Add[i][j].H))
					op.GeoM.Scale(Scale, Scale)
					screen.DrawImage(s, op)
				}
			}
		}
	}
	//父类
	t.MapBase.RenderWall(screen, offsetX, offsetY)
}

//后期加入自定义的可行走区域
func (t *TownE1) GetBlock1AeraUpdate(x, y int) bool {
	if y == 8 && x == 22 || y == 8 && x == 21 || y == 20 && x == 22 || y == 17 && x == 19 || y == 20 && x == 21 || y == 26 && x == 22 || y == 18 && x == 47 || y == 18 && x == 48 || y == 15 && x == 44 || y == 16 && x == 47 || y == 24 && x == 35 || y == 22 && x == 32 || y == 26 && x == 21 || y == 20 && x == 30 || y == 20 && x == 24 {
		return true
	} else {
		return false
	}
}

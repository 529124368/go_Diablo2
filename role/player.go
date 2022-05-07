package role

import (
	"embed"
	"game/controller"
	"game/engine/ws"
	"game/interfaces"
	"game/status"
	"game/tools"
	"math"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	OFFSETX int = -30
	OFFSETY int = -50
)

type Player struct {
	PlayerName               string                  //玩家名字
	X                        float64                 //玩家世界坐标X
	Y                        float64                 //玩家世界坐标Y
	State                    uint8                   //玩家状态
	Direction                uint8                   //玩家当前方向
	OldDirection             uint8                   //玩家旧的方向
	MouseX                   int                     //鼠标X坐标
	MouseY                   int                     //鼠标Y坐标
	SkillName                string                  //技能名称
	image                    *embed.FS               //静态资源获取
	map_c                    interfaces.MapInterface //地图
	status                   *status.StatusManage    //状态
	plist_sheet              *texturepacker.SpriteSheet
	plist_sheet_2            *texturepacker.SpriteSheet
	plist_png                *ebiten.Image
	plist_png_2              *ebiten.Image
	opS                      *ebiten.DrawImageOptions
	op                       *ebiten.DrawImageOptions
	newPath                  []uint8
	turnLoop                 uint8
	WsCon                    *ws.WsNetManage //net
	newpositonX, newpositonY float64
	newDir                   uint8
}

//创建玩家
func NewPlayer(x, y float64, state, dir uint8, mx, my int, images *embed.FS, m interfaces.MapInterface, s *status.StatusManage, con *ws.WsNetManage) *Player {

	play := &Player{
		PlayerName:   "",
		X:            x, //地图坐标X
		Y:            y, //地图坐标Y
		State:        state,
		Direction:    dir,
		OldDirection: dir,
		MouseX:       mx,
		MouseY:       my,
		SkillName:    "", //技能名字
		image:        images,
		map_c:        m,
		status:       s,
		turnLoop:     0,
		opS:          &ebiten.DrawImageOptions{},
		op:           &ebiten.DrawImageOptions{},
		WsCon:        con,
	}
	return play
}

//Load Images
func (p *Player) LoadImages(name string, num uint8) {
	if num == 2 {
		//加载玩家素材第一部分
		plist, _ := p.image.ReadFile("resource/man/warrior/" + name + ".png")
		plist_json, _ := p.image.ReadFile("resource/man/warrior/" + name + ".json")
		pli, pln := tools.GetImageFromPlistPaletted(plist, plist_json)
		p.plist_sheet = pli
		p.plist_png = ebiten.NewImageFromImage(pln)
		//加载玩家素材第二部分
		plist, _ = p.image.ReadFile("resource/man/warrior/" + name + "_act.png")
		plist_json, _ = p.image.ReadFile("resource/man/warrior/" + name + "_act.json")
		pli, pln = tools.GetImageFromPlistPaletted(plist, plist_json)
		p.plist_sheet_2 = pli
		p.plist_png_2 = ebiten.NewImageFromImage(pln)
	} else {
		//加载玩家素材第一部分
		plist, _ := p.image.ReadFile("resource/man/warrior/" + name + ".png")
		plist_json, _ := p.image.ReadFile("resource/man/warrior/" + name + ".json")
		pli, pln := tools.GetImageFromPlistPaletted(plist, plist_json)
		p.plist_sheet = pli
		p.plist_png = ebiten.NewImageFromImage(pln)
	}

	p.SetPlayerState(0, 0)
}

//设置玩家状态
func (p *Player) SetPlayerState(s, d uint8) {
	p.State = s
	if p.status.Flg {
		p.Direction = d
	}
}

//更新玩家旧的方向
func (p *Player) UpdateOldPlayerDir(d uint8) {
	p.OldDirection = d
}

//获取图片
func (p *Player) GetAnimator(flg, name string, block uint8) (*ebiten.Image, int, int) {
	if flg == "man" {
		//判断加载素材的第几部分
		if block == 1 {
			return p.plist_png.SubImage(p.plist_sheet.Sprites[name+".png"].Frame).(*ebiten.Image), p.plist_sheet.Sprites[name+".png"].SpriteSourceSize.Min.X, p.plist_sheet.Sprites[name+".png"].SpriteSourceSize.Min.Y
		} else {
			return p.plist_png_2.SubImage(p.plist_sheet_2.Sprites[name+".png"].Frame).(*ebiten.Image), p.plist_sheet_2.Sprites[name+".png"].SpriteSourceSize.Min.X, p.plist_sheet_2.Sprites[name+".png"].SpriteSourceSize.Min.Y
		}
	} else {
		return nil, 0, 0
	}
}

//暗黑破坏神 16方位 移动 鼠标控制
func (p *Player) GetMouseController(dir uint8) {
	if p.status.Flg {
		speed := 0.0
		//判断是否走路
		if p.status.IsWalk && (p.Direction != dir || p.State != tools.Walk) {
			speed = tools.SPEED
			p.SetPlayerState(tools.Walk, dir)
		}
		if !p.status.IsWalk && (p.Direction != dir || p.State != tools.RUN) {
			speed = tools.SPEED_RUN
			p.SetPlayerState(tools.RUN, dir)
		}
		//移动判断
		moveX, moveY := tools.CalculateSpeed(dir, speed)
		if p.CanWalk(moveX, moveY, dir) {
			p.status.CamerOffsetX += -moveX
			p.status.CamerOffsetY += -moveY
			p.Y += moveY
			p.X += moveX
		}
		p.status.Flg = false
	}
}

//暗黑破坏神 16方位 移动 鼠标控制
func (p *Player) GetMouseControllerCopy(dir uint8) {
	if p.status.Flg {
		speed := 0.0
		//判断是否走路
		if p.status.IsWalk && (p.Direction != dir || p.State != tools.Walk) {
			speed = tools.SPEED
			p.SetPlayerState(tools.Walk, dir)
		}
		if !p.status.IsWalk && (p.Direction != dir || p.State != tools.RUN) {
			speed = tools.SPEED_RUN
			p.SetPlayerState(tools.RUN, dir)
		}
		//移动判断
		moveX, moveY := tools.CalculateSpeed(dir, speed)
		if p.CanWalk(moveX, moveY, dir) {
			p.Y += moveY
			p.X += moveX
		}
		p.status.Flg = false
	}
}

//判断是否可以行走
func (p *Player) CanWalk(xS, yS float64, dir uint8) bool {
	x, y := tools.GetFloorPositionAt(p.X+xS-110, p.Y+yS+70)
	if x >= p.status.ReadMapSizeWidth || y >= p.status.ReadMapSizeHeight || x < 0 || y < 0 {
		p.SetPlayerState(tools.IDLE, dir)
		return false
	}
	//根据地图判断是否可以走
	if p.map_c.GetBlock1Aera(x, y) || p.map_c.GetBlock1AeraUpdate(x, y) {
		return true
	} else {
		p.SetPlayerState(tools.IDLE, dir)
		return false
	}
}

//本机玩家移动
func (p *Player) PlayerMove() {
	if p.newpositonX != 0 && p.newpositonY != 0 && (math.Abs(p.X-p.newpositonX) > 1 && math.Abs(p.Y-p.newpositonY) > 1) {
		p.status.Flg = true
		//判断人物方位
		if p.OldDirection != p.Direction && !controller.MouseRightPress() {
			if !p.status.CalculateEnd {
				p.newPath = tools.CalculateDirPath(p.OldDirection, p.Direction)
				p.status.CalculateEnd = true
			}
			if len(p.newPath) >= 3 {
				if p.turnLoop >= uint8(len(p.newPath)) {
					p.turnLoop = uint8(len(p.newPath) - 1)
					p.newDir = p.newPath[p.turnLoop]
					p.UpdateOldPlayerDir(p.Direction)
				} else {
					p.newDir = p.newPath[p.turnLoop]
				}
				p.turnLoop++
				p.SetPlayerState(tools.IDLE, p.newDir)
			} else {
				//直接切换方向
				p.status.CalculateEnd = false
				p.turnLoop = 0
				p.UpdateOldPlayerDir(p.Direction)
				p.GetMouseController(p.Direction)
			}
		} else {
			p.status.CalculateEnd = false
			p.turnLoop = 0
			p.GetMouseController(p.newDir)
		}
	} else {
		p.newpositonX = 0
		p.newpositonY = 0
	}
}

//非控制玩家移动
func (p *Player) PlayerMoveCopy() {
	if p.newpositonX != 0 && p.newpositonY != 0 && (math.Abs(p.X-p.newpositonX) > 1 && math.Abs(p.Y-p.newpositonY) > 1) {
		p.status.Flg = true
		//判断人物方位
		if p.OldDirection != p.Direction && !controller.MouseRightPress() {
			if !p.status.CalculateEnd {
				p.newPath = tools.CalculateDirPath(p.OldDirection, p.Direction)
				p.status.CalculateEnd = true
			}
			if len(p.newPath) >= 3 {
				if p.turnLoop >= uint8(len(p.newPath)) {
					p.turnLoop = uint8(len(p.newPath) - 1)
					p.newDir = p.newPath[p.turnLoop]
					p.UpdateOldPlayerDir(p.Direction)
				} else {
					p.newDir = p.newPath[p.turnLoop]
				}
				p.turnLoop++
				p.SetPlayerState(tools.IDLE, p.newDir)
			} else {
				//直接切换方向
				p.status.CalculateEnd = false
				p.turnLoop = 0
				p.UpdateOldPlayerDir(p.Direction)
				p.GetMouseControllerCopy(p.Direction)
			}
		} else {
			p.status.CalculateEnd = false
			p.turnLoop = 0
			p.GetMouseControllerCopy(p.newDir)
		}
	} else {
		p.newpositonX = 0
		p.newpositonY = 0
	}
}

//玩家到新位置的预算
func (p *Player) PlayerNextMovePositon(mouseX, mouseY int, dir uint8) {
	p.newDir = dir
	p.newpositonX = p.X + float64(mouseX) - 395
	p.newpositonY = p.Y + float64(mouseY) - 240
}

//GC
func (p *Player) GC() {
	p.plist_sheet = nil
	p.plist_sheet_2 = nil
	p.plist_png = nil
	p.plist_png_2 = nil
}

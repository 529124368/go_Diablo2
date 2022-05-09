package role

import (
	"embed"
	"game/controller"
	"game/engine/ws"
	"game/interfaces"
	"game/status"
	"game/tools"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	OFFSETX int = -30
	OFFSETY int = -50
)

type Player struct {
	PlayerBase                                                   //继承
	PlayerName                           string                  //玩家名字
	MouseX, MouseY                       int                     //鼠标X坐标 鼠标Y坐标
	SkillName                            string                  //技能名称
	map_c                                interfaces.MapInterface //地图
	opS, op                              *ebiten.DrawImageOptions
	newPath                              []uint8
	turnLoop                             uint8
	WsCon                                *ws.WsNetManage //net
	newpositonX, newpositonY             float64
	newDir                               uint8
	FrameSpeed, FrameNums, Counts, count int
	imgOffset                            [4]tools.OffsetXY //动作图片偏移
}

//创建玩家
func NewPlayer(x, y float64, state, dir uint8, mx, my int, images *embed.FS, m interfaces.MapInterface, s *status.StatusManage, con *ws.WsNetManage) *Player {

	play := &Player{
		PlayerName: "",
		MouseX:     mx,
		MouseY:     my,
		SkillName:  "", //技能名字
		map_c:      m,
		turnLoop:   0,
		opS:        &ebiten.DrawImageOptions{},
		op:         &ebiten.DrawImageOptions{},
		WsCon:      con,
	}
	play.X = x //地图坐标X
	play.Y = y //地图坐标Y
	play.State = state
	play.Direction = dir
	play.OldDirection = dir
	play.image = images
	play.status = s
	return play
}

//加載玩家素材
func (p *Player) LoadImages(name string, num uint8) {
	p.PlayerBase.LoadImages(name, num)
	p.imgOffset = tools.GetOffetByAction(name)
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
		} else {
			p.newpositonX = 0
			p.newpositonY = 0
			p.status.CalculateEnd = false
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
		p.status.IsRun = true
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
				p.UpdateOldPlayerDir(p.newDir)
				p.GetMouseController(p.newDir)
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
	//玩家停止
	if p.State == tools.IDLE && p.status.IsRun {
		p.status.IsRun = false
		//网络
		if p.status.IsNetPlay {
			p.WsCon.SendMessage("@@MoveEnd|" + p.PlayerName + "|" + strconv.FormatFloat(p.X, 'f', 0, 64) + "|" + strconv.FormatFloat(p.Y, 'f', 0, 64))
		}
	}
}

//玩家到新位置的预算
func (p *Player) PlayerNextMovePositon(mouseX, mouseY int, dir uint8) {
	p.newDir = dir
	p.newpositonX = p.X + float64(mouseX) - 395
	p.newpositonY = p.Y + float64(mouseY) - 240
	if p.status.IsNetPlay {
		//网络
		p.WsCon.SendMessage("@@Move|" + p.PlayerName + "|" + strconv.FormatFloat(p.newpositonX, 'f', 0, 64) + "|" + strconv.FormatFloat(p.newpositonY, 'f', 0, 64) + "|" + strconv.Itoa(int(p.newDir)))
	}
}

//渲染本机角色
func (p *Player) Render(screen *ebiten.Image) {
	p.count++
	//Change player Frame
	if p.count > p.FrameSpeed {
		p.Counts++
		p.count = 0
		if p.Counts >= p.FrameNums {
			p.Counts = 0
		}
	}
	var name string
	block := 1
	//nameSkill := ""
	switch p.State {
	case tools.ATTACK:
		name = strconv.Itoa(int(p.Direction)) + "_attack_" + strconv.Itoa(p.Counts)
	case tools.SkILL:
		block = 2
		if p.Counts >= 14 {
			p.Counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_skill_" + strconv.Itoa(p.Counts)
	case tools.IDLE:
		name = strconv.Itoa(int(p.Direction)) + "_stand_" + strconv.Itoa(p.Counts)
	case tools.Walk:
		if p.Counts >= 8 {
			p.Counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_run_" + strconv.Itoa(p.Counts)
	case tools.RUN:
		block = 2
		if p.Counts >= 8 {
			p.Counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_run2_" + strconv.Itoa(p.Counts)
	}

	imagess, x, y := p.GetAnimator("man", name, uint8(block))
	//Idel -> Walk Offset
	if p.State == tools.Walk {
		x += p.imgOffset[0].X
		y += p.imgOffset[0].Y
	}
	//Idel -> RUN Offset
	if p.State == tools.RUN {
		x += p.imgOffset[1].X
		y += p.imgOffset[1].Y
	}

	//Idel -> Walk -> Attack Offset
	if p.State == tools.ATTACK {
		x += p.imgOffset[2].X
		y += p.imgOffset[2].Y
	}

	//Idel -> SKILL-> Offset
	if p.State == tools.SkILL {
		x += p.imgOffset[3].X
		y += p.imgOffset[3].Y
	}

	//Draw Shadow
	p.opS.GeoM.Reset()
	p.opS.Filter = ebiten.FilterLinear
	p.opS.GeoM.Rotate(-0.5)
	p.opS.GeoM.Scale(1, 0.5)
	p.opS.ColorM.Scale(0, 0, 0, 1)
	p.opS.GeoM.Translate(float64(tools.LAYOUTX/2+x+p.status.ShadowOffsetX+p.status.UIOFFSETX), float64(tools.LAYOUTY/2+y))
	screen.DrawImage(imagess, p.opS)
	//Draw Player
	p.op.GeoM.Reset()
	p.op.GeoM.Translate(float64(tools.LAYOUTX/2+OFFSETX+x+p.status.UIOFFSETX), float64(tools.LAYOUTY/2+OFFSETY+y))
	p.op.Filter = ebiten.FilterLinear
	screen.DrawImage(imagess, p.op)

	//Draw Skill
	// if g.player.State == ATTACK {
	// 	imagey, x, y := g.player.GetAnimator("skill", nameSkill)
	// 	//skill option
	// 	opSkill = &ebiten.DrawImageOptions{}
	// 	opSkill.GeoM.Translate(float64(SCREENWIDTH/2+x), float64(SCREENHEIGHT/2+y))
	// 	opSkill.CompositeMode = ebiten.CompositeModeLighter
	// 	opSkill.GeoM.Scale(1.5, 1.5)
	// 	opSkill.Filter = ebiten.FilterLinear
	// 	screen.DrawImage(imagey, opSkill)
	// }
}

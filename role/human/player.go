package human

import (
	"embed"
	"game/baseClass"
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
	baseClass.PlayerBase                         //继承
	PlayerName           string                  //玩家名字
	MouseX, MouseY       int                     //鼠标X坐标 鼠标Y坐标
	SkillName            string                  //技能名称
	map_c                interfaces.MapInterface //地图
	newPath              []uint8
	turnLoOp             uint8
	WsCon                *ws.WsNetManage   //net
	imgOffset            [4]tools.OffsetXY //动作图片偏移
}

//创建玩家
func NewPlayer(x, y float64, state, dir uint8, mx, my int, images *embed.FS, m interfaces.MapInterface, s *status.StatusManage, con *ws.WsNetManage) *Player {
	play := &Player{
		PlayerName: "",
		MouseX:     mx,
		MouseY:     my,
		SkillName:  "", //技能名字
		map_c:      m,
		turnLoOp:   0,
		WsCon:      con,
	}
	play.X = x //地图坐标X
	play.Y = y //地图坐标Y
	play.State = state
	play.Direction = dir
	play.OldDirection = dir
	play.Asset = images
	play.Status = s
	play.OpS = &ebiten.DrawImageOptions{}
	play.Op = &ebiten.DrawImageOptions{}
	return play
}

//加載玩家素材
func (p *Player) LoadImages(name, path string, num uint8) {
	p.PlayerBase.LoadImages(name, path, num)
	p.imgOffset = tools.GetOffetByAction(name)
}

//暗黑破坏神 16方位 移动 鼠标控制
func (p *Player) GetMouseController(dir uint8) {
	if p.FlagCanAction {
		speed := 0.0
		//判断是否走路
		if p.Status.IsWalk && (p.Direction != dir || p.State != tools.Walk) {
			speed = tools.SPEED
			p.SetPlayerState(tools.Walk, dir)
		}
		if !p.Status.IsWalk && (p.Direction != dir || p.State != tools.RUN) {
			speed = tools.SPEED_RUN
			p.SetPlayerState(tools.RUN, dir)
		}
		//移动判断
		moveX, moveY := tools.CalculateSpeed(dir, speed)
		if p.CanWalk(moveX, moveY, dir) {
			p.Status.CamerOffsetX += -moveX
			p.Status.CamerOffsetY += -moveY
			p.Y += moveY
			p.X += moveX
		} else {
			p.NewpositonX = 0
			p.NewpositonY = 0
			p.Status.CalculateEnd = false
		}

	}
}

//判断是否可以行走
func (p *Player) CanWalk(xS, yS float64, dir uint8) bool {
	x, y := tools.GetFloorPositionAt(p.X+xS-110, p.Y+yS+70)
	if x >= p.Status.ReadMapSizeWidth || y >= p.Status.ReadMapSizeHeight || x < 0 || y < 0 {
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
	if p.NewpositonX != 0 && p.NewpositonY != 0 && (math.Abs(p.X-p.NewpositonX) > 1 && math.Abs(p.Y-p.NewpositonY) > 1) {
		p.Status.IsRun = true
		p.FlagCanAction = true
		//判断人物方位
		if p.OldDirection != p.Direction && !controller.MouseRightPress() {
			if !p.Status.CalculateEnd {
				p.newPath = tools.CalculateDirPath(p.OldDirection, p.Direction)
				p.Status.CalculateEnd = true
			}
			if len(p.newPath) >= 3 {
				if p.turnLoOp >= uint8(len(p.newPath)) {
					p.turnLoOp = uint8(len(p.newPath) - 1)
					p.NewDir = p.newPath[p.turnLoOp]
					p.UpdateOldPlayerDir(p.Direction)
				} else {
					p.NewDir = p.newPath[p.turnLoOp]
				}
				p.turnLoOp++
				p.SetPlayerState(tools.IDLE, p.NewDir)
			} else {
				//直接切换方向
				p.Status.CalculateEnd = false
				p.turnLoOp = 0
				p.UpdateOldPlayerDir(p.NewDir)
				p.GetMouseController(p.NewDir)
			}
		} else {
			p.Status.CalculateEnd = false
			p.turnLoOp = 0
			p.GetMouseController(p.NewDir)
		}
	} else {
		p.FlagCanAction = false
		p.NewpositonX = 0
		p.NewpositonY = 0
	}
	//玩家停止
	if p.State == tools.IDLE && p.Status.IsRun {
		p.Status.IsRun = false
		//网络
		if p.Status.IsNetPlay {
			p.WsCon.SendMessage("@@MoveEnd|" + p.PlayerName + "|" + strconv.FormatFloat(p.X, 'f', 0, 64) + "|" + strconv.FormatFloat(p.Y, 'f', 0, 64))
		}
	}
}

//玩家到新位置的预算
func (p *Player) PlayerNextMovePositon(mouseX, mouseY int, dir uint8) {
	p.NewDir = dir
	p.NewpositonX = p.X + float64(mouseX) - 395
	p.NewpositonY = p.Y + float64(mouseY) - 240
	if p.Status.IsNetPlay {
		//网络
		p.WsCon.SendMessage("@@Move|" + p.PlayerName + "|" + strconv.FormatFloat(p.NewpositonX, 'f', 0, 64) + "|" + strconv.FormatFloat(p.NewpositonY, 'f', 0, 64) + "|" + strconv.Itoa(int(p.NewDir)))
	}
}

//渲染角色
func (p *Player) Render(screen *ebiten.Image) {
	p.PlayerMove()
	p.ChangeFrame()
	p.PlayerBase.Render()
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
	p.OpS.GeoM.Reset()
	p.OpS.Filter = ebiten.FilterLinear
	p.OpS.GeoM.Rotate(-0.45)
	p.OpS.GeoM.Scale(1, 0.5)
	p.OpS.ColorM.Scale(0, 0, 0, 1)
	p.OpS.GeoM.Translate(float64(tools.LAYOUTX/2+x+p.Status.ShadowOffsetX+p.Status.UIOFFSETX), float64(tools.LAYOUTY/2+y+p.Status.ShadowOffsetY))
	screen.DrawImage(imagess, p.OpS)
	//Draw Player
	p.Op.GeoM.Reset()
	p.Op.GeoM.Translate(float64(tools.LAYOUTX/2+OFFSETX+x+p.Status.UIOFFSETX), float64(tools.LAYOUTY/2+OFFSETY+y))
	p.Op.Filter = ebiten.FilterLinear
	screen.DrawImage(imagess, p.Op)

	//Draw Skill
	// if g.player.State == ATTACK {
	// 	imagey, x, y := g.player.GetAnimator("skill", nameSkill)
	// 	//skill Option
	// 	OpSkill = &ebiten.DrawImageOptions{}
	// 	OpSkill.GeoM.Translate(float64(SCREENWIDTH/2+x), float64(SCREENHEIGHT/2+y))
	// 	OpSkill.CompositeMode = ebiten.CompositeModeLighter
	// 	OpSkill.GeoM.Scale(1.5, 1.5)
	// 	OpSkill.Filter = ebiten.FilterLinear
	// 	screen.DrawImage(imagey, OpSkill)
	// }
}

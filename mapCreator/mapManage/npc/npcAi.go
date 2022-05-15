package npc

import (
	"embed"
	"game/baseClass"
	"game/status"
	"game/tools"
	"math"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type AIEndPoint struct {
	X, Y float64
	Dir  uint8
}
type NpcAI struct {
	baseClass.PlayerBase                     //继承
	PlayerName                        string //玩家名字
	AIWait, AICount, AIf, AIPathCount int
	AIpath                            []AIEndPoint
}

//创建NPC
func NewPlayerAI(x, y float64, state, dir uint8, images *embed.FS) *NpcAI {
	play := &NpcAI{
		PlayerName: "",
		AIWait:     0,
		AICount:    0,
		AIf:        0,
	}
	play.X = x //地图坐标X
	play.Y = y //地图坐标Y
	play.State = state
	play.Direction = dir
	play.Asset = images
	play.OpS = &ebiten.DrawImageOptions{}
	play.Op = &ebiten.DrawImageOptions{}
	return play
}

//暗黑破坏神 16方位 移动 AI
func (p *NpcAI) GetMouseControllerAI(dir uint8) {
	if p.FlagCanAction {
		speed := 0.0
		//判断是否走路
		speed = 0.5
		p.SetPlayerState(tools.Walk, dir)
		//移动判断
		moveX, moveY := tools.CalculateSpeed(dir, speed)
		p.Y += moveY
		p.X += moveX
	}
}

//停止AI NPC移动
func (p *NpcAI) StopPlayerMoveAI() {
	p.NewpositonX = 0
	p.NewpositonY = 0
}

//控制AI NPC新位置的预算
func (p *NpcAI) UpdatePlayerNextMovePositonAI(NewpositonX, NewpositonY float64, dir uint8) {
	p.AICount = 0
	p.NewDir = dir
	p.NewpositonX = NewpositonX
	p.NewpositonY = NewpositonY
	p.AIf++
}

//渲染NPC
func (p *NpcAI) Render(screen *ebiten.Image) {
	p.PlayerMoveAI()
	p.ChangeFrame()
	p.PlayerBase.Render()
	//写入缓存
	var name strings.Builder
	block := 1
	name.WriteString(strconv.Itoa(int(p.Direction)))
	switch p.State {
	case tools.IDLE:
		name.WriteString("_stand_")
	case tools.Walk:
		if p.Counts >= 8 {
			p.Counts = 0
		}
		name.WriteString("_walk_")
	}
	name.WriteString(strconv.Itoa(p.Counts))
	imagess, x, y := p.GetAnimator("man", name.String(), uint8(block))
	//Draw Shadow
	p.OpS.GeoM.Reset()
	p.OpS.Filter = ebiten.FilterLinear
	p.OpS.GeoM.Rotate(-0.5)
	p.OpS.GeoM.Scale(1, 0.5)
	p.OpS.ColorM.Scale(0, 0, 0, 1)
	p.OpS.GeoM.Translate(float64(int(p.X)+x-32-25)+status.Config.CamerOffsetX, float64(int(p.Y)+y+35-30)+status.Config.CamerOffsetY)
	screen.DrawImage(imagess, p.OpS)
	//Draw Player
	p.Op.GeoM.Reset()
	p.Op.GeoM.Translate(float64(int(p.X)+x-25)+status.Config.CamerOffsetX, float64(int(p.Y)+y-30)+status.Config.CamerOffsetY)
	p.Op.Filter = ebiten.FilterLinear
	screen.DrawImage(imagess, p.Op)
}

//控制AI NPC移动
func (p *NpcAI) PlayerMoveAI() {
	//AI 移动判断
	p.AIMove()
	if p.NewDir == 5 || p.NewDir == 6 || p.NewDir == 7 || p.NewDir == 4 {
		//移动
		if p.NewpositonX != 0 && p.NewpositonY != 0 && (math.Abs(p.X-p.NewpositonX) > 0 || math.Abs(p.Y-p.NewpositonY) > 0) {
			p.FlagCanAction = true
			//直接切换方向
			p.GetMouseControllerAI(p.NewDir)
		} else {
			p.FlagCanAction = false
			p.State = tools.IDLE
			if p.NewpositonX != 0 && p.NewpositonY != 0 {
				p.X = p.NewpositonX
				p.Y = p.NewpositonY
			}
			p.NewpositonX = 0
			p.NewpositonY = 0
		}
	} else {
		//移动
		if p.NewpositonX != 0 && p.NewpositonY != 0 && (math.Abs(p.X-p.NewpositonX) > 0 && math.Abs(p.Y-p.NewpositonY) > 0) {
			p.FlagCanAction = true
			//直接切换方向
			p.GetMouseControllerAI(p.NewDir)
		} else {
			p.FlagCanAction = false
			p.State = tools.IDLE
			if p.NewpositonX != 0 && p.NewpositonY != 0 {
				p.X = p.NewpositonX
				p.Y = p.NewpositonY
			}
			p.NewpositonX = 0
			p.NewpositonY = 0
		}
	}
}

//AI 移动判断
func (p *NpcAI) AIMove() {
	p.AICount++
	p.AIf %= p.AIPathCount
	if !p.FlagCanAction && p.AICount >= p.AIWait {
		p.UpdatePlayerNextMovePositonAI(p.AIpath[p.AIf].X, p.AIpath[p.AIf].Y, p.AIpath[p.AIf].Dir)
	}
}

//设置AI路径
/**
** @params path  AI 移动的路径点
** @params speed AI移动速度
**/
func (p *NpcAI) SetAIPath(path []AIEndPoint, speed int) {
	p.AIPathCount = len(path)
	p.AIpath = path
	p.AIWait = speed
}

//改变帧数
func (p *NpcAI) ChangeFrame() {
	//根据状态改变帧数
	if p.State == tools.IDLE {
		p.FrameNums = 8
		p.FrameSpeed = 5
	} else {
		p.FrameNums = 8
		p.FrameSpeed = 7
	}
}

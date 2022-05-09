package npc

import (
	"embed"
	"game/baseClass"
	"game/status"
	"game/tools"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type NpcAI struct {
	baseClass.PlayerBase        //继承
	PlayerName           string //玩家名字
	//imgOffset            [4]tools.OffsetXY //动作图片偏移
}

//创建NPC
func NewPlayerAI(x, y float64, state, dir uint8, s *status.StatusManage, images *embed.FS) *NpcAI {
	play := &NpcAI{
		PlayerName: "",
	}
	play.X = x //地图坐标X
	play.Y = y //地图坐标Y
	play.State = state
	play.Status = s
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
		speed = tools.SPEED
		p.SetPlayerState(tools.Walk, dir)
		//移动判断
		moveX, moveY := tools.CalculateSpeed(dir, speed)
		p.Y += moveY
		p.X += moveX
		p.FlagCanAction = false
	}
}

//控制AI NPC移动
func (p *NpcAI) PlayerMoveAI() {
	if p.NewpositonX != 0 && p.NewpositonY != 0 && (math.Abs(p.X-p.NewpositonX) > 1 && math.Abs(p.Y-p.NewpositonY) > 1) {
		p.FlagCanAction = true
		//直接切换方向
		p.GetMouseControllerAI(p.NewDir)
	} else {
		p.State = tools.IDLE
		p.NewpositonX = 0
		p.NewpositonY = 0
	}
}

//停止AI NPC移动
func (p *NpcAI) StopPlayerMoveAI() {
	p.NewpositonX = 0
	p.NewpositonY = 0
}

//控制AI NPC新位置的预算
func (p *NpcAI) UpdatePlayerNextMovePositonAI(NewpositonX, NewpositonY float64, dir uint8) {
	p.NewDir = dir
	p.NewpositonX = NewpositonX
	p.NewpositonY = NewpositonY
}

//渲染NPC
func (p *NpcAI) Render(screen *ebiten.Image) {
	p.ChangeFrame()
	p.PlayerBase.Render()
	var name string
	block := 1
	switch p.State {
	case tools.IDLE:
		name = strconv.Itoa(int(p.Direction)) + "_stand_" + strconv.Itoa(p.Counts)
	case tools.Walk:
		if p.Counts >= 8 {
			p.Counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_walk_" + strconv.Itoa(p.Counts)
	}

	imagess, x, y := p.GetAnimator("man", name, uint8(block))

	//Draw Shadow
	p.OpS.GeoM.Reset()
	p.OpS.Filter = ebiten.FilterLinear
	p.OpS.GeoM.Rotate(-0.5)
	p.OpS.GeoM.Scale(1, 0.5)
	p.OpS.ColorM.Scale(0, 0, 0, 1)
	p.OpS.GeoM.Translate(float64(int(p.X)+x-32-25)+p.Status.CamerOffsetX, float64(int(p.Y)+y+35-30)+p.Status.CamerOffsetY)
	screen.DrawImage(imagess, p.OpS)
	//Draw Player
	p.Op.GeoM.Reset()
	p.Op.GeoM.Translate(float64(int(p.X)+x-25)+p.Status.CamerOffsetX, float64(int(p.Y)+y-30)+p.Status.CamerOffsetY)
	p.Op.Filter = ebiten.FilterLinear
	screen.DrawImage(imagess, p.Op)
}

//改变帧数
func (p *NpcAI) ChangeFrame() {
	//根据状态改变帧数
	if p.State == tools.IDLE {
		p.FrameNums = 8
		p.FrameSpeed = 5
	} else {
		p.FrameNums = 8
		p.FrameSpeed = 6
	}
}

package human

import (
	"embed"
	"game/baseClass"
	"game/status"
	"game/tools"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerAI struct {
	baseClass.PlayerBase                   //继承
	PlayerName           string            //玩家名字
	SkillName            string            //技能名称
	imgOffset            [4]tools.OffsetXY //动作图片偏移
}

//创建玩家
func NewPlayerAI(x, y float64, state, dir uint8, s *status.StatusManage, images *embed.FS) *PlayerAI {

	play := &PlayerAI{
		PlayerName: "",
		SkillName:  "", //技能名字
	}
	play.X = x //地图坐标X
	play.Y = y //地图坐标Y
	play.State = state
	play.Status = s
	play.Direction = dir
	play.OldDirection = dir
	play.Asset = images
	play.OpS = &ebiten.DrawImageOptions{}
	play.Op = &ebiten.DrawImageOptions{}
	return play
}

//加載玩家素材
func (p *PlayerAI) LoadImages(name string, num uint8) {
	p.PlayerBase.LoadImages(name, num)
	p.imgOffset = tools.GetOffetByAction(name)
}

//暗黑破坏神 16方位 移动 鼠标控制 AI
func (p *PlayerAI) GetMouseControllerAI(dir uint8) {
	if p.Status.Flg {
		speed := 0.0
		//判断是否走路
		speed = tools.SPEED
		p.SetPlayerState(tools.Walk, dir)
		//移动判断
		moveX, moveY := tools.CalculateSpeed(dir, speed)
		p.Y += moveY
		p.X += moveX
		p.Status.Flg = false
	}
}

//控制AI玩家移动
func (p *PlayerAI) PlayerMoveAI() {
	if p.NewpositonX != 0 && p.NewpositonY != 0 && (math.Abs(p.X-p.NewpositonX) > 1 && math.Abs(p.Y-p.NewpositonY) > 1) {
		p.Status.Flg = true
		//直接切换方向
		p.UpdateOldPlayerDir(p.NewDir)
		p.GetMouseControllerAI(p.NewDir)
	} else {
		p.State = tools.IDLE
		p.NewpositonX = 0
		p.NewpositonY = 0
	}
}

//停止AI玩家移动
func (p *PlayerAI) StopPlayerMoveAI() {
	p.NewpositonX = 0
	p.NewpositonY = 0
}

//控制AI玩家新位置的预算
func (p *PlayerAI) UpdatePlayerNextMovePositonAI(NewpositonX, NewpositonY float64, dir uint8) {
	p.NewDir = dir
	p.NewpositonX = NewpositonX
	p.NewpositonY = NewpositonY
}

//渲染角色
func (p *PlayerAI) Render(screen *ebiten.Image) {
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

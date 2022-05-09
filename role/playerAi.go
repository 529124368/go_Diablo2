package role

import (
	"embed"
	"game/status"
	"game/tools"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerAI struct {
	PlayerBase                                  //继承
	PlayerName                           string //玩家名字
	SkillName                            string //技能名称
	opS, op                              *ebiten.DrawImageOptions
	newpositonX, newpositonY             float64
	newDir                               uint8
	FrameSpeed, FrameNums, Counts, count int
	imgOffset                            [4]tools.OffsetXY //动作图片偏移
}

//创建玩家
func NewPlayerAI(x, y float64, state, dir uint8, s *status.StatusManage, images *embed.FS) *PlayerAI {

	play := &PlayerAI{
		PlayerName: "",
		SkillName:  "", //技能名字
		opS:        &ebiten.DrawImageOptions{},
		op:         &ebiten.DrawImageOptions{},
	}
	play.X = x //地图坐标X
	play.Y = y //地图坐标Y
	play.State = state
	play.status = s
	play.Direction = dir
	play.OldDirection = dir
	play.image = images
	return play
}

//暗黑破坏神 16方位 移动 鼠标控制 AI
func (p *PlayerAI) GetMouseControllerAI(dir uint8) {
	if p.status.Flg {
		speed := 0.0
		//判断是否走路
		speed = tools.SPEED
		p.SetPlayerState(tools.Walk, dir)
		//移动判断
		moveX, moveY := tools.CalculateSpeed(dir, speed)
		p.Y += moveY
		p.X += moveX
		p.status.Flg = false
	}
}

//控制AI玩家移动
func (p *PlayerAI) PlayerMoveAI() {
	if p.newpositonX != 0 && p.newpositonY != 0 && (math.Abs(p.X-p.newpositonX) > 1 && math.Abs(p.Y-p.newpositonY) > 1) {
		p.status.Flg = true
		//直接切换方向
		p.UpdateOldPlayerDir(p.newDir)
		p.GetMouseControllerAI(p.newDir)
	} else {
		p.State = tools.IDLE
		p.newpositonX = 0
		p.newpositonY = 0
	}
}

//停止AI玩家移动
func (p *PlayerAI) StopPlayerMoveAI() {
	p.newpositonX = 0
	p.newpositonY = 0
}

//控制AI玩家新位置的预算
func (p *PlayerAI) UpdatePlayerNextMovePositonAI(newpositonX, newpositonY float64, dir uint8) {
	p.newDir = dir
	p.newpositonX = newpositonX
	p.newpositonY = newpositonY
}

//渲染本机角色
func (p *PlayerAI) Render(screen *ebiten.Image) {
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

	//渲染Ai角色
	//Draw Shadow
	p.opS.GeoM.Reset()
	p.opS.Filter = ebiten.FilterLinear
	p.opS.GeoM.Rotate(-0.5)
	p.opS.GeoM.Scale(1, 0.5)
	p.opS.ColorM.Scale(0, 0, 0, 1)
	p.opS.GeoM.Translate(float64(int(p.X)+x-32-25)+p.status.CamerOffsetX, float64(int(p.Y)+y+35-30)+p.status.CamerOffsetY)
	screen.DrawImage(imagess, p.opS)
	//Draw Player
	p.op.GeoM.Reset()
	p.op.GeoM.Translate(float64(int(p.X)+x-25)+p.status.CamerOffsetX, float64(int(p.Y)+y-30)+p.status.CamerOffsetY)
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

package monster

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
type MonsterAI struct {
	baseClass.PlayerBase                     //继承
	MonsterrName                      string //玩家名字
	AIWait, AICount, AIf, AIPathCount int
	AIpath                            []AIEndPoint
}

//创建Monster
func NewMonsterAI(x, y float64, state, dir uint8, images *embed.FS) *MonsterAI {
	monster := &MonsterAI{
		MonsterrName: "",
		AIWait:       0,
		AICount:      0,
		AIf:          0,
	}
	monster.X = x //地图坐标X
	monster.Y = y //地图坐标Y
	monster.State = state
	monster.Direction = dir
	monster.Asset = images
	monster.OpS = &ebiten.DrawImageOptions{}
	monster.Op = &ebiten.DrawImageOptions{}
	return monster
}

//暗黑破坏神 16方位 移动 AI
func (p *MonsterAI) GetMouseControllerAI(dir uint8, dx, dy, std float64) {
	if p.FlagCanAction {
		speed := 0.0
		//判断是否走路
		speed = tools.SPEED
		p.SetPlayerState(tools.Walk, dir)
		//移动判断
		moveX, moveY := tools.CalculateSpeed(dir, speed, dx, dy, std)
		p.Y += moveY
		p.X += moveX
	}
}

//停止AI Monster移动
func (p *MonsterAI) StopPlayerMoveAI() {
	p.NewpositonX = 0
	p.NewpositonY = 0
}

//控制AI NPC新位置的预算
func (p *MonsterAI) UpdatePlayerNextMovePositonAI(NewpositonX, NewpositonY float64, dir uint8) {
	p.AICount = 0
	p.NewDir = dir
	p.NewpositonX = NewpositonX
	p.NewpositonY = NewpositonY
	p.AIf++
}

//渲染NPC
func (p *MonsterAI) Render(screen *ebiten.Image) {
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
	p.OpS.GeoM.Translate(float64(int(p.X)+x-30-25)+status.Config.CamerOffsetX, float64(int(p.Y)+y+35-30)+status.Config.CamerOffsetY)
	screen.DrawImage(imagess, p.OpS)
	//Draw Monster
	p.Op.GeoM.Reset()
	p.Op.GeoM.Translate(float64(int(p.X)+x-25)+status.Config.CamerOffsetX, float64(int(p.Y)+y-30)+status.Config.CamerOffsetY)
	p.Op.Filter = ebiten.FilterLinear
	screen.DrawImage(imagess, p.Op)
}

//控制AI Monster移动
func (p *MonsterAI) PlayerMoveAI() {
	//AI 移动判断
	p.AIMove()
	dx := math.Abs(p.X - p.NewpositonX)
	dy := math.Abs(p.Y - p.NewpositonY)
	std := math.Sqrt(dx*dx + dy*dy)
	//移动
	if p.NewpositonX != 0 && p.NewpositonY != 0 && std > 1 {
		p.FlagCanAction = true
		//直接切换方向
		p.GetMouseControllerAI(p.NewDir, dx, dy, std)
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

//AI 移动判断
func (p *MonsterAI) AIMove() {
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
func (p *MonsterAI) SetAIPath(path []AIEndPoint, speed int) {
	p.AIPathCount = len(path)
	p.AIpath = path
	p.AIWait = speed
}

//改变帧数
func (p *MonsterAI) ChangeFrame() {
	//根据状态改变帧数
	if p.State == tools.IDLE {
		p.FrameNums = 20
		p.FrameSpeed = 5
	} else {
		p.FrameNums = 10
		p.FrameSpeed = 3
	}
}

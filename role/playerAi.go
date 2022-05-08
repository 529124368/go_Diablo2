package role

import (
	"game/tools"
	"math"
)

//暗黑破坏神 16方位 移动 鼠标控制 AI
func (p *Player) GetMouseControllerAI(dir uint8) {
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
func (p *Player) PlayerMoveAI() {
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
func (p *Player) StopPlayerMoveAI() {
	p.newpositonX = 0
	p.newpositonY = 0
}

//控制AI玩家新位置的预算
func (p *Player) UpdatePlayerNextMovePositonAI(newpositonX, newpositonY float64, dir uint8) {
	p.newDir = dir
	p.newpositonX = newpositonX
	p.newpositonY = newpositonY
}

//GC
func (p *Player) GC() {
	p.plist_sheet = nil
	p.plist_sheet_2 = nil
	p.plist_png = nil
	p.plist_png_2 = nil
}

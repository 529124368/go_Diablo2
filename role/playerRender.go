package role

import (
	"game/tools"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

//渲染本机角色
func (p *Player) Render(screen *ebiten.Image, counts int) {
	var name string
	block := 1
	//nameSkill := ""
	switch p.State {
	case tools.ATTACK:
		name = strconv.Itoa(int(p.Direction)) + "_attack_" + strconv.Itoa(counts)
	case tools.SkILL:
		block = 2
		if counts >= 14 {
			counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_skill_" + strconv.Itoa(counts)
	case tools.IDLE:
		name = strconv.Itoa(int(p.Direction)) + "_stand_" + strconv.Itoa(counts)
	case tools.Walk:
		if counts >= 8 {
			counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_run_" + strconv.Itoa(counts)
	case tools.RUN:
		block = 2
		if counts >= 8 {
			counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_run2_" + strconv.Itoa(counts)
	}

	imagess, x, y := p.GetAnimator("man", name, uint8(block))
	//Idel -> Walk Offset
	if p.State == tools.Walk {
		//ba1
		// x += -4
		// y += -3
		//ba2
		// x += 4
		// y += -18
		//ba3
		x += 3
		y += -7
	}
	//Idel -> RUN Offset
	if p.State == tools.RUN {
		//ba1
		// x += -4
		// y += -3
		//ba2
		// x += 4
		// y += -18
		//ba3
		x += 8
		y += -10
	}

	//Idel -> Walk -> Attack Offset
	if p.State == tools.ATTACK {
		//ba1
		// x += -50
		// y += -30
		//ba2
		// x += -55
		// y += -35
		//ba3
		x += -55
		y += -20
	}

	//Idel -> SKILL-> Offset
	if p.State == tools.SkILL {
		//ba1
		// x += -50
		// y += -30
		//ba2
		// x += -55
		// y += -35
		//ba3
		x += -10
		y += -15
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

//渲染其他客户端角色
func (p *Player) RenderCopy(screen *ebiten.Image, counts int) {
	var name string
	block := 1
	//nameSkill := ""
	switch p.State {
	case tools.ATTACK:
		name = strconv.Itoa(int(p.Direction)) + "_attack_" + strconv.Itoa(counts)
	case tools.SkILL:
		block = 2
		if counts >= 14 {
			counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_skill_" + strconv.Itoa(counts)
	case tools.IDLE:
		name = strconv.Itoa(int(p.Direction)) + "_stand_" + strconv.Itoa(counts)
	case tools.Walk:
		if counts >= 8 {
			counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_run_" + strconv.Itoa(counts)
	case tools.RUN:
		block = 2
		if counts >= 8 {
			counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_run2_" + strconv.Itoa(counts)
	}

	imagess, x, y := p.GetAnimator("man", name, uint8(block))
	//Idel -> Walk Offset
	if p.State == tools.Walk {
		//ba1
		// x += -4
		// y += -3
		//ba2
		// x += 4
		// y += -18
		//ba3
		x += 3
		y += -7
	}
	//Idel -> RUN Offset
	if p.State == tools.RUN {
		//ba1
		// x += -4
		// y += -3
		//ba2
		// x += 4
		// y += -18
		//ba3
		x += 8
		y += -10
	}

	//Idel -> Walk -> Attack Offset
	if p.State == tools.ATTACK {
		//ba1
		// x += -50
		// y += -30
		//ba2
		// x += -55
		// y += -35
		//ba3
		x += -55
		y += -20
	}

	//Idel -> SKILL-> Offset
	if p.State == tools.SkILL {
		//ba1
		// x += -50
		// y += -30
		//ba2
		// x += -55
		// y += -35
		//ba3
		x += -10
		y += -15
	}
	//Draw Shadow
	p.opS.GeoM.Reset()
	p.opS.Filter = ebiten.FilterLinear
	p.opS.GeoM.Rotate(-0.5)
	p.opS.GeoM.Scale(1, 0.5)
	p.opS.ColorM.Scale(0, 0, 0, 1)
	p.opS.GeoM.Translate(float64(int(p.X)+x-32)+p.status.CamerOffsetX, float64(int(p.Y)+y+35)+p.status.CamerOffsetY)
	screen.DrawImage(imagess, p.opS)
	//Draw Player
	p.op.GeoM.Reset()
	p.op.GeoM.Translate(float64(int(p.X)+x)+p.status.CamerOffsetX, float64(int(p.Y)+y)+p.status.CamerOffsetY)
	p.op.Filter = ebiten.FilterLinear
	screen.DrawImage(imagess, p.op)
}

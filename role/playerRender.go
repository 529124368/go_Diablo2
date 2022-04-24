package role

import (
	"game/tools"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

//渲染角色
func (p *Player) Render(screen *ebiten.Image, counts int) {
	var name string
	block := 1
	//nameSkill := ""
	switch p.State {
	case tools.ATTACK:
		name = strconv.Itoa(int(p.Direction)) + "_attack_" + strconv.Itoa(counts) + ".png"
	case tools.SkILL:
		block = 2
		if counts >= 14 {
			counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_skill_" + strconv.Itoa(counts) + ".png"
	case tools.IDLE:
		name = strconv.Itoa(int(p.Direction)) + "_stand_" + strconv.Itoa(counts) + ".png"
	case tools.Walk:
		if counts >= 8 {
			counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_run_" + strconv.Itoa(counts) + ".png"
	case tools.RUN:
		block = 2
		if counts >= 8 {
			counts = 0
		}
		name = strconv.Itoa(int(p.Direction)) + "_run2_" + strconv.Itoa(counts) + ".png"
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
	opS.GeoM.Reset()
	opS.GeoM.Translate(float64(tools.LAYOUTX/2+x+p.status.ShadowOffsetX+p.status.UIOFFSETX), float64(tools.LAYOUTY/2+y+p.status.ShadowOffsetY))
	opS.Filter = ebiten.FilterLinear
	opS.GeoM.Rotate(-0.5)
	opS.GeoM.Scale(1, 0.5)
	opS.ColorM.Scale(0, 0, 0, 1)
	screen.DrawImage(imagess, opS)
	//Draw Player
	op.GeoM.Reset()
	op.GeoM.Translate(float64(tools.LAYOUTX/2+OFFSETX+x+p.status.UIOFFSETX), float64(tools.LAYOUTY/2+OFFSETY+y))
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(imagess, op)

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

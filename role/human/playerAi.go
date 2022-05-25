package human

import (
	"embed"
	"game/baseClass"
	"game/shader"
	"game/status"
	"game/tools"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerAI struct {
	baseClass.PlayerBase                   //继承
	PlayerName           string            //玩家名字
	SkillName            string            //技能名称
	imgOffset            [4]tools.OffsetXY //动作图片偏移
	Op                   *ebiten.DrawRectShaderOptions
	IsOutLine            float32
}

//创建玩家
func NewPlayerAI(x, y float64, state, dir uint8, images *embed.FS) *PlayerAI {
	play := &PlayerAI{
		PlayerName: "",
		SkillName:  "", //技能名字
	}
	play.X = x //地图坐标X
	play.Y = y //地图坐标Y
	play.State = state
	play.Direction = dir
	play.Asset = images
	play.OpS = &ebiten.DrawImageOptions{}
	play.Op = &ebiten.DrawRectShaderOptions{}
	return play
}

//加載玩家素材
func (p *PlayerAI) LoadImages(name, path string, num uint8) {
	p.PlayerBase.LoadImages(name, path, num)
	p.imgOffset = tools.GetOffetByAction(name)
}

//控制AI玩家新位置的预算
func (p *PlayerAI) UpdatePlayerPositonAI(NewpositonX, NewpositonY float64, dir uint8, types string) {
	p.Direction = dir
	p.X = NewpositonX
	p.Y = NewpositonY
	//根据状态切换速度
	if types == "r" {
		p.State = tools.RUN
	} else if types == "w" {
		p.State = tools.Walk
	}
}

//控制AI玩家状态
func (p *PlayerAI) UpdatePlayerState(s uint8) {
	if p.State == tools.IDLE && s != p.State {
		p.Counts = 0
	}
	p.State = s
}

//是否被选中
func (p *PlayerAI) In(mousex, mousey, playX, playY int) {
	x, y, err := tools.CalculateWorldToScreen(int(p.X), int(p.Y), playX, playY)
	if err == nil {
		if tools.Distance(int64(x), int64(y), int64(mousex), int64(mousey)) <= 10 {
			if p.IsOutLine == 1 {
				p.IsOutLine = 0
			} else {
				p.IsOutLine = 1
			}

		}
	}
}

//渲染角色
func (p *PlayerAI) Render(screen *ebiten.Image) {
	//改变帧数
	p.ChangeFrame()
	//渲染角色
	p.PlayerBase.Render()
	//写入缓存
	var name strings.Builder
	block := 1
	//nameSkill := ""
	name.WriteString(strconv.Itoa(int(p.Direction)))
	switch p.State {
	case tools.ATTACK:
		name.WriteString("_attack_")
	case tools.SkILL:
		block = 2
		if p.Counts >= 14 {
			p.Counts = 0
		}
		name.WriteString("_skill_")
	case tools.IDLE:
		name.WriteString("_stand_")
	case tools.Walk:
		if p.Counts >= 8 {
			p.Counts = 0
		}
		name.WriteString("_run_")
	case tools.RUN:
		block = 2
		if p.Counts >= 8 {
			p.Counts = 0
		}
		name.WriteString("_run2_")
	}
	name.WriteString(strconv.Itoa(p.Counts))

	imagess, x, y := p.GetAnimator("man", name.String(), uint8(block))
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
	p.OpS.GeoM.Translate(float64(int(p.X)+x-32-25)+status.Config.CamerOffsetX, float64(int(p.Y)+y+35-30)+status.Config.CamerOffsetY)
	screen.DrawImage(imagess, p.OpS)
	//Draw Player
	p.Op.GeoM.Reset()
	p.Op.Uniforms = map[string]interface{}{
		"IsOutLine": float32(p.IsOutLine),
	}
	p.Op.Images[0] = imagess
	p.Op.GeoM.Translate(float64(int(p.X)+x-25)+status.Config.CamerOffsetX, float64(int(p.Y)+y-30)+status.Config.CamerOffsetY)
	w, h := imagess.Size()
	screen.DrawRectShader(w, h, shader.Shader, p.Op)
}

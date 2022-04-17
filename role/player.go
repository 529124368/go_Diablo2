package role

import (
	"embed"
	"game/maps"
	"game/status"
	"game/tools"
	"image"
	"runtime"
	"strconv"
	"strings"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	plist_sheet       *texturepacker.SpriteSheet
	plist_skill_sheet *texturepacker.SpriteSheet
	plist_skill_png   *image.NRGBA
	plist_png         *image.Paletted
	loadedSkill       string
)

type Player struct {
	X            float64
	Y            float64
	State        uint8                //玩家状态
	Direction    uint8                //玩家当前方向
	OldDirection uint8                //玩家旧的方向
	MouseX       int                  //鼠标X坐标
	MouseY       int                  //鼠标Y坐标
	SkillName    string               //技能名称
	image        *embed.FS            //静态资源获取
	map_c        *maps.MapBase        //地图
	status       *status.StatusManage //状态
	hp           float64              //血
	mp           float64              //蓝
}

//Create Player Class
func NewPlayer(x, y float64, state, dir uint8, mx, my int, images *embed.FS, m *maps.MapBase, s *status.StatusManage) *Player {
	play := &Player{
		X:            x,
		Y:            y,
		State:        state,
		Direction:    dir,
		OldDirection: dir,
		MouseX:       mx,
		MouseY:       my,
		SkillName:    "",
		image:        images,
		map_c:        m,
		status:       s,
		hp:           100,
		mp:           100,
	}
	return play
}

//Load Images
func (p *Player) LoadImages() {

	//加载玩家素材
	plist, _ := p.image.ReadFile("resource/man/warrior/ba1.png")
	plist_json, _ := p.image.ReadFile("resource/man/warrior/ba1.json")
	plist_sheet, plist_png = tools.GetImageFromPlistPaletted(plist, plist_json)
	//skill load
	// go func() {
	// 	loadedSkill = "liehuo"
	// 	plist, _ := p.image.ReadFile("resource/man/skill/liehuo.png")
	// 	plist_json, _ := p.image.ReadFile("resource/man/skill/liehuo.json")
	// 	plist_skill_sheet, plist_skill_png = tools.GetImageFromPlist(plist, plist_json)
	// 	runtime.GC()
	// 	wg.Done()
	// }()
	p.SetPlayerState(0, 0)
}

//TODO 加载技能素材
func (p *Player) loadSkillImages(name string) {
	go func() {
		loadedSkill = name
		plist, _ := p.image.ReadFile("resource/man/skill/" + name + ".png")
		plist_json, _ := p.image.ReadFile("resource/man/skill/" + name + ".json")
		plist_skill_sheet, plist_skill_png = tools.GetImageFromPlist(plist, plist_json)
		runtime.GC()
	}()
}

//设置玩家状态
func (p *Player) SetPlayerState(s, d uint8) {
	p.State = s
	if p.status.Flg {
		p.Direction = d
	}

}

//更新玩家旧的方向
func (p *Player) UpdateOldPlayerDir(d uint8) {
	p.OldDirection = d
}

//获取图片
func (p *Player) GetAnimator(flg, name string) (*ebiten.Image, int, int) {
	if flg == "man" {
		return ebiten.NewImageFromImage(plist_png.SubImage(plist_sheet.Sprites[name].Frame)), plist_sheet.Sprites[name].SpriteSourceSize.Min.X, plist_sheet.Sprites[name].SpriteSourceSize.Min.Y
	} else {
		if p.SkillName != loadedSkill {
			p.loadSkillImages(p.SkillName)
		}
		xy := strings.Split(plist_skill_sheet.Meta.Version, "_")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		return ebiten.NewImageFromImage(plist_skill_png.SubImage(plist_skill_sheet.Sprites[name].Frame)), x, y
	}
}

//暗黑破坏神 16方位 移动 鼠标控制
func (p *Player) GetMouseController(dir uint8) {
	if dir == 2 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(-tools.SPEED, tools.SPEED)
		p.status.MoveOffsetX += -tools.SPEED
		p.status.MoveOffsetY += tools.SPEED
		p.Y -= tools.SPEED
		p.X += tools.SPEED
		p.status.Flg = false
	}

	if dir == 3 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)

		}
		//p.map_c.OpBg.GeoM.Translate(-tools.SPEED, -tools.SPEED)
		p.status.MoveOffsetX += -tools.SPEED
		p.status.MoveOffsetY += -tools.SPEED
		p.Y += tools.SPEED
		p.X += tools.SPEED
		p.status.Flg = false
	}
	if dir == 0 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(tools.SPEED, -tools.SPEED)
		p.status.MoveOffsetX += tools.SPEED
		p.status.MoveOffsetY += -tools.SPEED
		p.Y += tools.SPEED
		p.X -= tools.SPEED
		p.status.Flg = false
	}
	if dir == 1 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(tools.SPEED, tools.SPEED)
		p.status.MoveOffsetX += tools.SPEED
		p.status.MoveOffsetY += tools.SPEED
		p.Y -= tools.SPEED
		p.X -= tools.SPEED
		p.status.Flg = false
	}
	if dir == 5 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(tools.SPEED, 0)
		p.status.MoveOffsetX += tools.SPEED
		p.status.MoveOffsetY += 0
		p.X -= tools.SPEED
		p.status.Flg = false
	}

	if dir == 6 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(0, tools.SPEED)
		p.status.MoveOffsetX += 0
		p.status.MoveOffsetY += tools.SPEED
		p.Y -= tools.SPEED
		p.status.Flg = false
	}

	if dir == 7 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(-tools.SPEED, 0)
		p.status.MoveOffsetX += -tools.SPEED
		p.status.MoveOffsetY += 0
		p.X += tools.SPEED
		p.status.Flg = false
	}

	if dir == 4 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(0, -tools.SPEED)
		p.status.MoveOffsetX += 0
		p.status.MoveOffsetY += -tools.SPEED
		p.Y += tools.SPEED
		p.status.Flg = false
	}
	if dir == 12 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(1-tools.SPEED, tools.SPEED)
		p.status.MoveOffsetX += 1 - tools.SPEED
		p.status.MoveOffsetY += tools.SPEED
		p.Y -= tools.SPEED
		p.X += tools.SPEED - 1
		p.status.Flg = false
	}
	if dir == 2 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(-tools.SPEED, tools.SPEED)
		p.status.MoveOffsetX += -tools.SPEED
		p.status.MoveOffsetY += tools.SPEED
		p.Y -= tools.SPEED
		p.X += tools.SPEED
		p.status.Flg = false
	}
	if dir == 13 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(-tools.SPEED, tools.SPEED-1)
		p.status.MoveOffsetX += -tools.SPEED
		p.status.MoveOffsetY += tools.SPEED - 1
		p.Y -= tools.SPEED - 1
		p.X += tools.SPEED
		p.status.Flg = false
	}
	if dir == 10 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(tools.SPEED, tools.SPEED-1)
		p.status.MoveOffsetX += tools.SPEED
		p.status.MoveOffsetY += tools.SPEED - 1
		p.Y -= tools.SPEED - 1
		p.X -= tools.SPEED
		p.status.Flg = false
	}
	if dir == 1 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(tools.SPEED, tools.SPEED)
		p.status.MoveOffsetX += tools.SPEED
		p.status.MoveOffsetY += tools.SPEED
		p.Y -= tools.SPEED
		p.X -= tools.SPEED
		p.status.Flg = false
	}
	if dir == 11 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(tools.SPEED-1, tools.SPEED)
		p.status.MoveOffsetX += tools.SPEED - 1
		p.status.MoveOffsetY += tools.SPEED
		p.Y -= tools.SPEED
		p.X -= tools.SPEED - 1
		p.status.Flg = false
	}
	if dir == 9 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(tools.SPEED, 1-tools.SPEED)
		p.status.MoveOffsetX += tools.SPEED
		p.status.MoveOffsetY += 1 - tools.SPEED
		p.Y += tools.SPEED - 1
		p.X -= tools.SPEED
		p.status.Flg = false
	}
	if dir == 0 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(tools.SPEED, -tools.SPEED)
		p.status.MoveOffsetX += tools.SPEED
		p.status.MoveOffsetY += -tools.SPEED
		p.Y += tools.SPEED
		p.X -= tools.SPEED
		p.status.Flg = false
	}
	if dir == 8 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(tools.SPEED-1, -tools.SPEED)
		p.status.MoveOffsetX += tools.SPEED - 1
		p.status.MoveOffsetY += -tools.SPEED
		p.Y += tools.SPEED
		p.X -= tools.SPEED - 1
		p.status.Flg = false
	}
	//
	if dir == 15 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(1-tools.SPEED, -tools.SPEED)
		p.status.MoveOffsetX += 1 - tools.SPEED
		p.status.MoveOffsetY += -tools.SPEED
		p.Y += tools.SPEED
		p.X += tools.SPEED - 1
		p.status.Flg = false
	}
	if dir == 3 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(-tools.SPEED, -tools.SPEED)
		p.status.MoveOffsetX += -tools.SPEED
		p.status.MoveOffsetY += -tools.SPEED
		p.Y += tools.SPEED
		p.X += tools.SPEED
		p.status.Flg = false
	}
	if dir == 14 && p.status.Flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		//p.map_c.OpBg.GeoM.Translate(-tools.SPEED, 1-tools.SPEED)
		p.status.MoveOffsetX += -tools.SPEED
		p.status.MoveOffsetY += 1 - tools.SPEED
		p.Y += tools.SPEED - 1
		p.X += tools.SPEED
		p.status.Flg = false
	}
}

//TODO
func (p *Player) Attack() {

}

//TODO
func (p *Player) DeadEvent() {

}

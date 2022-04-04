package role

import (
	"embed"
	"game/maps"
	"game/tools"
	"image"
	"runtime"
	"strconv"
	"strings"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

var plist_sheet, plist_skill_sheet *texturepacker.SpriteSheet
var plist_skill_png *image.NRGBA
var plist_png *image.Paletted
var loadedSkill string

type Player struct {
	X         float64
	Y         float64
	State     int
	Direction int
	MouseX    int
	MouseY    int
	SkillName string
	image     *embed.FS
	map_c     *maps.MapBase
}

//Create Player Class
func NewPlayer(x, y float64, state, dir, mx, my int, images *embed.FS, m *maps.MapBase) *Player {
	play := &Player{
		X:         x,
		Y:         y,
		State:     state,
		Direction: dir,
		MouseX:    mx,
		MouseY:    my,
		SkillName: "",
		image:     images,
		map_c:     m,
	}
	return play
}

//Load Images
func (p *Player) LoadImages() {

	//player load
	plist, _ := p.image.ReadFile("resource/man/warrior/ba.png")
	plist_json, _ := p.image.ReadFile("resource/man/warrior/ba.json")
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

//Load Skill Images
func (p *Player) loadSkillImages(name string) {
	go func() {
		loadedSkill = name
		plist, _ := p.image.ReadFile("resource/man/skill/" + name + ".png")
		plist_json, _ := p.image.ReadFile("resource/man/skill/" + name + ".json")
		plist_skill_sheet, plist_skill_png = tools.GetImageFromPlist(plist, plist_json)
		runtime.GC()
	}()
}

//Set Player Status
func (p *Player) SetPlayerState(s, d int) {
	p.State = s
	p.Direction = d
}

//TODO
func (p *Player) Attack() {

}

//TODO
func (p *Player) DeadEvent() {

}

//Get Animator
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

//Mouse Controller For 16 Direction
func (p *Player) GetMouseController(dir int, flg bool) bool {
	if dir == 2 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(-tools.SPEED, tools.SPEED)
		p.Y -= tools.SPEED
		p.X += tools.SPEED
		flg = false
	}

	if dir == 3 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)

		}
		p.map_c.OpBg.GeoM.Translate(-tools.SPEED, -tools.SPEED)
		p.Y += tools.SPEED
		p.X += tools.SPEED
		flg = false
	}
	if dir == 0 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(tools.SPEED, -tools.SPEED)
		p.Y += tools.SPEED
		p.X -= tools.SPEED
		flg = false
	}
	if dir == 1 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(tools.SPEED, tools.SPEED)
		p.Y -= tools.SPEED
		p.X -= tools.SPEED
		flg = false
	}
	if dir == 5 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(tools.SPEED, 0)
		p.X -= tools.SPEED
		flg = false
	}

	if dir == 6 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(0, tools.SPEED)
		p.Y -= tools.SPEED
		flg = false
	}

	if dir == 7 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(-tools.SPEED, 0)
		p.X += tools.SPEED
		flg = false
	}

	if dir == 4 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(0, -tools.SPEED)
		p.Y += tools.SPEED
		flg = false
	}
	if dir == 12 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(1-tools.SPEED, tools.SPEED)
		p.Y -= tools.SPEED
		p.X += tools.SPEED - 1
		flg = false
	}
	if dir == 2 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(-tools.SPEED, tools.SPEED)
		p.Y -= tools.SPEED
		p.X += tools.SPEED
		flg = false
	}
	if dir == 13 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(-tools.SPEED, tools.SPEED-1)
		p.Y -= tools.SPEED - 1
		p.X += tools.SPEED
		flg = false
	}
	if dir == 10 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(tools.SPEED, tools.SPEED-1)
		p.Y -= tools.SPEED - 1
		p.X -= tools.SPEED
		flg = false
	}
	if dir == 1 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(tools.SPEED, tools.SPEED)
		p.Y -= tools.SPEED
		p.X -= tools.SPEED
		flg = false
	}
	if dir == 11 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(tools.SPEED-1, tools.SPEED)
		p.Y -= tools.SPEED
		p.X -= tools.SPEED - 1
		flg = false
	}
	if dir == 9 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(tools.SPEED, 1-tools.SPEED)
		p.Y += tools.SPEED - 1
		p.X -= tools.SPEED
		flg = false
	}
	if dir == 0 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(tools.SPEED, -tools.SPEED)
		p.Y += tools.SPEED
		p.X -= tools.SPEED
		flg = false
	}
	if dir == 8 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(tools.SPEED-1, -tools.SPEED)
		p.Y += tools.SPEED
		p.X -= tools.SPEED - 1
		flg = false
	}
	//
	if dir == 15 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(1-tools.SPEED, -tools.SPEED)
		p.Y += tools.SPEED
		p.X += tools.SPEED - 1
		flg = false
	}
	if dir == 3 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(-tools.SPEED, -tools.SPEED)
		p.Y += tools.SPEED
		p.X += tools.SPEED
		flg = false
	}
	if dir == 14 && flg {
		if p.Direction != dir || p.State != tools.RUN {
			p.SetPlayerState(tools.RUN, dir)
		}
		p.map_c.OpBg.GeoM.Translate(-tools.SPEED, 1-tools.SPEED)
		p.Y += tools.SPEED - 1
		p.X += tools.SPEED
		flg = false
	}
	return flg
}

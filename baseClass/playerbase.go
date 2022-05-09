package baseClass

import (
	"embed"
	"game/status"
	"game/tools"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

//基础类
type PlayerBase struct {
	X                                    float64              //玩家世界坐标X
	Y                                    float64              //玩家世界坐标Y
	State                                uint8                //玩家状态
	Status                               *status.StatusManage //状态
	Direction                            uint8                //玩家当前方向
	OldDirection                         uint8                //玩家旧的方向
	Asset                                *embed.FS            //静态资源获取
	Plist_sheet, Plist_sheet_2           *texturepacker.SpriteSheet
	Plist_png, Plist_png_2               *ebiten.Image
	FrameSpeed, FrameNums, Counts, Count int
	NewpositonX, NewpositonY             float64
	NewDir                               uint8
	OpS, Op                              *ebiten.DrawImageOptions
}

//加載素材
func (p *PlayerBase) LoadImages(name string, num uint8) {
	//加载玩家素材第一部分
	plist, _ := p.Asset.ReadFile("resource/man/warrior/" + name + ".png")
	plist_json, _ := p.Asset.ReadFile("resource/man/warrior/" + name + ".json")
	pli, pln := tools.GetImageFromPlistPaletted(plist, plist_json)
	p.Plist_sheet = pli
	p.Plist_png = ebiten.NewImageFromImage(pln)
	if num == 2 {
		//加载玩家素材第二部分
		plist, _ = p.Asset.ReadFile("resource/man/warrior/" + name + "_act.png")
		plist_json, _ = p.Asset.ReadFile("resource/man/warrior/" + name + "_act.json")
		pli, pln = tools.GetImageFromPlistPaletted(plist, plist_json)
		p.Plist_sheet_2 = pli
		p.Plist_png_2 = ebiten.NewImageFromImage(pln)
	}
	p.SetPlayerState(0, 0)
}

//设置玩家状态
func (p *PlayerBase) SetPlayerState(s, d uint8) {
	p.State = s
	if p.Status.Flg {
		p.Direction = d
	}
}

//更新旧的方向
func (p *PlayerBase) UpdateOldPlayerDir(d uint8) {
	p.OldDirection = d
}

//获取图片
func (p *PlayerBase) GetAnimator(flg, name string, block uint8) (*ebiten.Image, int, int) {
	if flg == "man" {
		//判断加载素材的第几部分
		if block == 1 {
			return p.Plist_png.SubImage(p.Plist_sheet.Sprites[name+".png"].Frame).(*ebiten.Image), p.Plist_sheet.Sprites[name+".png"].SpriteSourceSize.Min.X, p.Plist_sheet.Sprites[name+".png"].SpriteSourceSize.Min.Y
		} else {
			return p.Plist_png_2.SubImage(p.Plist_sheet_2.Sprites[name+".png"].Frame).(*ebiten.Image), p.Plist_sheet_2.Sprites[name+".png"].SpriteSourceSize.Min.X, p.Plist_sheet_2.Sprites[name+".png"].SpriteSourceSize.Min.Y
		}
	} else {
		return nil, 0, 0
	}
}

//渲染角色
func (p *PlayerBase) Render() {
	p.Count++
	//Change player Frame
	if p.Count > p.FrameSpeed {
		p.Counts++
		p.Count = 0
		if p.Counts >= p.FrameNums {
			p.Counts = 0
		}
	}
}

//GC
func (p *PlayerBase) GC() {
	p.Plist_sheet = nil
	p.Plist_sheet_2 = nil
	p.Plist_png = nil
	p.Plist_png_2 = nil
}

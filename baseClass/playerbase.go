package baseClass

import (
	"embed"
	"game/tools"
	"strings"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

// 基础类
type PlayerBase struct {
	FlagCanAction                        bool      //是否可以移动标志
	X                                    float64   //玩家世界坐标X
	Y                                    float64   //玩家世界坐标Y
	State                                uint8     //玩家状态
	Direction                            uint8     //玩家当前方向
	OldDirection                         uint8     //玩家旧的方向
	Asset                                *embed.FS //静态资源获取
	Plist_sheet, Plist_sheet_2           *texturepacker.SpriteSheet
	Plist_png, Plist_png_2               *ebiten.Image
	FrameSpeed, FrameNums, Counts, Count int
	NewpositonX, NewpositonY             float64
	NewDir                               uint8
	OpS, Op                              *ebiten.DrawImageOptions
}

// 加載素材
func (p *PlayerBase) LoadImages(name, path string, num uint8) {
	//写入缓存
	var st strings.Builder
	st.WriteString("resource")
	st.WriteString(path)
	st.WriteString(name)
	st.WriteString(".png")
	//加载玩家素材第一部分
	plist, _ := p.Asset.ReadFile(st.String())
	st.Reset()
	st.WriteString("resource")
	st.WriteString(path)
	st.WriteString(name)
	st.WriteString(".json")
	plist_json, _ := p.Asset.ReadFile(st.String())
	pli, pln := tools.GetImageFromPlistPaletted(plist, plist_json)
	p.Plist_sheet = pli
	p.Plist_png = ebiten.NewImageFromImage(pln)
	if num == 2 {
		//加载玩家素材第二部分
		st.Reset()
		st.WriteString("resource")
		st.WriteString(path)
		st.WriteString(name)
		st.WriteString("_act.png")
		plist, _ = p.Asset.ReadFile(st.String())
		st.Reset()
		st.WriteString("resource")
		st.WriteString(path)
		st.WriteString(name)
		st.WriteString("_act.json")
		plist_json, _ = p.Asset.ReadFile(st.String())
		pli, pln = tools.GetImageFromPlistPaletted(plist, plist_json)
		p.Plist_sheet_2 = pli
		p.Plist_png_2 = ebiten.NewImageFromImage(pln)
	}
	p.SetPlayerState(0, 0)
}

// 复用素材
func (p *PlayerBase) RepeatedImages(s *texturepacker.SpriteSheet, m *ebiten.Image) {
	p.Plist_png = m
	p.Plist_sheet = s
}

// 设置玩家状态
func (p *PlayerBase) SetPlayerState(s, d uint8) {
	p.State = s
	p.Direction = d
}

// 获取图片
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

// 渲染角色
func (p *PlayerBase) Render() {
	p.Count++
	//Change player Frame
	if p.Count > p.FrameSpeed {
		p.Counts++
		if (p.State == tools.ATTACK || p.State == tools.SkILL) && p.Counts == p.FrameNums {
			p.SetPlayerState(tools.IDLE, p.Direction)
		}
		p.Counts %= p.FrameNums
		p.Count = 0
	}
}

// GC
func (p *PlayerBase) GC() {
	p.Plist_sheet = nil
	p.Plist_sheet_2 = nil
	p.Plist_png = nil
	p.Plist_png_2 = nil
}

// 改变帧数
func (p *PlayerBase) ChangeFrame() {
	//根据状态改变帧数
	if p.State == tools.IDLE {
		p.FrameNums = 16
		p.FrameSpeed = 5
	} else if p.State == tools.ATTACK {
		p.FrameNums = 16
		p.FrameSpeed = 1
	} else if p.State == tools.SkILL {
		p.FrameNums = 14
		p.FrameSpeed = 1
	} else {
		p.FrameNums = 8
		p.FrameSpeed = 3
	}
}

package layout

import (
	"embed"
	"fmt"
	"game/maps"
	"game/status"
	"game/tools"
	"image"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	plist_png     *image.NRGBA
	plist_R_png   *image.NRGBA
	plist_sheet   *texturepacker.SpriteSheet
	plist_R_sheet *texturepacker.SpriteSheet
	isClick       bool = false
)

//Create UI Class
type UI struct {
	image             *embed.FS
	Compents          []*icon
	HiddenCompents    []*icon
	MiniPanelCompents []*icon
	status            *status.StatusManage
	maps              *maps.MapBase
}

func NewUI(images *embed.FS, s *status.StatusManage, m *maps.MapBase) *UI {
	ui := &UI{
		image:             images,
		Compents:          make([]*icon, 0, 12),
		HiddenCompents:    make([]*icon, 0, 6),
		MiniPanelCompents: make([]*icon, 0, 6),
		status:            s,
		maps:              m,
	}
	return ui
}

func (u *UI) LoadGameImages() {
	u.ClearSlice(10)
	var len float64 = 0
	// go func() {
	// 	plist, _ := u.image.ReadFile("resource/UI/0000.png")
	// 	plist_json, _ := u.image.ReadFile("resource/man/warrior/ba.json")
	// 	plist_sheet, plist_png = tools.GetImageFromPlistPaletted(plist, plist_json)
	// 	runtime.GC()
	// }()
	s, _ := u.image.ReadFile("resource/UI/0000.png")
	mgUI := tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), tools.ISNORCOM)

	s, _ = u.image.ReadFile("resource/UI/HP.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(28, 480-float64(mgUI.Bounds().Max.Y+13), mgUI, 0, nil), tools.ISNORCOM)

	len += 115

	s, _ = u.image.ReadFile("resource/UI/chisha.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0001.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0002.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0003.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0004.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), tools.ISNORCOM)

	s, _ = u.image.ReadFile("resource/UI/liehuo.png")
	mgUI1 := tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(627, 480-float64(mgUI1.Bounds().Max.Y), mgUI1, 0, nil), tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0005.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), tools.ISNORCOM)

	s, _ = u.image.ReadFile("resource/UI/MP.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(684, 480-float64(mgUI.Bounds().Max.Y+13), mgUI, 1, nil), tools.ISNORCOM)

	s, _ = u.image.ReadFile("resource/UI/skill_btn.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(204, 441, mgUI, 0, func(i *icon) {
		if isClick == false {
			isClick = true
			go func() {
				on := *i.images
				s, _ = u.image.ReadFile("resource/UI/skill_btn_down.png")
				mgUI = tools.GetEbitenImage(s)
				i.images = mgUI
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.images = &on
				runtime.GC()
				isClick = false
			}()
		}
	}, 201, 446, 228, 468), tools.ISNORCOM)
	u.AddComponent(QuickCreate(562, 441, mgUI, 0, func(i *icon) {
		if isClick == false {
			isClick = true
			go func() {
				on := *i.images
				s, _ = u.image.ReadFile("resource/UI/skill_btn_down.png")
				mgUI = tools.GetEbitenImage(s)
				i.images = mgUI
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.images = &on
				runtime.GC()
				isClick = false
			}()
		}
	}, 559, 443, 586, 472), tools.ISNORCOM)

	//Draw Eq
	s, _ = u.image.ReadFile("resource/UI/eq_0.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395, 0, mgUI, 0, nil), tools.ISHIDDENCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/eq_1.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395+256, 0, mgUI, 0, nil), tools.ISHIDDENCOM)

	len = 395
	s, _ = u.image.ReadFile("resource/UI/bag_0.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395, 176, mgUI, 0, nil), tools.ISHIDDENCOM)

	len += float64(mgUI.Bounds().Max.X)
	s, _ = u.image.ReadFile("resource/UI/bag_1.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395+256, 176, mgUI, 0, nil), tools.ISHIDDENCOM)

	//Close btn
	s, _ = u.image.ReadFile("resource/UI/close_btn_on.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(414, 384, mgUI, 0, func(i *icon) {
		if isClick == false {
			isClick = true
			go func() {
				on := *i.images
				s, _ = u.image.ReadFile("resource/UI/close_btn_down.png")
				mgUI = tools.GetEbitenImage(s)
				i.images = mgUI
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.images = &on
				u.setHidden(tools.ISHIDDENCOM)
				go func() {
					for _, v := range u.MiniPanelCompents {
						v.SetPosition(100, 0)
						if v.clickMaxX != 0 {
							v.clickMaxX += 100
							v.clickMinX += 100
						}
					}
				}()
				runtime.GC()
				//恢复因打开包裹导致的人物偏移
				u.status.UIOFFSETX = 0
				//恢复影子偏移
				u.status.ShadowOffsetX = -350
				u.status.ShadowOffsetY = 365
				//恢复地图偏移
				u.maps.ChangeMapTranslate(200, 0)
				isClick = false
			}()
		}
	}, 412, 388, 439, 416), tools.ISHIDDENCOM)

	//循环获取uixy.go 里登录的物品信息
	items := getItems()
	for k, v := range items {
		s, _ = u.image.ReadFile("resource/UI/" + k + ".png")
		mgUI = tools.GetEbitenImage(s)
		for _, b := range v {
			res := strings.Split(b, "_")
			x, _ := strconv.ParseFloat(res[0], 64)
			y, _ := strconv.ParseFloat(res[1], 64)
			lay, _ := strconv.Atoi(res[2])
			re, _ := strconv.Atoi(res[3])
			u.AddComponent(QuickCreate(x, y, mgUI, uint8(lay), nil), uint8(re))
		}
	}
	//注册mini板打开按钮
	s, _ = u.image.ReadFile("resource/UI/open_minipanel_btn.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(390, 443, mgUI, 0, func(i *icon) {
		if isClick == false {
			isClick = true
			go func() {
				if u.status.OpenMiniPanel {
					u.setHidden(tools.ISMINICOM)
					s, _ = u.image.ReadFile("resource/UI/open_minipanel_down.png")
					mgUI = tools.GetEbitenImage(s)
					i.images = mgUI
					time.Sleep(tools.CLOSEBTNSLEEP)
					s, _ = u.image.ReadFile("resource/UI/close_minipanel_btn.png")
					mgUI = tools.GetEbitenImage(s)
					i.images = mgUI
				} else {
					u.SetDisplay(tools.ISMINICOM)
					s, _ = u.image.ReadFile("resource/UI/close_minipanel_down.png")
					mgUI = tools.GetEbitenImage(s)
					i.images = mgUI
					time.Sleep(tools.CLOSEBTNSLEEP)
					s, _ = u.image.ReadFile("resource/UI/open_minipanel_btn.png")
					mgUI = tools.GetEbitenImage(s)
					i.images = mgUI
				}
				runtime.GC()
				isClick = false
			}()
		}
	}, 387, 448, 397, 467), tools.ISNORCOM)
	//注册mini板
	s, _ = u.image.ReadFile("resource/UI/miniPanel.png")
	mgUI = tools.GetEbitenImage(s)
	baseX := float64(tools.LAYOUTX/2 - mgUI.Bounds().Max.X/2)
	u.AddComponent(QuickCreate(baseX, 406, mgUI, 0, nil), tools.ISMINICOM)
	baseX += 4
	//
	s, _ = u.image.ReadFile("resource/UI/mini_menu_man.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, nil), tools.ISMINICOM)
	//
	baseX += float64(mgUI.Bounds().Max.X) + 4
	s, _ = u.image.ReadFile("resource/UI/mini_menu_wea.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, func(i *icon) {
		if isClick == false {
			isClick = true
			go func() {
				on := *i.images
				s, _ = u.image.ReadFile("resource/UI/mini_menu_wea_down.png")
				mgUI = tools.GetEbitenImage(s)
				if !u.status.OpenBag {
					if x, _ := u.MiniPanelCompents[0].GetPosition(); x > 209 {
						go func() {
							//设置因打开包裹导致的人物偏移
							u.status.UIOFFSETX = -200
							//修改地图偏移
							u.maps.ChangeMapTranslate(-200, 0)
							//修改人物影子偏移
							u.status.ShadowOffsetX = u.status.ShadowOffsetX + 14
							u.status.ShadowOffsetY = u.status.ShadowOffsetY - 79
							for _, v := range u.MiniPanelCompents {
								v.SetPosition(-100, 0)
								if v.clickMaxX != 0 {
									//按钮点击范围偏移
									v.clickMaxX -= 100
									v.clickMinX -= 100
								}
							}
						}()
					}
				}
				i.images = mgUI
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.images = &on
				u.SetDisplay(tools.ISHIDDENCOM)
				runtime.GC()
				isClick = false
			}()
		}
	}, 337, 416, 350, 429), tools.ISMINICOM)
	baseX += float64(mgUI.Bounds().Max.X) + 4
	s, _ = u.image.ReadFile("resource/UI/mini_menu_j.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, nil), tools.ISMINICOM)
	baseX += float64(mgUI.Bounds().Max.X) + 4
	s, _ = u.image.ReadFile("resource/UI/mini_menu_m.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, nil), tools.ISMINICOM)
	baseX += float64(mgUI.Bounds().Max.X) + 4
	s, _ = u.image.ReadFile("resource/UI/mini_menu_mess.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, nil), tools.ISMINICOM)
	baseX += float64(mgUI.Bounds().Max.X) + 4
	s, _ = u.image.ReadFile("resource/UI/mini_menu_s.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, nil), tools.ISMINICOM)
	baseX += float64(mgUI.Bounds().Max.X) + 4
	s, _ = u.image.ReadFile("resource/UI/mini_menu_st.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, nil), tools.ISMINICOM)

	u.setHidden(tools.ISHIDDENCOM)
	u.setHidden(tools.ISMINICOM)

}

func (u *UI) LoadGameLoginImages() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("has error is :", r)
		}
	}()
	var len float64 = 0
	var scales float64 = 0.8
	s, _ := u.image.ReadFile("resource/UI/login0.png")
	mgUI := tools.GetEbitenImage(s)
	op := newIcon()
	op.SetPosition(len, 0)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login1.png")
	mgUI = tools.GetEbitenImage(s)
	op1 := newIcon()
	op1.SetPosition(len, 0)
	op1.op.GeoM.Scale(1, scales)
	op1.addImage(mgUI)
	u.AddComponent(op1, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login2.png")
	mgUI = tools.GetEbitenImage(s)
	op2 := newIcon()
	op2.SetPosition(len, 0)
	op2.op.GeoM.Scale(1, scales)
	op2.addImage(mgUI)
	u.AddComponent(op2, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login3.png")
	mgUI = tools.GetEbitenImage(s)
	op3 := newIcon()
	op3.SetPosition(len, 0)
	op3.op.GeoM.Scale(1, scales)
	op3.addImage(mgUI)
	u.AddComponent(op3, tools.ISNORCOM)

	len = 0
	var offset float64 = 340
	s, _ = u.image.ReadFile("resource/UI/login8.png")
	mgUI = tools.GetEbitenImage(s)
	op8 := newIcon()
	op8.SetPosition(len, offset)
	op8.op.GeoM.Scale(1, scales)
	op8.addImage(mgUI)
	u.AddComponent(op8, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login9.png")
	mgUI = tools.GetEbitenImage(s)
	op9 := newIcon()
	op9.SetPosition(len, offset)
	op9.op.GeoM.Scale(1, scales)
	op9.addImage(mgUI)
	u.AddComponent(op9, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login10.png")
	mgUI = tools.GetEbitenImage(s)
	op10 := newIcon()
	op10.SetPosition(len, offset)
	op10.op.GeoM.Scale(1, scales)
	op10.addImage(mgUI)
	u.AddComponent(op10, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login11.png")
	mgUI = tools.GetEbitenImage(s)
	op11 := newIcon()
	op11.SetPosition(len, offset)
	op11.op.GeoM.Scale(1, scales)
	op11.addImage(mgUI)
	u.AddComponent(op11, tools.ISNORCOM)

	len = 0

	s, _ = u.image.ReadFile("resource/UI/login4.png")
	mgUI = tools.GetEbitenImage(s)
	op4 := newIcon()
	op4.SetPosition(len, float64(mgUI.Bounds().Max.Y))
	op4.op.GeoM.Scale(1, scales)
	op4.addImage(mgUI)
	u.AddComponent(op4, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login5.png")
	mgUI = tools.GetEbitenImage(s)
	op5 := newIcon()
	op5.SetPosition(len, float64(mgUI.Bounds().Max.Y))
	op5.op.GeoM.Scale(1, scales)
	op5.addImage(mgUI)
	u.AddComponent(op5, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login6.png")
	mgUI = tools.GetEbitenImage(s)
	op6 := newIcon()
	op6.SetPosition(len, float64(mgUI.Bounds().Max.Y))
	op6.op.GeoM.Scale(1, scales)
	op6.addImage(mgUI)
	u.AddComponent(op6, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login7.png")
	mgUI = tools.GetEbitenImage(s)
	op7 := newIcon()
	op7.SetPosition(len, float64(mgUI.Bounds().Max.Y))
	op7.op.GeoM.Scale(1, scales)
	op7.addImage(mgUI)
	u.AddComponent(op7, tools.ISNORCOM)

	go func() {
		plist, _ := u.image.ReadFile("resource/UI/logo.png")
		plist_json, _ := u.image.ReadFile("resource/UI/logo.json")
		plist_sheet, plist_png = tools.GetImageFromPlist(plist, plist_json)
		runtime.GC()
	}()

}

func (u *UI) LoadGameCharaSelectImages() {
	u.ClearSlice(1)
	s, _ := u.image.ReadFile("resource/UI/charactSelect.png")
	mgUI := tools.GetEbitenImage(s)
	op := newIcon()
	op.SetPosition(0, 0)
	op.op.GeoM.Scale(1, 0.8)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)
	w := &sync.WaitGroup{}
	w.Add(2)
	go func() {
		plist, _ := u.image.ReadFile("resource/UI/logo.png")
		plist_json, _ := u.image.ReadFile("resource/UI/logo.json")
		plist_sheet, plist_png = tools.GetImageFromPlist(plist, plist_json)
		w.Done()
	}()
	go func() {
		plist, _ := u.image.ReadFile("resource/UI/selectRoles.png")
		plist_json, _ := u.image.ReadFile("resource/UI/selectRoles.json")
		plist_R_sheet, plist_R_png = tools.GetImageFromPlist(plist, plist_json)
		w.Done()
	}()
	w.Wait()
	go func() {
		runtime.GC()
	}()

}

func (u *UI) GetAnimator(flg, name string) (*ebiten.Image, int, int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if flg == "role" {

		return ebiten.NewImageFromImage(plist_R_png.SubImage(plist_R_sheet.Sprites[name].Frame)), plist_R_sheet.Sprites[name].SpriteSourceSize.Min.X, plist_R_sheet.Sprites[name].SpriteSourceSize.Min.Y

	} else {

		return ebiten.NewImageFromImage(plist_png.SubImage(plist_sheet.Sprites[name].Frame)), plist_sheet.Sprites[name].SpriteSourceSize.Min.X, plist_sheet.Sprites[name].SpriteSourceSize.Min.Y
	}
}

//Add Component
func (u *UI) AddComponent(s *icon, ImageType uint8) {
	u.Compents = append(u.Compents, s)
	if ImageType == tools.ISHIDDENCOM {
		u.HiddenCompents = append(u.HiddenCompents, s)
	} else if ImageType == tools.ISMINICOM {
		u.MiniPanelCompents = append(u.MiniPanelCompents, s)
	}
}

//
func (u *UI) SetDisplay(ImageType uint8) {
	if ImageType == tools.ISHIDDENCOM {
		u.status.OpenBag = true
		for _, v := range u.HiddenCompents {
			v.isDisplay = true
		}
	} else {
		u.status.OpenMiniPanel = true
		for _, v := range u.MiniPanelCompents {
			v.isDisplay = true
		}
	}

}

//
func (u *UI) setHidden(ImageType uint8) {
	if ImageType == tools.ISHIDDENCOM {
		u.status.OpenBag = false
		for _, v := range u.HiddenCompents {
			v.isDisplay = false
		}
	} else {
		u.status.OpenMiniPanel = false
		for _, v := range u.MiniPanelCompents {
			v.isDisplay = false
		}
	}

}

func (u *UI) ClearSlice(cap int) {
	u.Compents = make([]*icon, 0, cap)
	u.HiddenCompents = make([]*icon, 0, cap/2)
	u.MiniPanelCompents = make([]*icon, 0, cap/2)
}

//Render UI
func (u *UI) DrawUI(screen *ebiten.Image) {
	for _, v := range u.Compents {
		if v.layer == 0 && v.isDisplay {
			screen.DrawImage(v.images, v.op)
		}
	}
	for _, v := range u.Compents {
		if v.layer == 1 && v.isDisplay {
			screen.DrawImage(v.images, v.op)
		}
	}
}

//Event Listen
func (u *UI) EventLoop() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, v := range u.Compents {
			if v.hasEvent == 1 && v.isDisplay {
				x, y := ebiten.CursorPosition()
				if x >= v.clickMinX && x <= v.clickMaxX && y >= v.clickMinY && y <= v.clickMaxY {
					//设置不可以行走
					u.status.Flg = false
					//call back func
					v.f(v)
				}
			}
		}
	}

}

//GC for loading
func (u *UI) ClearGlobalVariable() {
	plist_R_sheet = nil
	plist_R_png = nil
}

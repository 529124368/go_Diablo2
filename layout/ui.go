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
	Compents          []*Sprite            //普通UI存放集合
	HiddenCompents    []*Sprite            //可以被隐藏的UI组件集合
	MiniPanelCompents []*Sprite            //MINI板的UI集合
	ItemsCompents     []*SpriteItems       //Items的UI集合
	status            *status.StatusManage //状态管理器
	maps              *maps.MapBase        //地图
}

func NewUI(images *embed.FS, s *status.StatusManage, m *maps.MapBase) *UI {
	ui := &UI{
		image:             images,
		Compents:          make([]*Sprite, 0, 12),
		HiddenCompents:    make([]*Sprite, 0, 6),
		MiniPanelCompents: make([]*Sprite, 0, 6),
		ItemsCompents:     make([]*SpriteItems, 0, 50),
		status:            s,
		maps:              m,
	}
	return ui
}

//加载进入游戏UI
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
	u.AddComponent(QuickCreate(204, 441, mgUI, 0, func(i spriteInterface) {
		if isClick == false {
			isClick = true
			go func() {
				on := *i.(*Sprite).images
				s, _ = u.image.ReadFile("resource/UI/skill_btn_down.png")
				mgUI = tools.GetEbitenImage(s)
				i.(*Sprite).images = mgUI
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).images = &on
				runtime.GC()
				isClick = false
			}()
		}
	}, 201, 446, 228, 468), tools.ISNORCOM)
	u.AddComponent(QuickCreate(562, 441, mgUI, 0, func(i spriteInterface) {
		if isClick == false {
			isClick = true
			go func() {
				on := *i.(*Sprite).images
				s, _ = u.image.ReadFile("resource/UI/skill_btn_down.png")
				mgUI = tools.GetEbitenImage(s)
				i.(*Sprite).images = mgUI
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).images = &on
				runtime.GC()
				isClick = false
			}()
		}
	}, 559, 443, 586, 472), tools.ISNORCOM)

	//描画装备栏和包裹UI
	s, _ = u.image.ReadFile("resource/UI/eq_0.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395, 0, mgUI, 0, nil), tools.ISHIDDEN)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/eq_1.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395+256, 0, mgUI, 0, nil), tools.ISHIDDEN)

	len = 395
	s, _ = u.image.ReadFile("resource/UI/bag_0.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395, 176, mgUI, 0, nil), tools.ISHIDDEN)

	len += float64(mgUI.Bounds().Max.X)
	s, _ = u.image.ReadFile("resource/UI/bag_1.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395+256, 176, mgUI, 0, nil), tools.ISHIDDEN)

	//Close btn
	s, _ = u.image.ReadFile("resource/UI/close_btn_on.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(414, 384, mgUI, 0, func(i spriteInterface) {
		if isClick == false {
			isClick = true
			go func() {
				on := *i.(*Sprite).images
				s, _ = u.image.ReadFile("resource/UI/close_btn_down.png")
				mgUI = tools.GetEbitenImage(s)
				i.(*Sprite).images = mgUI
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).images = &on
				u.setHidden(tools.ISHIDDEN)
				go func() {
					for _, v := range u.MiniPanelCompents {
						v.SetPosition(100, 0)

						if v.clickMaxX != 0 {
							v.clickMinX += 100
							v.clickMaxX += 100
						}
					}
				}()
				runtime.GC()
				//恢复因打开包裹导致的人物偏移
				u.status.UIOFFSETX = 0
				//恢复影子偏移
				u.status.ShadowOffsetX = -350
				u.status.ShadowOffsetY = 365
				//恢复玩家中心位置
				u.status.PLAYERCENTERX = 388
				//恢复地图偏移
				u.maps.ChangeMapTranslate(200, 0)
				isClick = false
			}()
		}
	}, 412, 388, 439, 416), tools.ISHIDDEN)

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
			u.AddComponent(QuickCreateItems(x, y, mgUI, uint8(lay), nil), uint8(re))
		}
	}

	//注册mini板打开按钮
	s, _ = u.image.ReadFile("resource/UI/open_minipanel_btn.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(390, 443, mgUI, 0, func(i spriteInterface) {
		if isClick == false {
			isClick = true
			go func() {
				if u.status.OpenMiniPanel {
					u.setHidden(tools.ISMINICOM)
					s, _ = u.image.ReadFile("resource/UI/open_minipanel_down.png")
					mgUI = tools.GetEbitenImage(s)
					i.(*Sprite).images = mgUI
					time.Sleep(tools.CLOSEBTNSLEEP)
					s, _ = u.image.ReadFile("resource/UI/close_minipanel_btn.png")
					mgUI = tools.GetEbitenImage(s)
					i.(*Sprite).images = mgUI
				} else {
					u.SetDisplay(tools.ISMINICOM)
					s, _ = u.image.ReadFile("resource/UI/close_minipanel_down.png")
					mgUI = tools.GetEbitenImage(s)
					i.(*Sprite).images = mgUI
					time.Sleep(tools.CLOSEBTNSLEEP)
					s, _ = u.image.ReadFile("resource/UI/open_minipanel_btn.png")
					mgUI = tools.GetEbitenImage(s)
					i.(*Sprite).images = mgUI
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
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, func(i spriteInterface) {
		if isClick == false {
			isClick = true
			go func() {
				on := *i.(*Sprite).images
				s, _ = u.image.ReadFile("resource/UI/mini_menu_wea_down.png")
				mgUI = tools.GetEbitenImage(s)
				if !u.status.OpenBag {
					//判断MINI板子的最左端坐标是否超过最大极限
					if x, _ := u.MiniPanelCompents[0].GetPosition(); x > 209 {
						go func() {
							//设置因打开包裹导致的人物偏移
							u.status.UIOFFSETX = -200
							//修改地图偏移
							u.maps.ChangeMapTranslate(-200, 0)
							//修改玩家中心位置
							u.status.PLAYERCENTERX -= 200
							//修改人物影子偏移
							u.status.ShadowOffsetX = u.status.ShadowOffsetX + 14
							u.status.ShadowOffsetY = u.status.ShadowOffsetY - 79
							for _, v := range u.MiniPanelCompents {
								v.SetPosition(-100, 0)
								if v.clickMaxX != 0 {
									v.clickMinX -= 100
									v.clickMaxX -= 100
								}
							}
						}()
					}
				}
				i.(*Sprite).images = mgUI
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).images = &on
				u.SetDisplay(tools.ISHIDDEN)
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

	u.setHidden(tools.ISHIDDEN)
	u.setHidden(tools.ISMINICOM)

}

//加载登录游戏UI
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
	op := newSprite()
	op.SetPosition(len, 0)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login1.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, 0)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login2.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, 0)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login3.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, 0)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len = 0
	var offset float64 = 340
	s, _ = u.image.ReadFile("resource/UI/login8.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, offset)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login9.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, offset)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login10.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, offset)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login11.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, offset)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len = 0

	s, _ = u.image.ReadFile("resource/UI/login4.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, float64(mgUI.Bounds().Max.Y))
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login5.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, float64(mgUI.Bounds().Max.Y))
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login6.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, float64(mgUI.Bounds().Max.Y))
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login7.png")
	mgUI = tools.GetEbitenImage(s)
	op = newSprite()
	op.SetPosition(len, float64(mgUI.Bounds().Max.Y))
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, tools.ISNORCOM)

	go func() {
		plist, _ := u.image.ReadFile("resource/UI/logo.png")
		plist_json, _ := u.image.ReadFile("resource/UI/logo.json")
		plist_sheet, plist_png = tools.GetImageFromPlist(plist, plist_json)
		runtime.GC()
	}()

}

//加载游戏选择角色UI
func (u *UI) LoadGameCharaSelectImages() {
	u.ClearSlice(1)
	s, _ := u.image.ReadFile("resource/UI/charactSelect.png")
	mgUI := tools.GetEbitenImage(s)
	op := newSprite()
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

//图集获取图片
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

//组件注册
func (u *UI) AddComponent(s spriteInterface, ImageType uint8) {
	if ImageType == tools.ISHIDDEN {
		//将UI压入通用集合
		u.Compents = append(u.Compents, s.(*Sprite))
		//将UI压入可隐藏集合
		u.HiddenCompents = append(u.HiddenCompents, s.(*Sprite))
	} else if ImageType == tools.ISMINICOM {
		//将UI压入通用集合
		u.Compents = append(u.Compents, s.(*Sprite))
		//将UI压入MINI板集合
		u.MiniPanelCompents = append(u.MiniPanelCompents, s.(*Sprite))
	} else if ImageType == tools.ISITEMS {
		//将UI压入Items集合
		u.ItemsCompents = append(u.ItemsCompents, s.(*SpriteItems))
	} else {
		//将UI压入通用集合
		u.Compents = append(u.Compents, s.(*Sprite))
	}
}

//显示UI
func (u *UI) SetDisplay(ImageType uint8) {
	if ImageType == tools.ISHIDDEN || ImageType == tools.ISITEMS {
		u.status.OpenBag = true
		for _, v := range u.HiddenCompents {
			v.isDisplay = true
		}
		for _, v := range u.ItemsCompents {
			v.isDisplay = true
		}
	} else {
		u.status.OpenMiniPanel = true
		for _, v := range u.MiniPanelCompents {
			v.isDisplay = true
		}
	}

}

//隐藏UI
func (u *UI) setHidden(ImageType uint8) {
	if ImageType == tools.ISHIDDEN || ImageType == tools.ISITEMS {
		u.status.OpenBag = false
		for _, v := range u.HiddenCompents {
			v.isDisplay = false
		}
		for _, v := range u.ItemsCompents {
			v.isDisplay = false
		}
	} else {
		u.status.OpenMiniPanel = false
		for _, v := range u.MiniPanelCompents {
			v.isDisplay = false
		}
	}

}

//清除切片
func (u *UI) ClearSlice(cap int) {
	u.Compents = make([]*Sprite, 0, cap)
	u.HiddenCompents = make([]*Sprite, 0, cap/2)
	u.MiniPanelCompents = make([]*Sprite, 0, cap/2)
	u.ItemsCompents = make([]*SpriteItems, 0, 50)
}

//渲染UI
func (u *UI) DrawUI(screen *ebiten.Image) {
	//渲染UI
	for _, v := range u.Compents {
		if v.layer == 0 && v.isDisplay {
			screen.DrawImage(v.images, v.op)
		}
	}
	//渲染层级为1的UI
	for _, v := range u.Compents {
		if v.layer == 1 && v.isDisplay {
			screen.DrawImage(v.images, v.op)
		}
	}
	//当包裹打开的时候，渲染包裹内物品和装备
	if u.status.OpenBag {
		for _, v := range u.ItemsCompents {
			screen.DrawImage(v.imageBg, v.opBg)
			screen.DrawImage(v.images, v.op)
		}

	}
}

//事件轮询
func (u *UI) EventLoop() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, v := range u.Compents {
			if v.hasEvent == 1 && v.isDisplay {
				x, y := ebiten.CursorPosition()
				if x >= v.clickMinX && x <= v.clickMaxX && y >= v.clickMinY && y <= v.clickMaxY {
					//设置不可以行走
					u.status.Flg = false
					//实行回调函数
					v.f(v)
				}
			}
		}
	}

}

//GC 加载清楚变量
func (u *UI) ClearGlobalVariable() {
	plist_R_sheet = nil
	plist_R_png = nil
}

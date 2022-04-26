package layout

import (
	"game/interfaces"
	"game/tools"
	"runtime"
	"strings"
	"time"
)

//加载进入游戏UI
func (u *UI) LoadGameImages() {
	u.ClearSlice(10)
	//
	var len float64 = 0
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
	u.AddComponent(QuickCreate(204, 441, mgUI, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
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
	}, true), tools.ISNORCOM)
	u.AddComponent(QuickCreate(562, 441, mgUI, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
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
	}, true), tools.ISNORCOM)

	//注册跑步走路动作切换
	s, _ = u.image.ReadFile("resource/UI/walk_button.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(254, 451, mgUI, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				if u.status.IsWalk {
					s, _ = u.image.ReadFile("resource/UI/walk_button_down.png")
					mgUI = tools.GetEbitenImage(s)
					i.(*Sprite).images = mgUI
					time.Sleep(tools.CLOSEBTNSLEEP)
					s, _ = u.image.ReadFile("resource/UI/run_button.png")
					mgUI = tools.GetEbitenImage(s)
					i.(*Sprite).images = mgUI
					u.status.IsWalk = false
				} else {
					s, _ = u.image.ReadFile("resource/UI/run_button_down.png")
					mgUI = tools.GetEbitenImage(s)
					i.(*Sprite).images = mgUI
					time.Sleep(tools.CLOSEBTNSLEEP)
					s, _ = u.image.ReadFile("resource/UI/walk_button.png")
					mgUI = tools.GetEbitenImage(s)
					i.(*Sprite).images = mgUI
					u.status.IsWalk = true
				}
				runtime.GC()
				isClick = false
			}()
		}
	}, true), tools.ISNORCOM)

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

	//关闭装备栏按钮
	s, _ = u.image.ReadFile("resource/UI/close_btn_on.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(414, 384, mgUI, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
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
				u.status.ShadowOffsetX = -345
				u.status.ShadowOffsetY = 345
				//恢复玩家中心位置
				u.status.PLAYERCENTERX = 388
				//恢复地图偏移
				u.mapManage.ChangeMapTranslate(200, 0)
				isClick = false
			}()
		}
	}, true), tools.ISHIDDEN)

	//背包物品LOOP Start
	//临时Map
	TempArray := make(map[string]int, 10)
	items := u.bag.BagLayout
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if strings.Contains(items[i][j], "_") {
				if _, ok := TempArray[items[i][j]]; !ok {
					TempArray[items[i][j]] = 0
					t := strings.Split(items[i][j], "_")
					s, _ = u.image.ReadFile("resource/items/" + t[0] + ".png")
					mgUI = tools.GetEbitenImage(s)
					x := 413 + j*29
					y := 254 + i*29
					u.AddComponent(QuickCreateItems(float64(x), float64(y), t[0], mgUI, 1, u.ItemsEvent(), 1, true), tools.ISITEMS)
				}
			} else if items[i][j] != "" {
				s, _ = u.image.ReadFile("resource/items/" + items[i][j] + ".png")
				mgUI = tools.GetEbitenImage(s)
				x := 413 + j*29
				y := 254 + i*29
				u.AddComponent(QuickCreateItems(float64(x), float64(y), items[i][j], mgUI, 1, u.ItemsEvent(), 1, true), tools.ISITEMS)
			}
		}
	}
	//手动销毁临时Map
	TempArray = nil
	//背包物品LOOP END
	xx, yy := 0, 0
	for i := 0; i < 10; i++ {
		//装备栏初始化
		if u.bag.BagLayout[4][i] != "" {
			if i == 0 {
				xx, yy = 530, 3
			}
			if i == 1 {
				xx, yy = 416, 60
			}
			if i == 2 {
				xx, yy = 647, 60
			}
			if i == 3 {
				xx, yy = 600, 32
			}
			if i == 4 {
				xx, yy = 530, 74
			}
			if i == 5 {
				xx, yy = 414, 177
			}
			if i == 6 {
				xx, yy = 487, 177
			}
			if i == 7 {
				xx, yy = 529, 178
			}
			if i == 8 {
				xx, yy = 600, 177
			}
			if i == 9 {
				xx, yy = 646, 177
			}

			//插入装备
			u.InsertToEquip(xx, yy, u.bag.BagLayout[4][i])
		}
	}

	//注册mini板打开按钮
	s, _ = u.image.ReadFile("resource/UI/open_minipanel_btn.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(390, 443, mgUI, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
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
	}, true), tools.ISNORCOM)
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
	s, _ = u.image.ReadFile("resource/UI/mini_menu_eq.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				on := *i.(*Sprite).images
				s, _ = u.image.ReadFile("resource/UI/mini_menu_eq_down.png")
				mgUI = tools.GetEbitenImage(s)
				if !u.status.OpenBag {
					//判断MINI板子的最左端坐标是否超过最大极限
					if x, _ := u.MiniPanelCompents[0].GetPosition(); x > 209 {
						go func() {
							//设置因打开包裹导致的人物偏移
							u.status.UIOFFSETX = -200
							//修改地图偏移
							u.mapManage.ChangeMapTranslate(-200, 0)
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
	}, true), tools.ISMINICOM)
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

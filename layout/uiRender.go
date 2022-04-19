package layout

import (
	"fmt"
	"game/tools"
	"runtime"
	"strings"
	"sync"
	"time"
)

//加载进入游戏UI
func (u *UI) LoadGameImages() {
	u.ClearSlice(10)
	//初始化背包 数据
	itemsLayout := [5][10]string{
		{"HP0", "HP0", "book_0,2", "dun-6_0,3", "dun-6_0,3", "sword_0,5", "sword_0,5", "", "dun_0,8", "dun_0,8"},
		{"body-3_1,0", "body-3_1,0", "book_0,2", "dun-6_0,3", "dun-6_0,3", "sword_0,5", "sword_0,5", "", "dun_0,8", "dun_0,8"},
		{"body-3_1,0", "body-3_1,0", "", "dun-6_0,3", "dun-6_0,3", "sword_0,5", "sword_0,5", "", "head-5_2,8", "head-5_2,8"},
		{"body-3_1,0", "body-3_1,0", "", "", "", "", "", "", "head-5_2,8", "head-5_2,8"},
		{"", "", "", "", "", "", "", "", "", ""},
		//头盔526,8  左手武器412,54 右手武器644,54 项链599,36 铠甲526,80 手套413,182 左戒指485,181 腰带527,181 右戒指599,183 靴子644,183
	}
	u.BagLayout = itemsLayout
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
	u.AddComponent(QuickCreate(204, 441, mgUI, 0, func(i spriteInterface) {
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
	u.AddComponent(QuickCreate(562, 441, mgUI, 0, func(i spriteInterface) {
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
	u.AddComponent(QuickCreate(414, 384, mgUI, 0, func(i spriteInterface) {
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
				u.status.ShadowOffsetX = -350
				u.status.ShadowOffsetY = 365
				//恢复玩家中心位置
				u.status.PLAYERCENTERX = 388
				//恢复地图偏移
				u.maps.ChangeMapTranslate(200, 0)
				isClick = false
			}()
		}
	}, true), tools.ISHIDDEN)

	//背包物品LOOP Start
	//临时Map
	TempArray := make(map[string]int, 10)
	items := u.BagLayout
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if strings.Contains(items[i][j], "_") {
				if _, ok := TempArray[items[i][j]]; !ok {
					TempArray[items[i][j]] = 0
					t := strings.Split(items[i][j], "_")
					s, _ = u.image.ReadFile("resource/UI/" + t[0] + ".png")
					mgUI = tools.GetEbitenImage(s)
					x := 413 + j*29
					y := 254 + i*29
					u.AddComponent(QuickCreateItems(float64(x), float64(y), t[0], mgUI, 1, u.ItemsEvent(), 1, true), 0)
				}
			} else if items[i][j] != "" {
				s, _ = u.image.ReadFile("resource/UI/" + items[i][j] + ".png")
				mgUI = tools.GetEbitenImage(s)
				x := 413 + j*29
				y := 254 + i*29
				u.AddComponent(QuickCreateItems(float64(x), float64(y), items[i][j], mgUI, 1, u.ItemsEvent(), 1, true), 0)
			}
		}
	}
	//手动销毁临时Map
	TempArray = nil
	//背包物品LOOP END

	//注册mini板打开按钮
	s, _ = u.image.ReadFile("resource/UI/open_minipanel_btn.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(390, 443, mgUI, 0, func(i spriteInterface) {
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
	u.AddComponent(QuickCreate(baseX, 410, mgUI, 0, func(i spriteInterface) {
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
	op := QuickCreate(len, 0, mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login1.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, 0, mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login2.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, 0, mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login3.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, 0, mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len = 0
	var offset float64 = 340
	s, _ = u.image.ReadFile("resource/UI/login8.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, offset, mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login9.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, offset, mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login10.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, offset, mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login11.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, offset, mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len = 0

	s, _ = u.image.ReadFile("resource/UI/login4.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, float64(mgUI.Bounds().Max.Y), mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login5.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, float64(mgUI.Bounds().Max.Y), mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login6.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, float64(mgUI.Bounds().Max.Y), mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login7.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(len, float64(mgUI.Bounds().Max.Y), mgUI, 0, nil)
	op.op.GeoM.Scale(1, scales)
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
	//清空
	u.ClearSlice(1)
	//角色选择背景图加载
	s, _ := u.image.ReadFile("resource/UI/charactSelect.png")
	mgUI := tools.GetEbitenImage(s)
	op := QuickCreate(0, 0, mgUI, 0, nil)
	op.op.GeoM.Scale(1, 0.8)
	u.AddComponent(op, tools.ISNORCOM)
	//游戏开始按钮
	s, _ = u.image.ReadFile("resource/UI/startGameButton.png")
	mgUI = tools.GetEbitenImage(s)
	op = QuickCreate(1250, 850, mgUI, 0, func(i spriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				on := *i.(*Sprite).images
				s, _ = u.image.ReadFile("resource/UI/startGameButton_down.png")
				mgUI = tools.GetEbitenImage(s)
				i.(*Sprite).images = mgUI
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).images = &on
				u.SetDisplay(tools.ISHIDDEN)
				//切换游戏场景到开门loading
				u.status.ChangeScenceFlg = true
				u.status.CurrentGameScence = tools.GAMESCENEOPENDOOR
				u.status.ChangeScenceFlg = false
				isClick = false
			}()
		}
	}, false)
	op.op.GeoM.Scale(0.5, 0.5)
	u.AddComponent(op, tools.ISNORCOM)
	//动画
	w := &sync.WaitGroup{}
	w.Add(2)
	//加载火焰
	go func() {
		plist, _ := u.image.ReadFile("resource/UI/logo.png")
		plist_json, _ := u.image.ReadFile("resource/UI/logo.json")
		plist_sheet, plist_png = tools.GetImageFromPlist(plist, plist_json)
		w.Done()
	}()
	//加载野蛮人
	go func() {
		plist, _ := u.image.ReadFile("resource/UI/selectRoles.png")
		plist_json, _ := u.image.ReadFile("resource/UI/selectRoles.json")
		plist_R_sheet, plist_R_png = tools.GetImageFromPlist(plist, plist_json)
		w.Done()
	}()
	w.Wait()
	//手动GC
	go func() {
		runtime.GC()
	}()

}

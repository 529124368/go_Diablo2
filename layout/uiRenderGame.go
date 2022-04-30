package layout

import (
	"game/interfaces"
	"game/tools"
	"runtime"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

//加载进入游戏UI
func (u *UI) LoadGameImages() {
	u.ClearSlice(10)
	//UI
	plist, _ := u.image.ReadFile("resource/UI/UI.png")
	plist_json, _ := u.image.ReadFile("resource/UI/UI.json")
	pli, pic := tools.GetImageFromPlistPaletted(plist, plist_json)
	plist_sheet = pli
	plist_png = ebiten.NewImageFromImage(pic)
	//itemns
	plist, _ = u.image.ReadFile("resource/items/items.png")
	plist_json, _ = u.image.ReadFile("resource/items/items.json")
	pli1, pic1 := tools.GetImageFromPlist(plist, plist_json)
	plist_R_sheet = pli1
	plist_R_png = ebiten.NewImageFromImage(pic1)
	runtime.GC()
	//
	var len float64 = 0
	name := "0000"
	_, Y, _ := u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(len, 480-float64(Y), name, plist_sheet, 0, nil), tools.ISNORCOM)

	name = "HP"
	_, Y, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(28, 480-float64(Y+13), name, plist_sheet, 0, nil), tools.ISNORCOM)

	len += 115

	name = "chisha"
	X, Y, _ := u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(len, 480-float64(Y), name, plist_sheet, 0, nil), tools.ISNORCOM)

	len += float64(X)

	name = "0001"
	X, Y, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(len, 480-float64(Y), name, plist_sheet, 0, nil), tools.ISNORCOM)

	len += float64(X)

	name = "0002"
	X, Y, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(len, 480-float64(Y), name, plist_sheet, 0, nil), tools.ISNORCOM)

	len += float64(X)

	name = "0003"
	X, Y, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(len, 480-float64(Y), name, plist_sheet, 0, nil), tools.ISNORCOM)

	len += float64(X)

	name = "0004"
	X, Y, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(len, 480-float64(Y), name, plist_sheet, 0, nil), tools.ISNORCOM)
	len += float64(X)

	name = "liehuo"
	_, Y, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(629, 480-float64(Y), name, plist_sheet, 0, nil), tools.ISNORCOM)

	name = "0005"
	_, Y, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(len, 480-float64(Y), name, plist_sheet, 0, nil), tools.ISNORCOM)

	name = "MP"
	_, Y, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(684, 480-float64(Y+13), name, plist_sheet, 1, nil), tools.ISNORCOM)

	name = "skill_btn"
	u.AddComponent(QuickCreate(204, 441, name, plist_sheet, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				//减血
				u.DeleHP(10)
				on := i.(*Sprite).imagesName
				i.(*Sprite).imagesName = "skill_btn_down"
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).imagesName = on
				runtime.GC()
				isClick = false
			}()
		}
	}, true), tools.ISNORCOM)
	u.AddComponent(QuickCreate(562, 441, name, plist_sheet, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				//减蓝
				u.DeleMP(10)
				on := i.(*Sprite).imagesName
				i.(*Sprite).imagesName = "skill_btn_down"
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).imagesName = on
				runtime.GC()
				isClick = false
			}()
		}
	}, true), tools.ISNORCOM)

	//注册跑步走路动作切换
	name = "walk_button"
	u.AddComponent(QuickCreate(254, 451, name, plist_sheet, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				if u.status.IsWalk {
					i.(*Sprite).imagesName = "walk_button_down"
					time.Sleep(tools.CLOSEBTNSLEEP)
					i.(*Sprite).imagesName = "run_button"
					u.status.IsWalk = false
				} else {
					i.(*Sprite).imagesName = "run_button_down"
					time.Sleep(tools.CLOSEBTNSLEEP)
					i.(*Sprite).imagesName = "walk_button"
					u.status.IsWalk = true
				}
				runtime.GC()
				isClick = false
			}()
		}
	}, true), tools.ISNORCOM)

	//描画装备栏和包裹UI
	name = "eq_0"
	X, _, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(395, 0, name, plist_sheet, 0, nil), tools.ISHIDDEN)

	len += float64(X)

	name = "eq_1"
	u.AddComponent(QuickCreate(395+256, 0, name, plist_sheet, 0, nil), tools.ISHIDDEN)

	len = 395
	name = "bag_0"
	X, _, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(395, 176, name, plist_sheet, 0, nil), tools.ISHIDDEN)

	len += float64(X)
	name = "bag_1"
	u.AddComponent(QuickCreate(395+256, 176, name, plist_sheet, 0, nil), tools.ISHIDDEN)

	//关闭装备栏按钮
	name = "close_btn_on"
	u.AddComponent(QuickCreate(414, 384, name, plist_sheet, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				on := i.(*Sprite).imagesName
				i.(*Sprite).imagesName = "close_btn_down"
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).imagesName = on
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
					x := 413 + j*29
					y := 254 + i*29
					u.AddComponent(QuickCreateItems(float64(x), float64(y), t[0], plist_R_sheet, 1, u.ItemsEvent(), 1, true), tools.ISITEMS)
				}
			} else if items[i][j] != "" {
				x := 413 + j*29
				y := 254 + i*29
				u.AddComponent(QuickCreateItems(float64(x), float64(y), items[i][j], plist_R_sheet, 1, u.ItemsEvent(), 1, true), tools.ISITEMS)
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
	name = "open_minipanel_btn"
	u.AddComponent(QuickCreate(390, 443, name, plist_sheet, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				if u.status.OpenMiniPanel {
					u.setHidden(tools.ISMINICOM)
					i.(*Sprite).imagesName = "open_minipanel_down"
					time.Sleep(tools.CLOSEBTNSLEEP)
					i.(*Sprite).imagesName = "close_minipanel_btn"
				} else {
					u.SetDisplay(tools.ISMINICOM)
					i.(*Sprite).imagesName = "close_minipanel_down"
					time.Sleep(tools.CLOSEBTNSLEEP)
					i.(*Sprite).imagesName = "open_minipanel_btn"
				}
				runtime.GC()
				isClick = false
			}()
		}
	}, true), tools.ISNORCOM)
	//注册mini板
	name = "miniPanel"
	X, _, _ = u.GetImagesSize(tools.PlistN, name)
	baseX := float64(tools.LAYOUTX/2 - X/2)
	u.AddComponent(QuickCreate(baseX, 406, name, plist_sheet, 0, nil), tools.ISMINICOM)
	baseX += 4
	//
	name = "mini_menu_man"
	X, _, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(baseX, 410, name, plist_sheet, 0, nil), tools.ISMINICOM)
	//
	baseX += float64(X) + 4
	name = "mini_menu_eq"
	u.AddComponent(QuickCreate(baseX, 410, name, plist_sheet, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				on := i.(*Sprite).imagesName
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
							u.status.ShadowOffsetX = u.status.ShadowOffsetX + 24
							u.status.ShadowOffsetY = u.status.ShadowOffsetY - 84
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
				i.(*Sprite).imagesName = "mini_menu_eq_down"
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).imagesName = on
				u.SetDisplay(tools.ISHIDDEN)
				runtime.GC()
				isClick = false
			}()
		}
	}, true), tools.ISMINICOM)
	X, _, _ = u.GetImagesSize(tools.PlistN, name)
	baseX += float64(X) + 4
	name = "mini_menu_j"
	X, _, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(baseX, 410, name, plist_sheet, 0, nil), tools.ISMINICOM)
	baseX += float64(X) + 4
	name = "mini_menu_m"
	X, _, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(baseX, 410, name, plist_sheet, 0, nil), tools.ISMINICOM)
	baseX += float64(X) + 4
	name = "mini_menu_mess"
	X, _, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(baseX, 410, name, plist_sheet, 0, nil), tools.ISMINICOM)
	baseX += float64(X) + 4
	name = "mini_menu_s"
	X, _, _ = u.GetImagesSize(tools.PlistN, name)
	u.AddComponent(QuickCreate(baseX, 410, name, plist_sheet, 0, nil), tools.ISMINICOM)
	baseX += float64(X) + 4
	name = "mini_menu_st"
	u.AddComponent(QuickCreate(baseX, 410, name, plist_sheet, 0, nil), tools.ISMINICOM)

	u.setHidden(tools.ISHIDDEN)
	u.setHidden(tools.ISMINICOM)

}

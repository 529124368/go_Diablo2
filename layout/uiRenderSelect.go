package layout

import (
	"game/interfaces"
	"game/tools"
	"runtime"
	"sync"
	"time"
)

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
	op = QuickCreate(1250, 850, mgUI, 0, func(i interfaces.SpriteInterface) {
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

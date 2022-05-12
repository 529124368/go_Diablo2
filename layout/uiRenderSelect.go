package layout

import (
	"game/interfaces"
	"game/status"
	"game/tools"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

//加载游戏选择角色UI
func (u *UI) LoadGameCharaSelectImages() {
	//清空
	u.ClearSlice(1)
	//角色选择背景图加载
	s, _ := u.image.ReadFile("resource/UI/charactSelect.png")
	mgUI := tools.GetEbitenImage(s)
	selectSenceBg = mgUI
	//游戏开始按钮
	op := QuickCreate(1250, 850, "startGameButton", plist_sheet, 0, func(i interfaces.SpriteInterface) {
		if !isClick {
			isClick = true
			go func() {
				on := i.(*Sprite).imagesName
				i.(*Sprite).imagesName = "startGameButton_down"
				time.Sleep(tools.CLOSEBTNSLEEP)
				i.(*Sprite).imagesName = on
				u.SetDisplay(tools.ISHIDDEN)
				//切换游戏场景到开门loading
				status.Config.ChangeScenceFlg = true
				status.Config.CurrentGameScence = tools.GAMESCENEOPENDOOR
				status.Config.ChangeScenceFlg = false
				isClick = false
			}()
		}
	}, false)
	op.op.GeoM.Scale(0.5, 0.5)
	u.AddComponent(op, tools.ISNORCOM)
	if plist_png == nil {
		plist, _ := u.image.ReadFile("resource/UI/logo.png")
		plist_json, _ := u.image.ReadFile("resource/UI/logo.json")
		pli, pic := tools.GetImageFromPlist(plist, plist_json)
		plist_sheet = pli
		plist_png = ebiten.NewImageFromImage(pic)
	}

	//加载野蛮人
	plist, _ := u.image.ReadFile("resource/UI/selectRoles.png")
	plist_json, _ := u.image.ReadFile("resource/UI/selectRoles.json")
	pli, pic := tools.GetImageFromPlist(plist, plist_json)
	plist_R_sheet = pli
	plist_R_png = ebiten.NewImageFromImage(pic)
	runtime.GC()
}

package layout

import (
	"fmt"
	"game/tools"

	"github.com/hajimehoshi/ebiten/v2"
)

//加载登录游戏UI
func (u *UI) LoadGameLoginImages() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("has error is :", r)
		}
	}()
	plist, _ := u.image.ReadFile("resource/UI/logo.png")
	plist_json, _ := u.image.ReadFile("resource/UI/logo.json")
	pli, pic := tools.GetImageFromPlist(plist, plist_json)
	plist_sheet = pli
	plist_png = ebiten.NewImageFromImage(pic)

	var len float64 = 0
	var scales float64 = 0.8
	name := "login0"
	op := QuickCreate(len, 0, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += 256

	name = "login1"
	op = QuickCreate(len, 0, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += 256

	name = "login2"
	op = QuickCreate(len, 0, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += 256

	name = "login3"
	op = QuickCreate(len, 0, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len = 0
	var offset float64 = 340
	name = "login8"
	op = QuickCreate(len, offset, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += 256
	name = "login9"
	op = QuickCreate(len, offset, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += 256
	name = "login10"
	op = QuickCreate(len, offset, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += 256
	name = "login11"
	op = QuickCreate(len, offset, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len = 0
	name = "login4"
	op = QuickCreate(len, 256, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += 256
	name = "login5"
	op = QuickCreate(len, 256, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += 256
	name = "login6"
	op = QuickCreate(len, 256, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

	len += 256
	name = "login7"
	op = QuickCreate(len, 256, name, plist_sheet, 0, nil)
	op.op.GeoM.Scale(1, scales)
	u.AddComponent(op, tools.ISNORCOM)

}

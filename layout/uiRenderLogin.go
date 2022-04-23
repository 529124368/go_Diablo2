package layout

import (
	"fmt"
	"game/tools"
	"runtime"
)

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

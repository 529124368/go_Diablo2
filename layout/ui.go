package layout

import (
	"embed"
	"fmt"
	"game/tools"
	"image"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	plist_png     *image.NRGBA
	plist_R_png   *image.NRGBA
	plist_sheet   *texturepacker.SpriteSheet
	plist_R_sheet *texturepacker.SpriteSheet
)

//Create UI Class
type UI struct {
	image         *embed.FS
	compents      []*icon
	Hiddenompents []*icon
}

func NewUI(images *embed.FS) *UI {
	ui := &UI{
		image:         images,
		compents:      make([]*icon, 0, 12),
		Hiddenompents: make([]*icon, 0, 6),
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
	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), false)

	s, _ = u.image.ReadFile("resource/UI/HP.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(28, 480-float64(mgUI.Bounds().Max.Y+13), mgUI, 0, nil), false)

	len += 115

	s, _ = u.image.ReadFile("resource/UI/chisha.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0001.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0002.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0003.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0004.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), false)

	s, _ = u.image.ReadFile("resource/UI/liehuo.png")
	mgUI1 := tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(627, 480-float64(mgUI1.Bounds().Max.Y), mgUI1, 0, nil), false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0005.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil), false)

	s, _ = u.image.ReadFile("resource/UI/MP.png")
	mgUI = tools.GetEbitenImage(s)

	u.AddComponent(QuickCreate(684, 480-float64(mgUI.Bounds().Max.Y+13), mgUI, 1, nil), false)

	s, _ = u.image.ReadFile("resource/UI/skill_btn.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(204, 441, mgUI, 0, nil), false)
	u.AddComponent(QuickCreate(562, 441, mgUI, 0, nil), false)

	//Draw Eq
	s, _ = u.image.ReadFile("resource/UI/eq_0.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395, 0, mgUI, 0, nil), true)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/eq_1.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395+255, 0, mgUI, 0, nil), true)

	len = 395
	s, _ = u.image.ReadFile("resource/UI/bag_0.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395, 176, mgUI, 0, nil), true)

	len += float64(mgUI.Bounds().Max.X)
	s, _ = u.image.ReadFile("resource/UI/bag_1.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(395+255, 176, mgUI, 0, nil), true)

	s, _ = u.image.ReadFile("resource/UI/close_btn.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(414, 384, mgUI, 0, func(i *icon) {
		fmt.Println("close button !!!!!!!!!")
		u.setHidden()
	}, 412, 388, 439, 416), true)

	//Item add Start
	items := getItems()
	for k, v := range items {
		s, _ = u.image.ReadFile("resource/UI/" + k + ".png")
		mgUI = tools.GetEbitenImage(s)
		for _, b := range v {
			res := strings.Split(b, "_")
			x, _ := strconv.ParseFloat(res[0], 64)
			y, _ := strconv.ParseFloat(res[1], 64)
			lay, _ := strconv.Atoi(res[2])
			var isH bool
			if res[3] == "1" {
				isH = true
			} else {
				isH = false
			}
			u.AddComponent(QuickCreate(x, y, mgUI, uint8(lay), nil), isH)
		}
	}
	//

	u.setHidden()

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
	op.pos(len, 0)
	op.op.GeoM.Scale(1, scales)
	op.addImage(mgUI)
	u.AddComponent(op, false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login1.png")
	mgUI = tools.GetEbitenImage(s)
	op1 := newIcon()
	op1.pos(len, 0)
	op1.op.GeoM.Scale(1, scales)
	op1.addImage(mgUI)
	u.AddComponent(op1, false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login2.png")
	mgUI = tools.GetEbitenImage(s)
	op2 := newIcon()
	op2.pos(len, 0)
	op2.op.GeoM.Scale(1, scales)
	op2.addImage(mgUI)
	u.AddComponent(op2, false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login3.png")
	mgUI = tools.GetEbitenImage(s)
	op3 := newIcon()
	op3.pos(len, 0)
	op3.op.GeoM.Scale(1, scales)
	op3.addImage(mgUI)
	u.AddComponent(op3, false)

	len = 0
	var offset float64 = 340
	s, _ = u.image.ReadFile("resource/UI/login8.png")
	mgUI = tools.GetEbitenImage(s)
	op8 := newIcon()
	op8.pos(len, offset)
	op8.op.GeoM.Scale(1, scales)
	op8.addImage(mgUI)
	u.AddComponent(op8, false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login9.png")
	mgUI = tools.GetEbitenImage(s)
	op9 := newIcon()
	op9.pos(len, offset)
	op9.op.GeoM.Scale(1, scales)
	op9.addImage(mgUI)
	u.AddComponent(op9, false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login10.png")
	mgUI = tools.GetEbitenImage(s)
	op10 := newIcon()
	op10.pos(len, offset)
	op10.op.GeoM.Scale(1, scales)
	op10.addImage(mgUI)
	u.AddComponent(op10, false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login11.png")
	mgUI = tools.GetEbitenImage(s)
	op11 := newIcon()
	op11.pos(len, offset)
	op11.op.GeoM.Scale(1, scales)
	op11.addImage(mgUI)
	u.AddComponent(op11, false)

	len = 0

	s, _ = u.image.ReadFile("resource/UI/login4.png")
	mgUI = tools.GetEbitenImage(s)
	op4 := newIcon()
	op4.pos(len, float64(mgUI.Bounds().Max.Y))
	op4.op.GeoM.Scale(1, scales)
	op4.addImage(mgUI)
	u.AddComponent(op4, false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login5.png")
	mgUI = tools.GetEbitenImage(s)
	op5 := newIcon()
	op5.pos(len, float64(mgUI.Bounds().Max.Y))
	op5.op.GeoM.Scale(1, scales)
	op5.addImage(mgUI)
	u.AddComponent(op5, false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login6.png")
	mgUI = tools.GetEbitenImage(s)
	op6 := newIcon()
	op6.pos(len, float64(mgUI.Bounds().Max.Y))
	op6.op.GeoM.Scale(1, scales)
	op6.addImage(mgUI)
	u.AddComponent(op6, false)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login7.png")
	mgUI = tools.GetEbitenImage(s)
	op7 := newIcon()
	op7.pos(len, float64(mgUI.Bounds().Max.Y))
	op7.op.GeoM.Scale(1, scales)
	op7.addImage(mgUI)
	u.AddComponent(op7, false)

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
	op.pos(0, 0)
	op.op.GeoM.Scale(1, 0.8)
	op.addImage(mgUI)
	u.AddComponent(op, false)
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
func (u *UI) AddComponent(s *icon, hasHidden bool) {
	u.compents = append(u.compents, s)
	if hasHidden {
		u.Hiddenompents = append(u.Hiddenompents, s)
	}
}

//
func (u *UI) SetDisplay() {
	for _, v := range u.Hiddenompents {
		v.isDisplay = true
	}
}

//
func (u *UI) setHidden() {
	for _, v := range u.Hiddenompents {
		v.isDisplay = false
	}
}

func (u *UI) ClearSlice(cap int) {
	u.compents = make([]*icon, 0, cap)
	u.Hiddenompents = make([]*icon, 0, cap/2)
}

//Render UI
func (u *UI) DrawUI(screen *ebiten.Image) {
	for _, v := range u.compents {
		if v.layer == 0 && v.isDisplay {
			screen.DrawImage(v.images, v.op)
		}
	}
	for _, v := range u.compents {
		if v.layer == 1 && v.isDisplay {
			screen.DrawImage(v.images, v.op)
		}
	}
}

//Event Listen
func (u *UI) EventLoop() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		for _, v := range u.compents {
			if v.hasEvent == 1 && v.isDisplay {
				x, y := ebiten.CursorPosition()
				if x >= v.clickMinX && x <= v.clickMaxX && y >= v.clickMinY && y <= v.clickMaxY {
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

package layout

import (
	"embed"
	"fmt"
	"game/tools"
	"image"
	"runtime"
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
	op := newIcon()
	op.pos(len, 480-float64(mgUI.Bounds().Max.Y))
	op.addImage(mgUI)
	op.addEvnet(func(i *icon) {
		fmt.Println("click me!!!!!!!!!")
	})
	u.AddComponent(op)

	s, _ = u.image.ReadFile("resource/UI/HP.png")
	mgUI = tools.GetEbitenImage(s)
	opHP := newIcon()
	opHP.pos(28, 480-float64(mgUI.Bounds().Max.Y)-12)
	opHP.addImage(mgUI)
	opHP.addEvnet(func(i *icon) {
		fmt.Println("click me!!!!!!!!!")
	})
	u.AddComponent(opHP)

	len += 115

	s, _ = u.image.ReadFile("resource/UI/chisha.png")
	mgUI = tools.GetEbitenImage(s)
	op1 := newIcon()
	op1.pos(len, 480-float64(mgUI.Bounds().Max.Y))
	op1.addImage(mgUI)
	op1.addEvnet(func(i *icon) {
		fmt.Println("click me!!!!!!!!!")
	})
	u.AddComponent(op1)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0001.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil))

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0002.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil))

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0003.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil))

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0004.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil))

	s, _ = u.image.ReadFile("resource/UI/liehuo.png")
	mgUI1 := tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(627, 480-float64(mgUI1.Bounds().Max.Y), mgUI1, 0, nil))

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/0005.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(len, 480-float64(mgUI.Bounds().Max.Y), mgUI, 0, nil))

	s, _ = u.image.ReadFile("resource/UI/MP.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(684, 480-float64(mgUI.Bounds().Max.Y+13), mgUI, 1, nil))

	s, _ = u.image.ReadFile("resource/UI/skill_btn.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(204, 441, mgUI, 0, nil))
	u.AddComponent(QuickCreate(562, 441, mgUI, 0, nil))

	//TODO item Start
	s, _ = u.image.ReadFile("resource/UI/HP0.png")
	mgUI = tools.GetEbitenImage(s)
	u.AddComponent(QuickCreate(421, 441, mgUI, 0, nil))

	u.AddComponent(QuickCreate(452, 441, mgUI, 0, nil))

	//
	v1 := QuickCreate(415+28*2, 316-62, mgUI, 1, nil)
	u.AddHiddenComponent(v1)
	u.AddComponent(v1)
	v2 := QuickCreate(415+28*3, 316-62, mgUI, 1, nil)
	u.AddHiddenComponent(v2)
	u.AddComponent(v2)
	v3 := QuickCreate(415+28*4, 316-62, mgUI, 1, nil)
	u.AddHiddenComponent(v3)
	u.AddComponent(v3)
	//TODO End

	//Draw Eq
	s, _ = u.image.ReadFile("resource/UI/eq_0.png")
	mgUI = tools.GetEbitenImage(s)
	v4 := QuickCreate(395, 0, mgUI, 0, nil)
	u.AddHiddenComponent(v4)
	u.AddComponent(v4)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/eq_1.png")
	mgUI = tools.GetEbitenImage(s)
	v5 := QuickCreate(395+255, 0, mgUI, 0, nil)
	u.AddHiddenComponent(v5)
	u.AddComponent(v5)

	len = 395
	s, _ = u.image.ReadFile("resource/UI/bag_0.png")
	mgUI = tools.GetEbitenImage(s)
	v6 := QuickCreate(395, 176, mgUI, 0, nil)
	u.AddHiddenComponent(v6)
	u.AddComponent(v6)

	len += float64(mgUI.Bounds().Max.X)
	s, _ = u.image.ReadFile("resource/UI/bag_1.png")
	mgUI = tools.GetEbitenImage(s)
	v7 := QuickCreate(395+255, 176, mgUI, 0, nil)
	u.AddHiddenComponent(v7)
	u.AddComponent(v7)

	s, _ = u.image.ReadFile("resource/UI/close_btn.png")
	mgUI = tools.GetEbitenImage(s)
	v8 := QuickCreate(414, 384, mgUI, 0, func(i *icon) {
		fmt.Println("close button !!!!!!!!!")
		u.setHidden()
	})
	u.AddHiddenComponent(v8)
	u.AddComponent(v8)

	s, _ = u.image.ReadFile("resource/UI/head.png")
	mgUI = tools.GetEbitenImage(s)
	v9 := QuickCreate(531, 0, mgUI, 0, nil)
	u.AddHiddenComponent(v9)
	u.AddComponent(v9)

	s, _ = u.image.ReadFile("resource/UI/futou.png")
	mgUI = tools.GetEbitenImage(s)
	v10 := QuickCreate(413, 120-70, mgUI, 0, nil)
	u.AddHiddenComponent(v10)
	u.AddComponent(v10)

	s, _ = u.image.ReadFile("resource/UI/body.png")
	mgUI = tools.GetEbitenImage(s)
	v11 := QuickCreate(528, 130-63, mgUI, 0, nil)
	u.AddHiddenComponent(v11)
	u.AddComponent(v11)

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
	u.AddComponent(op)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login1.png")
	mgUI = tools.GetEbitenImage(s)
	op1 := newIcon()
	op1.pos(len, 0)
	op1.op.GeoM.Scale(1, scales)
	op1.addImage(mgUI)
	u.AddComponent(op1)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login2.png")
	mgUI = tools.GetEbitenImage(s)
	op2 := newIcon()
	op2.pos(len, 0)
	op2.op.GeoM.Scale(1, scales)
	op2.addImage(mgUI)
	u.AddComponent(op2)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login3.png")
	mgUI = tools.GetEbitenImage(s)
	op3 := newIcon()
	op3.pos(len, 0)
	op3.op.GeoM.Scale(1, scales)
	op3.addImage(mgUI)
	u.AddComponent(op3)

	len = 0
	var offset float64 = 340
	s, _ = u.image.ReadFile("resource/UI/login8.png")
	mgUI = tools.GetEbitenImage(s)
	op8 := newIcon()
	op8.pos(len, offset)
	op8.op.GeoM.Scale(1, scales)
	op8.addImage(mgUI)
	u.AddComponent(op8)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login9.png")
	mgUI = tools.GetEbitenImage(s)
	op9 := newIcon()
	op9.pos(len, offset)
	op9.op.GeoM.Scale(1, scales)
	op9.addImage(mgUI)
	u.AddComponent(op9)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login10.png")
	mgUI = tools.GetEbitenImage(s)
	op10 := newIcon()
	op10.pos(len, offset)
	op10.op.GeoM.Scale(1, scales)
	op10.addImage(mgUI)
	u.AddComponent(op10)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login11.png")
	mgUI = tools.GetEbitenImage(s)
	op11 := newIcon()
	op11.pos(len, offset)
	op11.op.GeoM.Scale(1, scales)
	op11.addImage(mgUI)
	u.AddComponent(op11)

	len = 0

	s, _ = u.image.ReadFile("resource/UI/login4.png")
	mgUI = tools.GetEbitenImage(s)
	op4 := newIcon()
	op4.pos(len, float64(mgUI.Bounds().Max.Y))
	op4.op.GeoM.Scale(1, scales)
	op4.addImage(mgUI)
	u.AddComponent(op4)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login5.png")
	mgUI = tools.GetEbitenImage(s)
	op5 := newIcon()
	op5.pos(len, float64(mgUI.Bounds().Max.Y))
	op5.op.GeoM.Scale(1, scales)
	op5.addImage(mgUI)
	u.AddComponent(op5)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login6.png")
	mgUI = tools.GetEbitenImage(s)
	op6 := newIcon()
	op6.pos(len, float64(mgUI.Bounds().Max.Y))
	op6.op.GeoM.Scale(1, scales)
	op6.addImage(mgUI)
	u.AddComponent(op6)

	len += float64(mgUI.Bounds().Max.X)

	s, _ = u.image.ReadFile("resource/UI/login7.png")
	mgUI = tools.GetEbitenImage(s)
	op7 := newIcon()
	op7.pos(len, float64(mgUI.Bounds().Max.Y))
	op7.op.GeoM.Scale(1, scales)
	op7.addImage(mgUI)
	u.AddComponent(op7)

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
	u.AddComponent(op)
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
func (u *UI) AddComponent(s *icon) {
	u.compents = append(u.compents, s)
}

//Add HiddenComponent
func (u *UI) AddHiddenComponent(s *icon) {
	u.Hiddenompents = append(u.Hiddenompents, s)
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
		//x, y := ebiten.CursorPosition()
		for _, v := range u.compents {
			if v.hasEvent == 1 && v.isDisplay {
				fmt.Println(v.images.Size())
				v.f(v)
			}
		}
	}

}

//GC for loading
func (u *UI) ClearGlobalVariable() {
	plist_R_sheet = nil
	plist_R_png = nil
}

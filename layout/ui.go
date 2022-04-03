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

var plist_png *image.NRGBA
var plist_R_png *image.Paletted
var plist_sheet, plist_R_sheet *texturepacker.SpriteSheet

//Create UI Class
type UI struct {
	image    *embed.FS
	compents []*icon
}

func NewUI(images *embed.FS) *UI {
	ui := &UI{
		image:    images,
		compents: make([]*icon, 0, 12),
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
	go func() {
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
		op2 := newIcon()
		op2.pos(len, 480-float64(mgUI.Bounds().Max.Y))
		op2.addImage(mgUI)
		u.AddComponent(op2)

		len += float64(mgUI.Bounds().Max.X)

		s, _ = u.image.ReadFile("resource/UI/0002.png")
		mgUI = tools.GetEbitenImage(s)
		op3 := newIcon()
		op3.pos(len, 480-float64(mgUI.Bounds().Max.Y))
		op3.addImage(mgUI)
		u.AddComponent(op3)

		len += float64(mgUI.Bounds().Max.X)

		s, _ = u.image.ReadFile("resource/UI/0003.png")
		mgUI = tools.GetEbitenImage(s)
		op4 := newIcon()
		op4.pos(len, 480-float64(mgUI.Bounds().Max.Y))
		op4.addImage(mgUI)
		u.AddComponent(op4)

		len += float64(mgUI.Bounds().Max.X)

		s, _ = u.image.ReadFile("resource/UI/0004.png")
		mgUI = tools.GetEbitenImage(s)
		op5 := newIcon()
		op5.pos(len, 480-float64(mgUI.Bounds().Max.Y))
		op5.addImage(mgUI)
		u.AddComponent(op5)

		s, _ = u.image.ReadFile("resource/UI/liehuo.png")
		mgUI1 := tools.GetEbitenImage(s)
		op6 := newIcon()
		op6.pos(627, 480-float64(mgUI1.Bounds().Max.Y))
		op6.addImage(mgUI1)
		u.AddComponent(op6)

		len += float64(mgUI.Bounds().Max.X)

		s, _ = u.image.ReadFile("resource/UI/0005.png")
		mgUI = tools.GetEbitenImage(s)
		op7 := newIcon()
		op7.pos(len, 480-float64(mgUI.Bounds().Max.Y))
		op7.addImage(mgUI)
		u.AddComponent(op7)

		s, _ = u.image.ReadFile("resource/UI/MP.png")
		mgUI = tools.GetEbitenImage(s)
		op8 := newIcon()
		op8.pos(684, 480-float64(mgUI.Bounds().Max.Y+13))
		op8.addImage(mgUI)
		u.AddComponent(op8)
		runtime.GC()
	}()
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
		plist_R_sheet, plist_R_png = tools.GetImageFromPlistPaletted(plist, plist_json)
		w.Done()
	}()
	w.Wait()
	go func() {
		runtime.GC()
	}()

}

func (u *UI) GetAnimator(flg, name string) (*ebiten.Image, int, int) {
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

func (u *UI) ClearSlice(cap int) {
	u.compents = make([]*icon, 0, cap)
}

//Render UI
func (u *UI) DrawUI(screen *ebiten.Image) {
	for _, v := range u.compents {
		screen.DrawImage(v.images, v.op)
	}
}

//Event Listen
func (u *UI) EventLoop() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		//x, y := ebiten.CursorPosition()
		for _, v := range u.compents {
			if v.flg == 1 {
				fmt.Println(v.images.Size())
				v.f(v)
			}
		}
	}

}

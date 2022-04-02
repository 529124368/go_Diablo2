package layout

import (
	"embed"
	"fmt"
	"game/tools"
	"image"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

var plist_png *image.Paletted
var plist_sheet *texturepacker.SpriteSheet

//UI
type UI struct {
	image    *embed.FS
	compents []*icon
}

func NewUI(images *embed.FS) *UI {
	ui := &UI{
		image:    images,
		compents: make([]*icon, 0, 2),
	}
	return ui
}

func (u *UI) LoadImages() {
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

}

// func (u *UI) GetAnimator(name string) (*ebiten.Image, int, int) {
// 	return ebiten.NewImageFromImage(plist_png.SubImage(plist_sheet.Sprites[name].Frame)), plist_sheet.Sprites[name].SpriteSourceSize.Min.X, plist_sheet.Sprites[name].SpriteSourceSize.Min.Y
// }
func (u *UI) AddComponent(s *icon) {
	u.compents = append(u.compents, s)
}

//render Ui
func (u *UI) DrawUI(screen *ebiten.Image) {
	for _, v := range u.compents {
		screen.DrawImage(v.images, v.op)
	}
}
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

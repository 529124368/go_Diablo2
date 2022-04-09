package layout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//颜色
type RGBColor struct {
	r float64
	g float64
	b float64
	a float64
}

//精灵子类 items专用精灵
type SpriteItems struct {
	Sprite      //继承
	opBg        *ebiten.DrawImageOptions
	imageBg     *ebiten.Image
	bgIsDisplay bool //背景图是否显示
	bgColor     *RGBColor
}

//创建精灵
func newSpriteItems() *SpriteItems {
	s := &SpriteItems{}
	//父类属性
	s.op = new(ebiten.DrawImageOptions)
	s.hasEvent = 0
	s.layer = 0
	s.imagex = 0
	s.imagey = 0
	s.isDisplay = true
	s.size = imageSize{0, 0}
	s.op.Filter = ebiten.FilterLinear
	//本类属性
	s.opBg = new(ebiten.DrawImageOptions)
	s.bgIsDisplay = true
	s.bgColor = &RGBColor{0, 0, 0, 0}
	return s
}

//快速创建items精灵组件
func QuickCreateItems(x, y float64, img *ebiten.Image, layer uint8, callBack func(i spriteInterface), s ...int) *SpriteItems {
	op := newSpriteItems()
	op.SetPosition(x, y)
	if len(s) == 4 {
		//添加点击范围
		op.addClickRange(s[0], s[1], s[2], s[3])
	}
	//判断是否有注册的UI事件
	if callBack != nil {
		op.addEvent(callBack)
	}
	//添加UI显示层级
	op.layer = layer
	//添加图片长度
	width, height := img.Size()
	op.size.width = width
	op.size.height = height
	//添加items背景图
	emptyImage := ebiten.NewImage(width, height)
	emptyImage.Fill(color.White)
	//透明蓝色
	re := width / 28
	ss := height / 28
	if re == 2 && ss == 2 {
		//透明红色
		op.opBg.ColorM.Scale(255, 0, 0, 0.26)
	} else if re == 2 && ss == 3 {
		//透明黄色
		op.opBg.ColorM.Scale(255, 241, 0, 0.26)
	} else if re == 1 && ss == 2 {
		//透明绿色
		op.opBg.ColorM.Scale(0, 255, 0, 0.26)
	} else {
		//透明蓝色
		op.opBg.ColorM.Scale(0, 0, 255, 0.26)
	}
	op.opBg.GeoM.Translate(x, y)
	op.imageBg = emptyImage
	//保存图片
	op.addImage(img)
	return op
}

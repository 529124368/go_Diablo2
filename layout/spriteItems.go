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
	touchEvnet  func(i spriteInterface, x, y int)
	clickEvnet  func(i spriteInterface, x, y int)
	itemName    string
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
func QuickCreateItems(x, y float64, name string, img *ebiten.Image, layer uint8, clickEvnet func(i spriteInterface, x, y int), d uint8, needClickRange ...bool) *SpriteItems {
	op := newSpriteItems()
	op.SetPosition(x, y)
	//判断是否有注册的UI事件
	if clickEvnet != nil {
		op.clickEvnet = clickEvnet
	}
	//item名字设定
	op.itemName = name
	//添加UI显示层级
	op.layer = layer
	//添加图片长度
	width, height := img.Size()
	op.size.width = width
	op.size.height = height
	if len(needClickRange) == 1 && needClickRange[0] {
		//添加点击范围
		op.addClickRange()
	}
	//添加items背景图
	emptyImage := ebiten.NewImage(width, height)
	emptyImage.Fill(color.White)
	//透明蓝色
	re := width / 28
	ss := height / 28
	//图片位置在装备区的单一颜色
	if y < 238 {
		//透明绿色
		op.opBg.ColorM.Scale(0, 255, 0, 0.13)
	} else {
		if re == 2 && ss == 2 {
			//透明红色
			op.opBg.ColorM.Scale(255, 0, 0, 0.13)
		} else {
			//透明蓝色
			op.opBg.ColorM.Scale(0, 0, 255, 0.13)
		}
	}

	op.opBg.GeoM.Translate(x, y)
	op.imageBg = emptyImage
	//背景BG是否显示
	if d == 1 {
		op.bgIsDisplay = true
	} else {
		op.bgIsDisplay = false
	}
	op.hasEvent = 1
	//添加鼠标悬停事件
	op.touchEvnet = func(i spriteInterface, x, y int) {
		//装备栏位置
		if i.(*SpriteItems).imagey < 238 {
			if x >= op.clickMinX && x <= op.clickMaxX && y >= op.clickMinY && y <= op.clickMaxY {
				i.(*SpriteItems).bgIsDisplay = true
			} else {
				i.(*SpriteItems).bgIsDisplay = false
			}
		}
	}
	//保存图片
	op.addImage(img)
	return op
}

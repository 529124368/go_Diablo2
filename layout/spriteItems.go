package layout

import (
	"game/interfaces"
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
	Sprite                                    //继承父结构体
	opBg             *ebiten.DrawImageOptions //背景图片坐标
	imageBg          *ebiten.Image            //背景图片
	opContent        *ebiten.DrawImageOptions //物品介绍文
	imageContent     *ebiten.Image            //物品介绍文坐标
	bgIsDisplay      bool                     //背景图是否显示
	contentIsDisplay bool                     //装备描述是否显示
	bgColor          *RGBColor
	touchEvent       func(i interfaces.SpriteInterface, x, y int)
	clickEvent       func(i interfaces.SpriteInterface, x, y int)
	itemName         string
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
	s.opContent = new(ebiten.DrawImageOptions)
	s.bgIsDisplay = true
	s.contentIsDisplay = false
	s.bgColor = &RGBColor{0, 0, 0, 0}
	return s
}

//快速创建items精灵组件
func QuickCreateItems(x, y float64, name string, img *ebiten.Image, layer uint8, clickEvnet func(i interfaces.SpriteInterface, x, y int), d uint8, needClickRange ...bool) *SpriteItems {
	op := newSpriteItems()
	op.SetPosition(x, y)
	//判断是否有注册的UI事件
	if clickEvnet != nil {
		op.clickEvent = clickEvnet
	}
	//item名字设定
	op.itemName = name
	//添加UI显示层级
	op.layer = layer
	//添加图片长度
	op.size.width, op.size.height = img.Size()
	if len(needClickRange) == 1 && needClickRange[0] {
		//添加点击范围
		op.AddClickRange()
	}
	//添加items背景图
	//如果图片位置位于左右手武器栏
	var GBImage *ebiten.Image

	if (x == 416 || x == 647) && y >= 45 && y <= 154 {
		GBImage = ebiten.NewImage(op.size.width, 116)
	} else {
		GBImage = ebiten.NewImage(op.size.width, op.size.height)
	}
	GBImage.Fill(color.White)
	//透明蓝色
	re := op.size.width / 28
	ss := op.size.height / 28
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
	if x == 416 && y == 60 || x == 647 && y == 60 {
		op.opBg.GeoM.Translate(x, y-15)
	} else if (x == 416 || x == 647) && y == 80 {
		//如果为2*2大小的装备的情况下需要y轴偏移下面距离
		op.opBg.GeoM.Translate(x, y-30)
	} else {
		op.opBg.GeoM.Translate(x, y)
	}

	op.imageBg = GBImage
	//背景BG是否显示
	if d == 1 {
		op.bgIsDisplay = true
	} else {
		op.bgIsDisplay = false
	}
	//物品详细
	ContImage := ebiten.NewImage(op.size.width*2, op.size.height*2)
	ContImage.Fill(color.White)
	op.opContent.ColorM.Scale(0, 0, 0, 0.7)
	op.opContent.GeoM.Translate(x-float64(op.size.width/2), y+25)
	op.imageContent = ContImage
	//
	op.hasEvent = 1
	//添加鼠标悬停事件
	op.touchEvent = func(i interfaces.SpriteInterface, x, y int) {
		//装备栏位置
		if i.(*SpriteItems).imagey < 238 {
			if x >= op.clickMinX && x <= op.clickMaxX && y >= op.clickMinY && y <= op.clickMaxY {
				i.(*SpriteItems).bgIsDisplay = true
			} else {
				i.(*SpriteItems).bgIsDisplay = false
			}
		}
		//显示装备信息
		if x >= op.clickMinX && x <= op.clickMaxX && y >= op.clickMinY && y <= op.clickMaxY {
			i.(*SpriteItems).contentIsDisplay = true
		} else {
			i.(*SpriteItems).contentIsDisplay = false
		}

	}
	//保存图片
	op.AddImage(img)
	return op
}

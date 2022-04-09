package layout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//图片大小
type imageSize struct {
	width  int
	height int
}

//精灵类
type Sprite struct {
	op        *ebiten.DrawImageOptions
	images    *ebiten.Image
	hasEvent  uint8
	layer     uint8
	isDisplay bool
	f         func(i spriteInterface)
	clickMinX int
	clickMinY int
	clickMaxX int
	clickMaxY int
	imagex    float64
	imagey    float64
	size      imageSize
}

//创建精灵
func newSprite() *Sprite {
	s := &Sprite{
		op:        new(ebiten.DrawImageOptions),
		hasEvent:  0,
		layer:     0,
		imagex:    0,
		imagey:    0,
		isDisplay: true,
		size:      imageSize{0, 0},
	}
	s.op.Filter = ebiten.FilterLinear
	return s
}

//获取精灵的图片屏幕坐标
func (s *Sprite) GetPosition() (float64, float64) {
	return s.imagex, s.imagey
}

//获取精灵长宽
func (s *Sprite) GetSpriteSize() (int, int) {
	return s.size.width, s.size.height
}

//设置精灵屏幕坐标
func (s *Sprite) SetPosition(x, y float64) {
	s.op.GeoM.Translate(x, y)
	s.imagex += x
	s.imagey += y
}

//添加图片
func (s *Sprite) addImage(m *ebiten.Image) {
	s.images = m
}

//给UI添加事件
func (s *Sprite) addEvent(fu func(s spriteInterface)) {
	s.hasEvent = 1
	s.f = fu

}

//添加按钮点击范围
func (s *Sprite) addClickRange(minX, minY, maxX, maxY int) {
	s.clickMinX = minX
	s.clickMinY = minY
	s.clickMaxX = maxX
	s.clickMaxY = maxY
}

//快速创建精灵组件
func QuickCreate(x, y float64, img *ebiten.Image, layer uint8, callBack func(i spriteInterface), s ...int) *Sprite {
	op := newSprite()
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
	//保存图片
	op.addImage(img)
	return op
}

//自定义画items的背景
func (s *Sprite) DrawItemsBgByCustom(positionX, postionY float64, width, height int, screen *ebiten.Image) {
	emptyImage := ebiten.NewImage(width, height)
	emptyImage.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	//透明蓝色
	op.ColorM.Scale(0, 0, 255, 0.26)
	op.GeoM.Translate(positionX, postionY)
	screen.DrawImage(emptyImage, op)
}

//快速画items的背景
func (s *Sprite) QuickDrawItemsBg(screen *ebiten.Image) {
	emptyImage := ebiten.NewImage(s.size.width, s.size.height)
	emptyImage.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	re := s.size.width / 28
	if re == 2 {
		//透明红色
		op.ColorM.Scale(255, 0, 255, 0.26)
	} else {
		//透明蓝色
		op.ColorM.Scale(0, 0, 255, 0.26)
	}

	op.GeoM.Translate(s.imagex, s.imagey)
	screen.DrawImage(emptyImage, op)
}

package layout

import (
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
	f         func(i *Sprite)
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
	i := &Sprite{
		op:        new(ebiten.DrawImageOptions),
		hasEvent:  0,
		layer:     0,
		imagex:    0,
		imagey:    0,
		isDisplay: true,
		size:      imageSize{0, 0},
	}
	i.op.Filter = ebiten.FilterLinear
	return i
}

//获取精灵的图片屏幕坐标
func (i *Sprite) GetPosition() (float64, float64) {
	return i.imagex, i.imagey
}

//获取精灵长宽
func (i *Sprite) GetSpriteSize() (int, int) {
	return i.size.width, i.size.height
}

//设置精灵屏幕坐标
func (i *Sprite) SetPosition(x, y float64) {
	i.op.GeoM.Translate(x, y)
	i.imagex += x
	i.imagey += y
}

//添加图片
func (i *Sprite) addImage(m *ebiten.Image) {
	i.images = m
}

//给UI添加事件
func (i *Sprite) addEvent(fu func(i *Sprite)) {
	i.hasEvent = 1
	i.f = fu

}

//添加按钮点击范围
func (i *Sprite) addClickRange(minX, minY, maxX, maxY int) {
	i.clickMinX = minX
	i.clickMinY = minY
	i.clickMaxX = maxX
	i.clickMaxY = maxY
}

//快速创建精灵组件
func QuickCreate(x, y float64, img *ebiten.Image, layer uint8, callBack func(i *Sprite), s ...int) *Sprite {
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

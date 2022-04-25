package maps

import (
	"embed"
	"game/status"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Scale float64 = 1
)

type ImgWall struct {
	Img *ebiten.Image
	H   int
}

type MapBase struct {
	Image   *embed.FS
	BgImage *ebiten.Image
	Status  *status.StatusManage //状态
	Img     [][]*ebiten.Image
	Img2    [][]ImgWall
	Img3    [][]ImgWall
}

//加载地图图片
func (m *MapBase) LoadMap() {
}

//改变地图坐标
func (m *MapBase) ChangeMapTranslate(x, y float64) {
	m.Status.MoveOffsetX += x
}

//渲染地图的地砖
func (m *MapBase) RenderFloor(screen *ebiten.Image, offsetX, offsetY float64) {
	//floor
	sumX := 0
	startY := 0
	for i := 0; i < m.Status.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < m.Status.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			//视野剔除
			if j > m.Status.MapTitleX-m.Status.MapZoom && j < m.Status.MapTitleX+m.Status.MapZoom && i > m.Status.MapTitleY-m.Status.MapZoom && i < m.Status.MapTitleY+m.Status.MapZoom {
				s := m.Img[i][j]
				if s != nil {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(3280+float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY)
					op.GeoM.Scale(Scale, Scale)
					screen.DrawImage(s, op)
					// debug  info
					//ebitenutil.DebugPrintAt(screen, "·"+strconv.Itoa(j)+","+strconv.Itoa(i), i*(-80)+sumX+int(offsetX)+74, startY+j*40+int(offsetY)+m.Img3[i][j].h+37)
				}
			}
		}
	}

}

//渲染地图的建筑
func (m *MapBase) RenderWall(screen *ebiten.Image, offsetX, offsetY float64) {
	//wall
	sumX := 0
	startY := 0
	for i := 0; i < m.Status.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < m.Status.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			//视野剔除
			if j > m.Status.MapTitleX-m.Status.MapZoom && j < m.Status.MapTitleX+m.Status.MapZoom && i > m.Status.MapTitleY-m.Status.MapZoom && i < m.Status.MapTitleY+m.Status.MapZoom {
				s := m.Img2[i][j].Img
				if s != nil {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(3280+float64(i)*(-80)+float64(sumX)+offsetX, float64(startY)+float64(j)*40+offsetY+float64(m.Img2[i][j].H))
					op.GeoM.Scale(Scale, Scale)
					screen.DrawImage(s, op)
				}
			}
		}
	}
}

//获取墙体区域
func (m *MapBase) GetBlock1Aera() [][]ImgWall {
	return m.Img2
}

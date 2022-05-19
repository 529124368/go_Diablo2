package layout

import (
	"embed"
	"errors"
	"fmt"
	"game/controller"
	"game/fonts"
	"game/interfaces"
	"game/status"
	"game/storage"
	"game/tools"
	"image"
	"image/color"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	count                      int  = 0
	change                     bool = false
	selectSenceBg              *ebiten.Image
	plist_png, plist_R_png     *ebiten.Image
	plist_sheet, plist_R_sheet *texturepacker.SpriteSheet
	isClick                    bool = false
	MouseIcon, MouseIconTake   *ebiten.Image
	mouseIconCopy              ebiten.Image
	PHp                        *ebiten.Image //玩家HP
	SkillIcon                  *ebiten.Image //技能显示
	opMouse                    *ebiten.DrawImageOptions
	HPImage                    *ebiten.Image            = nil //血条备份
	HPop                       *ebiten.DrawImageOptions       //血条备份
	DeletedHPSum, DeletedMPSum int                      = 0, 0
	MPImage                    *ebiten.Image            = nil //蓝条备份
	MPop                       *ebiten.DrawImageOptions       //蓝条备份
)

//UI类
type UI struct {
	image             *embed.FS
	Compents          []*Sprite       //普通UI存放集合
	HiddenCompents    []*Sprite       //可以被隐藏的UI组件集合
	MiniPanelCompents []*Sprite       //MINI板的UI集合
	ItemsCompents     []*SpriteItems  //Items的UI集合
	tempBag           [1]*SpriteItems //临时Items存放
	fCont             *fonts.FontBase
	mapManage         interfaces.MapInterface
	bag               *storage.Bag
}

func NewUI(images *embed.FS, f *fonts.FontBase, m interfaces.MapInterface, b *storage.Bag) *UI {

	ui := &UI{
		image:             images,
		Compents:          make([]*Sprite, 0, 12),
		HiddenCompents:    make([]*Sprite, 0, 6),
		MiniPanelCompents: make([]*Sprite, 0, 6),
		ItemsCompents:     make([]*SpriteItems, 0, 10),
		fCont:             f,
		mapManage:         m,
		bag:               b,
	}
	//鼠标Icon设置
	opMouse = &ebiten.DrawImageOptions{}
	ss, _ := ui.image.ReadFile("resource/UI/mouse.png")
	MouseIcon = tools.GetEbitenImage(ss)
	ss, _ = ui.image.ReadFile("resource/UI/mouse_take.png")
	MouseIconTake = tools.GetEbitenImage(ss)
	//
	ss, _ = ui.image.ReadFile("resource/UI/hp_bar.png")
	PHp = tools.GetEbitenImage(ss)
	//
	ss, _ = ui.image.ReadFile("resource/UI/skillIcon.png")
	SkillIcon = tools.GetEbitenImage(ss)
	return ui
}

//获取选择角色的背景图片
func (u *UI) GetSelectBGImage() *ebiten.Image {
	return selectSenceBg
}

//清除选择角色的背景图片
func (u *UI) GCSelectBGImage() {
	selectSenceBg = nil
}

//获取plist的图片大小信息
func (u *UI) GetImagesSize(types uint8, name string) (int, int, error) {
	if types == tools.PlistR {
		return plist_R_sheet.Sprites[name+".png"].SourceSize.X, plist_R_sheet.Sprites[name+".png"].SourceSize.Y, nil
	} else if types == tools.PlistN {
		return plist_sheet.Sprites[name+".png"].SourceSize.X, plist_sheet.Sprites[name+".png"].SourceSize.Y, nil
	}
	return 0, 0, errors.New("has error")
}

//减血
func (u *UI) DeleHP(num int) {
	//u.Compents[1] hard code
	if DeletedHPSum < 95 {
		DeletedHPSum += num
		img, _, _ := u.GetAnimator(tools.PlistN, u.Compents[1].imagesName)
		HPImage = img.SubImage(image.Rectangle{image.Point{176, 1033 + DeletedHPSum}, image.Point{256, 1113}}).(*ebiten.Image)
		HPop = new(ebiten.DrawImageOptions)
		HPop.GeoM.Translate(28, float64(387+DeletedHPSum))
	}
}

//减蓝
func (u *UI) DeleMP(num int) {
	//u.Compents[1] hard code
	if DeletedMPSum < 95 {
		DeletedMPSum += num
		img, _, _ := u.GetAnimator(tools.PlistN, u.Compents[9].imagesName)
		MPImage = img.SubImage(image.Rectangle{image.Point{165, 1115 + DeletedMPSum}, image.Point{245, 1195}}).(*ebiten.Image)
		MPop = new(ebiten.DrawImageOptions)
		MPop.GeoM.Translate(684, float64(387+DeletedMPSum))
	}
}

//图集获取图片
func (u *UI) GetAnimator(flg uint8, name string) (*ebiten.Image, int, int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if flg == tools.PlistR {
		return plist_R_png.SubImage(plist_R_sheet.Sprites[name+".png"].Frame).(*ebiten.Image), plist_R_sheet.Sprites[name+".png"].SpriteSourceSize.Min.X, plist_R_sheet.Sprites[name+".png"].SpriteSourceSize.Min.Y
	} else if flg == tools.PlistN {
		return plist_png.SubImage(plist_sheet.Sprites[name+".png"].Frame).(*ebiten.Image), plist_sheet.Sprites[name+".png"].SpriteSourceSize.Min.X, plist_sheet.Sprites[name+".png"].SpriteSourceSize.Min.Y
	}
	return nil, 0, 0
}

//组件注册
func (u *UI) AddComponent(s interfaces.SpriteInterface, ImageType uint8) {
	if ImageType == tools.ISHIDDEN {
		//将UI压入通用集合
		u.Compents = append(u.Compents, s.(*Sprite))
		//将UI压入可隐藏集合
		u.HiddenCompents = append(u.HiddenCompents, s.(*Sprite))
	} else if ImageType == tools.ISMINICOM {
		//将UI压入通用集合
		u.Compents = append(u.Compents, s.(*Sprite))
		//将UI压入MINI板集合
		u.MiniPanelCompents = append(u.MiniPanelCompents, s.(*Sprite))
	} else if ImageType == tools.ISITEMS {
		//将UI压入Items集合
		u.ItemsCompents = append(u.ItemsCompents, s.(*SpriteItems))
	} else {
		//将UI压入通用集合
		u.Compents = append(u.Compents, s.(*Sprite))
	}
}

//显示UI
func (u *UI) SetDisplay(ImageType uint8) {
	if ImageType == tools.ISHIDDEN {
		status.Config.OpenBag = true
		for _, v := range u.HiddenCompents {
			v.isDisplay = true
		}
	} else {
		status.Config.OpenMiniPanel = true
		for _, v := range u.MiniPanelCompents {
			v.isDisplay = true
		}
	}

}

//隐藏UI
func (u *UI) setHidden(ImageType uint8) {
	if ImageType == tools.ISHIDDEN {
		status.Config.OpenBag = false
		for _, v := range u.HiddenCompents {
			v.isDisplay = false
		}
	} else {
		status.Config.OpenMiniPanel = false
		for _, v := range u.MiniPanelCompents {
			v.isDisplay = false
		}
	}

}

//清除切片
func (u *UI) ClearSlice(cap int) {
	u.Compents = make([]*Sprite, 0, cap)
	u.HiddenCompents = make([]*Sprite, 0, cap/2)
	u.MiniPanelCompents = make([]*Sprite, 0, cap/2)
	u.ItemsCompents = make([]*SpriteItems, 0, 10)
}

//渲染UI
func (u *UI) DrawUI(screen *ebiten.Image) {
	//渲染UI
	for k, v := range u.Compents {
		if v.layer == 0 && v.isDisplay {
			if k == 1 && HPImage != nil {
				screen.DrawImage(HPImage, HPop)
			} else {
				img, _, _ := u.GetAnimator(tools.PlistN, v.imagesName)
				screen.DrawImage(img, v.op)
			}
		}
	}

	//渲染层级为1的UI
	for k, v := range u.Compents {
		if v.layer == 1 && v.isDisplay {
			if k == 9 && MPImage != nil {
				screen.DrawImage(MPImage, MPop)
			} else {
				img, _, _ := u.GetAnimator(tools.PlistN, v.imagesName)
				screen.DrawImage(img, v.op)
			}

		}
	}
	//当包裹打开的时候，渲染包裹内物品和装备 TODO
	if status.Config.OpenBag {
		for _, v := range u.ItemsCompents {
			//先渲染背景色
			if v.bgIsDisplay {
				screen.DrawImage(v.imageBg, v.opBg)
			}
			//再渲染物品
			img, _, _ := u.GetAnimator(tools.PlistR, v.imagesName)
			screen.DrawImage(img, v.op)
		}
		if !status.Config.IsTakeItem {
			//渲染物品信息
			for _, v := range u.ItemsCompents {
				//是否显示物品详细
				if v.contentIsDisplay {
					screen.DrawImage(v.imageContent, v.opContent)
					//Draw Text
					u.fCont.Render(screen, 1, int(v.imagex)-20, int(v.imagey)+50, "属性:xx\n攻击:+2\n敏捷:+4", 7.2, 150, color.RGBA{R: 255, G: 255, B: 255, A: 255})
				}
			}
		}
	}
	//玩家血条显示
	if status.Config.CurrentGameScence == tools.GAMESCENESTART {
		hop := new(ebiten.DrawImageOptions)
		hop.GeoM.Translate(float64(status.Config.PLAYERCENTERX-25), float64(status.Config.PLAYERCENTERY-60))
		screen.DrawImage(PHp, hop)
	}

	//显示Icon
	if status.Config.CurrentGameScence == tools.GAMESCENESTART && status.Config.IsAttack {

		//count++
		hop := new(ebiten.DrawImageOptions)
		hop.GeoM.Scale(0.15, 0.15)
		w, h := SkillIcon.Size()
		hop.GeoM.Translate(-float64(w)*0.15/2, -float64(h)*0.15)
		hop.GeoM.Rotate(90 * 2 * math.Pi / 360)
		hop.GeoM.Translate(float64(tools.LAYOUTX/2), float64(tools.LAYOUTY/2))
		screen.DrawImage(SkillIcon, hop)
	}
}

//事件轮询
func (u *UI) EventLoop(mouseX, mouseY int) {

	if controller.MouseleftPress() || controller.IsTouch() {
		//普通UI事件轮询
		for _, v := range u.Compents {
			if v.hasEvent == 1 && v.isDisplay {
				if mouseX >= v.clickMinX && mouseX <= v.clickMaxX && mouseY >= v.clickMinY && mouseY <= v.clickMaxY {
					//实行回调函数
					v.f(v)
				}
			}
		}
		//包裹打开的情况下监听
		if status.Config.OpenBag {
			//items UI事件轮询
			for _, v := range u.ItemsCompents {
				if v.hasEvent == 1 {
					if mouseX > v.clickMinX && mouseX < v.clickMaxX && mouseY > v.clickMinY && mouseY < v.clickMaxY {
						v.clickEvent(v, mouseX, mouseY)
					}
				}
			}
		}

		//点击包裹区域并且在包裹坐标范围内
		if status.Config.OpenBag && mouseX >= 408 && mouseY >= 6 && mouseX <= 698 && mouseY <= 372 && u.tempBag[0] != nil && status.Config.IsTakeItem {
			s := u.tempBag[0]
			//给鼠标加一个假偏移，防止双击
			if u.AddItemToBag(mouseX+status.Config.Mouseoffset, mouseY+status.Config.Mouseoffset, s.imagesName) {
				//清空缓冲区
				u.ClearTempBag()
			}
		}

	}
	//包裹打开的情况下监听
	if status.Config.OpenBag {
		//items UI事件轮询
		for _, v := range u.ItemsCompents {
			if v.hasEvent == 1 {
				v.touchEvent(v, mouseX, mouseY)
			}
		}
	}

}

//添加物品到包裹 or 装备栏
func (u *UI) AddItemToBag(mousex, mousey int, itemName string) bool {
	//屏幕坐标转换成包裹坐标
	x := int(mousey-254) / 29
	y := int(mousex-413) / 29
	sizeX, sizeY := tools.GetItemsCellSize(itemName)
	if sizeX != 0 && sizeY != 0 {
		//x y这个单元格有位置是否
		if x >= 0 && x <= 3 && y >= 0 && y <= 9 && u.bag.BagLayout[x][y] == "" {
			//是否相同size的时候
			if sizeX == 1 && sizeY == 1 {
				u.bag.BagLayout[x][y] = itemName
				layoutX := 413 + y*29
				layoutY := 254 + x*29
				u.AddComponent(QuickCreateItems(float64(layoutX), float64(layoutY), itemName, plist_R_sheet, 1, u.ItemsEvent(), 1, true), tools.ISITEMS)
				return true
			} else {
				//循环判断是否可以放下
				for i := 0; i < sizeX; i++ {
					for j := 0; j < sizeY; j++ {
						if x+j > 3 || y+i > 9 || u.bag.BagLayout[x+j][y+i] != "" {
							return false
						}
					}
				}
				name := strconv.Itoa(x) + "," + strconv.Itoa(y)
				for i := 0; i < sizeX; i++ {
					for j := 0; j < sizeY; j++ {
						u.bag.BagLayout[x+j][y+i] = itemName + "_" + name
					}
				}
				layoutX := 413 + y*29
				layoutY := 254 + x*29
				u.AddComponent(QuickCreateItems(float64(layoutX), float64(layoutY), itemName, plist_R_sheet, 1, u.ItemsEvent(), 1, true), tools.ISITEMS)
				return true
			}
		} else if mousex >= 397 && mousey >= 5 && mousex <= 705 && mousey <= 247 {
			//判断是否放入装备栏
			return u.JudgeCanToEquip(mousex, mousey, itemName)
		} else {
			return false
		}
	} else {
		return false
	}
}

//捡取物品到包裹
func (u *UI) AddItemToBagByHand(x, y int, itemName string) {
	xx := 413 + y*29
	yy := 254 + x*29
	u.AddComponent(QuickCreateItems(float64(xx), float64(yy), itemName, plist_R_sheet, 1, u.ItemsEvent(), 1, true), tools.ISITEMS)
}

//从包裹删除物品
func (u *UI) DelItemFromBag(imageX, imageY int) {
	//屏幕坐标转换成包裹坐标
	x := int(imageY-254) / 29
	y := int(imageX-413) / 29
	if x >= 0 && x <= 3 && y >= 0 && y <= 9 && u.bag.BagLayout[x][y] != "" {
		if strings.Contains(u.bag.BagLayout[x][y], "_") {
			itemName := u.bag.BagLayout[x][y]
			for i := 0; i < 4; i++ {
				for j := 0; j < 10; j++ {
					if u.bag.BagLayout[i][j] == itemName {
						u.bag.BagLayout[i][j] = ""
					}
				}
			}
		} else {
			u.bag.BagLayout[x][y] = ""
		}
		for k, v := range u.ItemsCompents {
			//根据具体的图片坐标删除 支持唯一性
			if v.imagex == float64(imageX) && v.imagey == float64(imageY) {
				if k != len(u.ItemsCompents)-1 {
					u.ItemsCompents = append(u.ItemsCompents[:k], u.ItemsCompents[k+1:]...)
				} else {
					u.ItemsCompents = u.ItemsCompents[:k]
				}
			}
		}
	} else if xx, _, key := u.JudgeIsEquipArea(imageX, imageY); xx != 0 {
		//删除装备栏
		for k, v := range u.ItemsCompents {
			//根据具体的图片坐标删除 支持唯一性
			if v.imagex == float64(imageX) && v.imagey == float64(imageY) {
				if k != len(u.ItemsCompents)-1 {
					u.ItemsCompents = append(u.ItemsCompents[:k], u.ItemsCompents[k+1:]...)
				} else {
					u.ItemsCompents = u.ItemsCompents[:k]
				}
				u.bag.BagLayout[4][key] = ""
				return
			}
		}
	}
}

//重新绘制鼠标ICON
func (u *UI) DrawMouseIcon(screen *ebiten.Image, mouseX, mouseY int) {
	opMouse.GeoM.Reset()
	opMouse.Filter = ebiten.FilterLinear
	opMouse.GeoM.Translate(float64(mouseX), float64(mouseY))
	if !change {
		screen.DrawImage(MouseIcon, opMouse)
	} else {
		screen.DrawImage(MouseIconTake, opMouse)
	}

}

//判断是否可以放入装备栏
func (u *UI) JudgeCanToEquip(mousex, mousey int, itemName string) bool {
	x, y, key := u.JudgeIsEquipArea(mousex, mousey)
	if x != 0 && u.bag.BagLayout[4][key] == "" {
		yy := plist_R_sheet.Sprites[itemName+".png"].SourceSize.Y
		//左右手武器并且图片高度为2格的情况下
		if (key == 1 || key == 2) && yy/28 == 2 {
			y += 20
		} else if (key == 1 || key == 2) && yy/28 == 4 {
			//左右手武器并且图片高度为4格的情况下
			y -= 15
		}
		u.bag.BagLayout[4][key] = itemName
		u.AddComponent(QuickCreateItems(float64(x), float64(y), itemName, plist_R_sheet, 1, u.ItemsEvent(), 0, true), 0)
		return true
	} else {
		return false
	}
}

//判断是否可以放入装备栏
func (u *UI) InsertToEquip(mousex, mousey int, itemName string) bool {
	x, y, key := u.JudgeIsEquipArea(mousex, mousey)
	if x != 0 {
		yy := plist_R_sheet.Sprites[itemName+".png"].SourceSize.Y
		//左右手武器并且图片高度为2格的情况下
		if (key == 1 || key == 2) && yy/28 == 2 {
			y += 20
		} else if (key == 1 || key == 2) && yy/28 == 4 {
			//左右手武器并且图片高度为4格的情况下
			y -= 15
		}
		u.bag.BagLayout[4][key] = itemName
		u.AddComponent(QuickCreateItems(float64(x), float64(y), itemName, plist_R_sheet, 1, u.ItemsEvent(), 0, true), 0)
		return true
	} else {
		return false
	}
}

//物品事件
func (u *UI) ItemsEvent() func(i interfaces.SpriteInterface, x, y int) {
	//注册监听
	item_event := func(i interfaces.SpriteInterface, x, y int) {
		if !isClick {
			isClick = true
			go func() {
				if !status.Config.IsTakeItem {
					//拿起物品flag设置
					status.Config.IsTakeItem = true
					s := i.(*SpriteItems)
					go func() {
						time.Sleep(tools.CLOSEBTNSLEEP)
						status.Config.Mouseoffset = 0
					}()
					//将拿起的物品放入临时区
					u.tempBag[0] = s
					mouseIconCopy = *MouseIcon
					img, _, _ := u.GetAnimator(tools.PlistR, s.imagesName)
					MouseIcon = img
					//拿起物品，从包裹中删除物品
					u.DelItemFromBag(int(s.imagex), int(s.imagey))
				}
			}()
		}
	}
	return item_event
}

//判断鼠标是否位于装备区
func (u *UI) JudgeIsEquipArea(mousex, mousey int) (int, int, uint8) {
	if mousex >= 530 && mousey >= 3 && mousex <= 584 && mousey <= 54 {
		//判断是否可以放入头盔
		return 530, 3, 0
	} else if mousex >= 416 && mousey >= 45 && mousex <= 469 && mousey <= 154 {
		//判断是否可以放入左武器
		return 416, 60, 1

	} else if mousex >= 530 && mousey >= 74 && mousex <= 583 && mousey <= 154 {
		//判断是否可以放入铠甲
		return 530, 74, 4

	} else if mousex >= 647 && mousey >= 45 && mousex <= 699 && mousey <= 154 {
		//判断是否可以放入右武器
		return 647, 60, 2

	} else if mousex >= 414 && mousey >= 177 && mousex <= 468 && mousey <= 230 {
		//判断是否可以放入手套
		return 414, 177, 5

	} else if mousex >= 646 && mousey >= 177 && mousex <= 699 && mousey <= 229 {
		//判断是否可以放入鞋
		return 646, 177, 9
	} else if mousex >= 600 && mousey >= 32 && mousex <= 622 && mousey <= 58 {
		//判断是否可以放入项链
		return 600, 32, 3
	} else if mousex >= 487 && mousey >= 177 && mousex <= 508 && mousey <= 204 {
		//判断是否可以放入左戒指
		return 487, 177, 6
	} else if mousex >= 600 && mousey >= 177 && mousex <= 622 && mousey <= 205 {
		//判断是否可以放入右戒指
		return 600, 177, 8
	} else if mousex >= 529 && mousey >= 178 && mousex <= 581 && mousey <= 203 {
		//判断是否可以放入腰带
		return 529, 178, 7
	} else {
		return 0, 0, 0
	}
}

//清空缓冲区物品信息 实现真正删除物品
func (u *UI) ClearTempBag() string {
	name := ""
	//鼠标还原
	MouseIcon = &mouseIconCopy
	//清理临时区
	name = u.tempBag[0].imagesName
	u.tempBag[0] = nil
	//恢复防止双击的鼠标偏移量
	status.Config.Mouseoffset = 500
	go func() {
		time.Sleep(tools.CLOSEBTNSLEEP)
		isClick = false
		status.Config.IsTakeItem = false
		status.Config.IsDropDeal = false
	}()
	return name
}

func ChangeMouseicon(i uint) {
	switch i {
	case 1:
		change = true
	default:
		change = false

	}
}

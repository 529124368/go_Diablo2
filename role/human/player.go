package human

import (
	"embed"
	"game/baseClass"
	"game/controller"
	"game/engine/ws"
	"game/engine/ws/pb"
	"game/interfaces"
	"game/layout"
	"game/status"
	"game/tools"
	"math"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	OFFSETX int = -30
	OFFSETY int = -50
)

type Player struct {
	baseClass.PlayerBase                           //继承
	PlayerName           string                    //玩家名字
	MouseX, MouseY       int                       //鼠标X坐标 鼠标Y坐标
	SkillName            string                    //技能名称
	map_c                interfaces.MapInterface   //地图
	uiContr              *layout.UI                //UI
	music                interfaces.MusicInterface //音乐
	newPath              []uint8
	turnLoOp             uint8
	WsCon                *ws.WsNetManage   //net
	imgOffset            [4]tools.OffsetXY //动作图片偏移
}

//创建玩家
func NewPlayer(x, y float64, state, dir uint8, mx, my int, images *embed.FS, m interfaces.MapInterface, con *ws.WsNetManage, ui *layout.UI, ms interfaces.MusicInterface) *Player {
	play := &Player{
		PlayerName: "",
		MouseX:     mx,
		MouseY:     my,
		SkillName:  "", //技能名字
		map_c:      m,
		turnLoOp:   0,
		WsCon:      con,
		uiContr:    ui,
		music:      ms,
	}
	play.X = x //地图坐标X
	play.Y = y //地图坐标Y
	play.State = state
	play.Direction = dir
	play.OldDirection = dir
	play.Asset = images
	play.OpS = &ebiten.DrawImageOptions{}
	play.Op = &ebiten.DrawImageOptions{}
	return play
}

//加載玩家素材
func (p *Player) LoadImages(name, path string, num uint8) {
	p.PlayerBase.LoadImages(name, path, num)
	p.imgOffset = tools.GetOffetByAction(name)
}

//暗黑破坏神 16方位 移动 鼠标控制
func (p *Player) GetMouseController(dir uint8, dx, dy, dis float64, count *int) {
	if p.FlagCanAction {
		speed := 0.0
		//判断是否走路
		if status.Config.IsWalk && (p.Direction != dir || p.State != tools.Walk) {
			speed = tools.SPEED
			p.SetPlayerState(tools.Walk, dir)
		}
		if !status.Config.IsWalk && (p.Direction != dir || p.State != tools.RUN) {
			speed = tools.SPEED_RUN
			p.SetPlayerState(tools.RUN, dir)
		}
		//移动判断
		moveX, moveY := tools.CalculateSpeed(dir, speed, dx, dy, dis)
		if p.CanWalk(moveX, moveY, dir) {
			//相机偏移
			status.Config.CamerOffsetX -= moveX
			status.Config.CamerOffsetY -= moveY
			//玩家偏移
			p.Y += moveY
			p.X += moveX
			//网络  %2 通过加大数字可以降低网服务器发送消息频率
			if status.Config.IsNetPlay && *count%2 == 0 {
				//网络
				act := ""
				if !status.Config.IsWalk {
					act = "r"
				} else {
					act = "w"
				}
				ps := &pb.Player{
					Name:  p.PlayerName,
					X:     p.X,
					Y:     p.Y,
					Dir:   uint32(p.NewDir),
					State: act,
				}
				p.WsCon.SendMessage(true, "@@Move", "", "", ps)
			}
		} else {
			p.NewpositonX = 0
			p.NewpositonY = 0
			status.Config.CalculateEnd = false
		}

	}
}

//判断是否可以行走
func (p *Player) CanWalk(xS, yS float64, dir uint8) bool {
	x, y := tools.GetFloorPositionAt(p.X+xS-110, p.Y+yS+70)
	if x >= status.Config.ReadMapSizeWidth || y >= status.Config.ReadMapSizeHeight || x < 0 || y < 0 {
		p.SetPlayerState(tools.IDLE, dir)
		return false
	}
	//根据地图判断是否可以走
	if p.map_c.GetBlock1Aera(x, y) || p.map_c.GetBlock1AeraUpdate(x, y) {
		return true
	}
	p.SetPlayerState(tools.IDLE, dir)
	return false
}

//本机玩家移动
func (p *Player) PlayerMove(count *int) {
	//判断人物方位
	if p.OldDirection != p.Direction {
		//切换方向
		p.ChangeDir()
	}
	dx := math.Abs(p.X - p.NewpositonX)
	dy := math.Abs(p.Y - p.NewpositonY)
	dis := math.Sqrt(dx*dx + dy*dy)
	if p.NewpositonX != 0 && p.NewpositonY != 0 && dis > 2 {
		status.Config.IsRun = true
		p.FlagCanAction = true
		if !status.Config.CalculateEnd {
			p.GetMouseController(p.NewDir, dx, dy, dis, count)
		}
	} else {
		p.FlagCanAction = false
		p.NewpositonX = 0
		p.NewpositonY = 0
	}

	//玩家停止
	if p.State == tools.IDLE && status.Config.IsRun {
		status.Config.IsRun = false
		//网络
		if status.Config.IsNetPlay {
			ps := &pb.Player{
				Name: p.PlayerName,
				X:    p.X,
				Y:    p.Y,
			}
			p.WsCon.SendMessage(true, "@@MoveEnd", "", "", ps)
		}
	}
}

//改变方向
func (p *Player) ChangeDir() {
	if !status.Config.CalculateEnd {
		p.newPath = tools.CalculateDirPath(p.OldDirection, p.Direction)
		status.Config.CalculateEnd = true
	}
	if len(p.newPath) >= 3 {
		if p.turnLoOp == uint8(len(p.newPath)) {
			//转变方向完成
			p.turnLoOp = uint8(len(p.newPath) - 1)
			p.NewDir = p.newPath[p.turnLoOp]
			//更新旧的方向
			p.OldDirection = p.NewDir
			p.turnLoOp = 0
			status.Config.CalculateEnd = false
		} else {
			p.NewDir = p.newPath[p.turnLoOp]
		}
		p.turnLoOp++
		p.SetPlayerState(tools.IDLE, p.NewDir)
	} else {
		//直接切换方向
		status.Config.CalculateEnd = false
		p.turnLoOp = 0
		//更新旧的方向
		p.OldDirection = p.Direction
	}
}

//玩家到新位置的预算
func (p *Player) PlayerNextMovePositon(mouseX, mouseY int, dir uint8) {
	p.NewDir = dir
	switch p.NewDir {
	case 5, 7:
		p.NewpositonX = p.X + float64(mouseX) - float64(status.Config.PLAYERCENTERX)
		p.NewpositonY = p.Y
	case 4, 6:
		p.NewpositonX = p.X
		p.NewpositonY = p.Y + float64(mouseY) - float64(status.Config.PLAYERCENTERY)
	default:
		p.NewpositonX = p.X + float64(mouseX) - float64(status.Config.PLAYERCENTERX)
		p.NewpositonY = p.Y + float64(mouseY) - float64(status.Config.PLAYERCENTERY)
	}
}

//角色控制
func (p *Player) PlayerContr(count *int) {
	//计算鼠标位置
	dir := tools.CaluteDir(status.Config.PLAYERCENTERX, status.Config.PLAYERCENTERY, int64(p.MouseX), int64(p.MouseY))

	//鼠标事件
	if controller.MouseleftPress() || controller.IsTouch() {
		//防止点击UI界面也移动
		if p.MouseY < 436 {
			p.FlagCanAction = true
		}
		//如果打开包裹，包裹已右位置不能点击移动
		if status.Config.OpenBag && p.MouseX >= tools.LAYOUTX/2 {
			p.FlagCanAction = false
		}
		//如果打开MINi板子，并且没有打开包裹 以下坐标不可以点击移动
		if status.Config.OpenMiniPanel && !status.Config.OpenBag && p.MouseX >= 305 && p.MouseX <= 475 && p.MouseY > 407 {
			p.FlagCanAction = false
		}
		//如果打开MINi板子，并且打开包裹 以下坐标不可以点击移动
		if status.Config.OpenMiniPanel && status.Config.OpenBag && p.MouseX >= 205 && p.MouseX <= 377 && p.MouseY > 407 {
			p.FlagCanAction = false
		}
		//如果拿起物品也不可以移动
		if status.Config.IsTakeItem {
			p.FlagCanAction = false
			//这个范围内就是丢弃物品
			if p.MouseY < 436 && p.MouseX < 390 {
				if !status.Config.IsDropDeal {
					status.Config.IsDropDeal = true
					//播放掉落物品动画
					status.Config.IsPlayDropAnmi = true
					//音乐
					p.music.PlayMusic("diaoluo.mp3", tools.MUSICMP3)
					//丢弃物品
					status.Config.DropItemName = p.uiContr.ClearTempBag()
				}
			}
		}
		//移动
		if p.FlagCanAction {
			//计算新的位置
			p.PlayerNextMovePositon(p.MouseX, p.MouseY, dir)
		}
	}
	//玩家移动监听
	p.PlayerMove(count)

	//攻击和技能
	if !p.FlagCanAction {
		//普通攻击
		if controller.MouseRightPress() && !status.Config.IsTakeItem {
			if p.State != tools.ATTACK {
				p.Counts = 0
			}
			p.SetPlayerState(tools.ATTACK, dir)
			//网络
			if status.Config.IsNetPlay && *count%2 == 0 {
				act := "act"
				ps := &pb.Player{
					Name:  p.PlayerName,
					X:     p.X,
					Y:     p.Y,
					Dir:   uint32(p.Direction),
					State: act,
				}
				p.WsCon.SendMessage(true, "@@Attack", "", "", ps)
			}
		}
		//技能
		if controller.MousePressF1() && !status.Config.IsTakeItem {
			//音乐
			p.music.PlayMusic("File00002184.wav", tools.MUSICWAV)
			//g.player.SkillName = "狂风"
			if p.State != tools.SkILL {
				p.Counts = 0
			}
			p.SetPlayerState(tools.SkILL, dir)
		}
	}
}

//渲染角色
func (p *Player) Render(screen *ebiten.Image) {
	//改变帧数
	p.ChangeFrame()
	//渲染角色
	p.PlayerBase.Render()
	//写入缓存
	var name strings.Builder
	block := 1
	//nameSkill := ""
	name.WriteString(strconv.Itoa(int(p.Direction)))
	switch p.State {
	case tools.ATTACK:
		status.Config.IsAttack = true
		name.WriteString("_attack_")
	case tools.SkILL:
		status.Config.IsAttack = true
		block = 2
		if p.Counts >= 14 {
			p.Counts = 0
		}
		name.WriteString("_skill_")
	case tools.IDLE:
		status.Config.IsAttack = false
		name.WriteString("_stand_")
	case tools.Walk:
		status.Config.IsAttack = false
		if p.Counts >= 8 {
			p.Counts = 0
		}
		name.WriteString("_run_")
	case tools.RUN:
		status.Config.IsAttack = false
		block = 2
		if p.Counts >= 8 {
			p.Counts = 0
		}
		name.WriteString("_run2_")
	}
	name.WriteString(strconv.Itoa(p.Counts))
	imagess, x, y := p.GetAnimator("man", name.String(), uint8(block))
	//Idel -> Walk Offset
	if p.State == tools.Walk {
		x += p.imgOffset[0].X
		y += p.imgOffset[0].Y
	}
	//Idel -> RUN Offset
	if p.State == tools.RUN {
		x += p.imgOffset[1].X
		y += p.imgOffset[1].Y
	}

	//Idel -> Walk -> Attack Offset
	if p.State == tools.ATTACK {
		x += p.imgOffset[2].X
		y += p.imgOffset[2].Y
	}

	//Idel -> SKILL-> Offset
	if p.State == tools.SkILL {
		x += p.imgOffset[3].X
		y += p.imgOffset[3].Y
	}

	//Draw Shadow
	p.OpS.GeoM.Reset()
	p.OpS.Filter = ebiten.FilterLinear
	p.OpS.GeoM.Rotate(-0.45)
	p.OpS.GeoM.Scale(1, 0.5)
	p.OpS.ColorM.Scale(0, 0, 0, 1)
	p.OpS.GeoM.Translate(float64(tools.LAYOUTX/2+x+status.Config.ShadowOffsetX+status.Config.UIOFFSETX), float64(tools.LAYOUTY/2+y+status.Config.ShadowOffsetY))
	screen.DrawImage(imagess, p.OpS)
	//Draw Player
	p.Op.GeoM.Reset()
	p.Op.GeoM.Translate(float64(tools.LAYOUTX/2+OFFSETX+x+status.Config.UIOFFSETX), float64(tools.LAYOUTY/2+OFFSETY+y))
	p.Op.Filter = ebiten.FilterLinear
	screen.DrawImage(imagess, p.Op)

	//Draw Skill
	// if g.player.State == ATTACK {
	// 	imagey, x, y := g.player.GetAnimator("skill", nameSkill)
	// 	//skill Option
	// 	OpSkill = &ebiten.DrawImageOptions{}
	// 	OpSkill.GeoM.Translate(float64(SCREENWIDTH/2+x), float64(SCREENHEIGHT/2+y))
	// 	OpSkill.CompositeMode = ebiten.CompositeModeLighter
	// 	OpSkill.GeoM.Scale(1.5, 1.5)
	// 	OpSkill.Filter = ebiten.FilterLinear
	// 	screen.DrawImage(imagey, OpSkill)
	// }
}

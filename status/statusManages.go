package status

type StatusManage struct {
	Flg             bool
	ChangeScenceFlg bool
	DoorCountFlg    bool
	LoadingFlg      bool
	MusicIsPlay     bool
	OpenBag         bool
	OpenMiniPanel   bool
	CalculateEnd    bool
	UIOFFSETX       int
	ShadowOffsetX   int
	ShadowOffsetY   int
	PLAYERCENTERX   int64
	PLAYERCENTERY   int64
	IsTakeItem      bool //是否拿起物品
	Mouseoffset     int
	MoveOffsetX     float64
	MoveOffsetY     float64
}

func NewStatusManage() *StatusManage {
	n := &StatusManage{
		Flg:             false,
		ChangeScenceFlg: false,
		DoorCountFlg:    false,
		LoadingFlg:      false,
		MusicIsPlay:     false,
		OpenBag:         false,
		OpenMiniPanel:   false,
		CalculateEnd:    false,
		UIOFFSETX:       0,
		ShadowOffsetX:   -348,
		ShadowOffsetY:   361,
		PLAYERCENTERX:   388, //LAYOUTX/2
		PLAYERCENTERY:   242, //LAYOUTY/2
		IsTakeItem:      false,
		Mouseoffset:     500,
		MoveOffsetX:     200,
		MoveOffsetY:     -1300,
	}
	return n
}

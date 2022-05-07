package status

type StatusManage struct {
	Flg               bool //是否可以移动标志
	ChangeScenceFlg   bool
	DoorCountFlg      bool
	LoadingFlg        bool
	MusicIsPlay       bool
	OpenBag           bool
	OpenMiniPanel     bool
	IsWalk            bool
	CalculateEnd      bool
	UIOFFSETX         int
	ShadowOffsetX     int
	PLAYERCENTERX     int64
	PLAYERCENTERY     int64
	IsTakeItem        bool //是否拿起物品
	Mouseoffset       int
	CamerOffsetX      float64
	CamerOffsetY      float64
	ReadMapSizeWidth  int
	ReadMapSizeHeight int
	MapTitleX         int
	MapTitleY         int
	MapZoom           int
	CurrentGameScence int
	DisPlayDebugInfo  bool //是否显示Debug信息
	IsPlayDropAnmi    bool //是否播放掉落物品动画
	IsDropDeal        bool //是否掉落物品处理中
	DisplaySort       bool //人物和物体渲染顺序
	Queue             chan []byte
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
		IsWalk:          true,
		CalculateEnd:    false,
		UIOFFSETX:       0,
		ShadowOffsetX:   -70,
		PLAYERCENTERX:   395, //LAYOUTX/2
		PLAYERCENTERY:   240, //LAYOUTY/2
		IsTakeItem:      false,
		Mouseoffset:     -1800,
		//玩家初始位置偏移设置
		CamerOffsetX: -5280 + 395,
		CamerOffsetY: -1880 + 240,
		//读取地图的尺寸
		ReadMapSizeWidth:  0,
		ReadMapSizeHeight: 0,
		//玩家初始逻辑地图坐标
		MapTitleX: 36,
		MapTitleY: 11,
		//
		MapZoom:           8,
		CurrentGameScence: 1,
		DisPlayDebugInfo:  false,
		IsPlayDropAnmi:    false,
		IsDropDeal:        false,
		DisplaySort:       false,
		Queue:             make(chan []byte),
	}
	return n
}

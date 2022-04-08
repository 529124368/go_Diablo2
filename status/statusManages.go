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
		ShadowOffsetX:   -350,
		ShadowOffsetY:   365,
	}
	return n
}

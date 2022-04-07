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
	ForCalculateEnd bool
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
		ForCalculateEnd: false,
	}
	return n
}

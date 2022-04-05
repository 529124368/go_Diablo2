package runTime

type StatusManage struct {
	Flg             bool
	ChangeScenceFlg bool
	DoorCountFlg    bool
	LoadingFlg      bool
	MusicIsPlay     bool
}

func NewStatusManage() *StatusManage {
	n := &StatusManage{
		Flg:             false,
		ChangeScenceFlg: false,
		DoorCountFlg:    false,
		LoadingFlg:      false,
		MusicIsPlay:     false,
	}
	return n
}

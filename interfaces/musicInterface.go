package interfaces

type MusicInterface interface {
	PlayMusic(name string, ty int)
	CloseBGMusic()
	PauseBGMusic()
	PlayBGMusic(name string, ty int)
	IsPlayingBGMusic() bool
}

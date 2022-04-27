package interfaces

import "github.com/hajimehoshi/ebiten/v2/audio"

type MusicInterface interface {
	PlayMusic(name string, ty int, types uint8)
	CloseMusic(types uint8)
	GetAudioContext(types uint8) *audio.Player
	PauseMusic(types uint8)
	IsPlayingMusic(types uint8) bool
}

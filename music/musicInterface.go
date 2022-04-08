package music

import "github.com/hajimehoshi/ebiten/v2/audio"

type MusicInterface interface {
	PlayMusic(name string, ty int)
	CloseMusic()
	GeyAudioContext() *audio.Player
}

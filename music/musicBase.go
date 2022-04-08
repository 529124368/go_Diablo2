package music

import (
	"bytes"
	"embed"
	"game/tools"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var cont *audio.Context

type MusicBGM struct {
	image        *embed.FS
	AudioContext *audio.Player
}

func NewMusicBGM(images *embed.FS) *MusicBGM {
	cont = audio.NewContext(48000)
	m := &MusicBGM{
		image:        images,
		AudioContext: nil,
	}
	return m
}

func (m *MusicBGM) PlayMusic(name string, ty int) {
	switch ty {
	case tools.MUSICMP3:
		go func() {
			bgm, _ := m.image.ReadFile("resource/BGM/" + name)
			ss, _ := mp3.Decode(cont, bytes.NewReader(bgm))
			m.AudioContext = nil
			m.AudioContext, _ = cont.NewPlayer(ss)
			m.AudioContext.Play()
		}()
	case tools.MUSICWAV:
		go func() {
			bgm, _ := m.image.ReadFile("resource/BGM/" + name)
			ss, _ := wav.Decode(cont, bytes.NewReader(bgm))
			m.AudioContext = nil
			m.AudioContext, _ = cont.NewPlayer(ss)
			m.AudioContext.Play()
		}()
	}
}

func (m *MusicBGM) CloseMusic() {
	go func() {
		m.AudioContext.Close()
	}()
}

func (m *MusicBGM) GeyAudioContext() *audio.Player {
	return m.AudioContext
}

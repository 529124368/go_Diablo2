package music

import (
	"bytes"
	"embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var cont *audio.Context

type MusicBGM struct {
	image        *embed.FS
	audioContext *audio.Player
}

func NewMusicBGM(images *embed.FS) *MusicBGM {
	cont = audio.NewContext(48000)
	m := &MusicBGM{
		image:        images,
		audioContext: nil,
	}
	return m
}

func (m *MusicBGM) PlayMusic(name, ty string) {
	switch ty {
	case "mp3":
		go func() {
			bgm, _ := m.image.ReadFile("resource/BGM/" + name)
			ss, _ := mp3.Decode(cont, bytes.NewReader(bgm))
			m.audioContext = nil
			m.audioContext, _ = cont.NewPlayer(ss)
			m.audioContext.Play()
		}()
	case "wav":
		go func() {
			bgm, _ := m.image.ReadFile("resource/BGM/" + name)
			ss, _ := wav.Decode(cont, bytes.NewReader(bgm))
			m.audioContext = nil
			m.audioContext, _ = cont.NewPlayer(ss)
			m.audioContext.Play()
		}()
	}
}

func (m *MusicBGM) CloseMusic() {
	go func() {
		m.audioContext.Close()
	}()
}

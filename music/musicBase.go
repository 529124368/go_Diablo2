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
	image           *embed.FS
	AudioContScence *audio.Player
	AudioContBGM    *audio.Player
	BgmMusicName    string
	ScenceMusicName string
}

func NewMusicBGM(images *embed.FS) *MusicBGM {
	cont = audio.NewContext(48000)
	m := &MusicBGM{
		image:           images,
		AudioContScence: nil,
		AudioContBGM:    nil,
		BgmMusicName:    "",
		ScenceMusicName: "",
	}
	return m
}

func (m *MusicBGM) PlayMusic(name string, ty int, types uint8) {
	//判断是否加载过得音乐
	if types == tools.SceneMusic || types == tools.BgmMusic && name != m.BgmMusicName {
		switch ty {
		case tools.MUSICMP3:
			go func() {
				bgm, _ := m.image.ReadFile("resource/BGM/" + name)
				ss, _ := mp3.DecodeWithSampleRate(48000, bytes.NewReader(bgm))
				if types == tools.SceneMusic {
					m.ScenceMusicName = name
					m.AudioContScence, _ = cont.NewPlayer(ss)
					m.AudioContScence.Play()

				} else {
					m.BgmMusicName = name
					m.AudioContBGM, _ = cont.NewPlayer(ss)
					m.AudioContBGM.Play()

				}
			}()
		case tools.MUSICWAV:
			go func() {
				bgm, _ := m.image.ReadFile("resource/BGM/" + name)
				ss, _ := wav.DecodeWithSampleRate(48000, bytes.NewReader(bgm))
				if types == tools.SceneMusic {
					m.ScenceMusicName = name
					m.AudioContScence, _ = cont.NewPlayer(ss)
					m.AudioContScence.Play()
				} else {
					m.BgmMusicName = name
					m.AudioContBGM, _ = cont.NewPlayer(ss)
					m.AudioContBGM.Play()
				}
			}()
		}
	}
}

func (m *MusicBGM) CloseMusic(types uint8) {
	if types == tools.BgmMusic {
		m.AudioContBGM.Close()
	} else {
		m.AudioContScence.Close()
	}
}

func (m *MusicBGM) PauseMusic(types uint8) {
	if types == tools.BgmMusic {
		m.AudioContBGM.Pause()
	} else {
		m.AudioContScence.Pause()
	}
}
func (m *MusicBGM) IsPlayingMusic(types uint8) bool {
	if types == tools.BgmMusic {
		return m.AudioContBGM.IsPlaying()
	} else {
		return m.AudioContScence.IsPlaying()
	}
}

func (m *MusicBGM) GetAudioContext(types uint8) *audio.Player {
	if types == tools.BgmMusic {
		return m.AudioContBGM
	} else {
		return m.AudioContScence
	}
}
